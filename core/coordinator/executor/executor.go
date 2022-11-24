package executor

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/wuqinqiang/easycar/tools/retry"

	"go.opentelemetry.io/otel/codes"

	. "github.com/wuqinqiang/easycar/tracing"
	"go.opentelemetry.io/otel/attribute"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/transport"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/dao"
	"github.com/wuqinqiang/easycar/core/entity"
	"github.com/wuqinqiang/easycar/core/transport/common"
	"golang.org/x/sync/errgroup"
)

type (

	// FilterFn is a function that filters branches
	FilterFn func(branch *entity.Branch) bool

	executor struct {
		manager transport.Manager
		timeout time.Duration
	}
)

func NewExecutor() *executor {
	return DefaultExecutor
}

func (e *executor) Close(ctx context.Context) error {
	return e.manager.Close(ctx)
}

func (e *executor) Phase1(ctx context.Context, g *entity.Global) error {
	branches, err := dao.GetTransaction().GetBranches(ctx, g.GetGId())
	if err != nil {
		return err
	}
	return e.execute(ctx, true, branches, func(branch *entity.Branch) bool {
		if branch.Success() {
			return false
		}

		return branch.TccTry() || branch.SAGANormal()
	})
}

func (e *executor) Phase2(ctx context.Context, g *entity.Global) error {
	branches, err := dao.GetTransaction().GetBranches(ctx, g.GetGId())
	if err != nil {
		return err
	}

	return e.execute(ctx, false, branches, func(branch *entity.Branch) bool {
		if branch.Success() {
			return false
		}
		// for commit
		if g.GotoCommit() {
			return branch.TccConfirm()
		}

		// phase1 failed、rollbackIng and rollbackFailed
		if branch.SAGACompensation() || branch.TccCancel() {
			return true
		}
		return false
	})
}

func (e *executor) stratify(branches entity.BranchList) []entity.BranchList {
	layerList := make([]entity.BranchList, len(branches))
	// sort branches by level
	sort.Slice(branches, func(i, j int) bool {
		return branches[i].Level < branches[j].Level
	})

	var (
		previousLevel consts.Level = 1 // first level default 1
		bucketIndex                = 0 // first level bucket index default 0
	)

	for i, branch := range branches {
		if i == 0 {
			previousLevel = branch.Level
		}
		if branch.Level > previousLevel {
			bucketIndex += 1
			previousLevel = branch.Level
		}
		layerList[bucketIndex] = append(layerList[bucketIndex], branch)
	}
	return layerList
}

func (e *executor) execute(ctx context.Context, shouldStratify bool, branches entity.BranchList, filterFn FilterFn) error {

	// filter branches
	for i := 0; i < len(branches); {
		if !filterFn(branches[i]) {
			branches = append(branches[:i], branches[i+1:]...)
		} else {
			i++
		}
	}

	// if it's phase2,no layering is required
	// issue:https://github.com/wuqinqiang/easycar/issues/43
	layeredList := []entity.BranchList{branches}
	if shouldStratify {
		layeredList = e.stratify(branches)
	}

	for _, branchGroup := range layeredList {
		if len(branchGroup) == 0 {
			continue
		}

		errGroup, _ := errgroup.WithContext(ctx)
		for _, branch := range branchGroup {
			b := branch
			errGroup.Go(func() error {
				bCtx, span := Tracer(ctx, "reqRM")
				span.SetAttributes(
					attribute.String("protocol", b.Protocol),
					attribute.String("reqUrl", b.Url),
					attribute.String("branchId", b.BranchId),
				)
				defer span.End()

				var (
					err         error
					errmsg      string
					branchState = consts.BranchSucceed
				)

				// request the RM
				if err = e.request(ctx, b); err != nil {
					logging.Errorf("[Executor] request branch %+v err:%v", b, err)
					branchState = consts.BranchFailState
					span.SetStatus(codes.Error, err.Error())
					errmsg = err.Error()
				}

				// update branch state
				b.State = branchState
				if _, erro := dao.GetTransaction().UpdateBranchStateByGid(bCtx, b.BranchId,
					b.State, errmsg); erro != nil {
					logging.Errorf("[Executor]update branch state error:%v", erro)
				}
				return err
			})
		}

		//in phase1, we have to stop execution and don't go to the next level RM if some RM is wrong
		if err := errGroup.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func (e *executor) request(ctx context.Context, b *entity.Branch) (err error) {
	transporter, err := e.manager.GetTransporter(common.Net(b.Protocol))
	if err != nil {
		return fmt.Errorf("[Executor]branchid:%vget transport error:%v", b.BranchId, err)
	}

	defer func() {
		if err != nil {
			if errors.Is(err, retry.ErrOverMaximumAttempt) {
				logging.Warnf("over maximum attempt")
			}
			err = fmt.Errorf("[Executor] Request branchid:%vrequest error:%v", b.BranchId, err)
		}
	}()

	var (
		reqOpts []common.Option
	)

	timeout := e.timeout
	if b.Timeout > 0 {
		timeout = time.Second * time.Duration(b.Timeout)
	}
	reqOpts = append(reqOpts, common.WithTimeout(timeout))

	req := common.NewReq([]byte(b.ReqData), []byte(b.ReqHeader), reqOpts...)
	req.AddEasyCarHeaders(b.GID, b.BranchId)

	r := retry.New(2, retry.WithMaxBackOffTime(1*time.Second))

	err = r.Run(func() error {
		_, err = transporter.Request(ctx, b.Url, req)
		return err
	})
	return
}
