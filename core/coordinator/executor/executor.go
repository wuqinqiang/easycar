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
	}
)

func NewExecutor() *executor {
	executor := DefaultExecutor
	return executor
}

func (e *executor) Close(ctx context.Context) error {
	return e.manager.Close(ctx)
}

func (e *executor) Phase1(ctx context.Context, _ *entity.Global, branches entity.BranchList) error {
	return e.execute(ctx, branches, func(branch *entity.Branch) bool {
		return branch.TccTry() || branch.SAGANormal()
	})
}

func (e *executor) Phase2(ctx context.Context, global *entity.Global, branches entity.BranchList) error {
	return e.execute(ctx, branches, func(branch *entity.Branch) bool {
		if global.State == consts.Phase1Success {
			return branch.TccConfirm()
		}
		// other phase1 failed
		if branch.SAGACompensation() || branch.TccCancel() {
			return true
		}
		return false
	})
}

func (e *executor) stratify(branches entity.BranchList, filterFn FilterFn) []entity.BranchList {
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
		if !filterFn(branch) {
			continue
		}
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

func (e *executor) execute(ctx context.Context, branches entity.BranchList, filterFn FilterFn) error {
	layerList := e.stratify(branches, filterFn)

	if len(layerList) == 0 {
		return nil
	}

	for _, branchItems := range layerList {
		if len(branchItems) == 0 {
			continue
		}

		errGroup, groupCtx := errgroup.WithContext(ctx)

		for _, branch := range branchItems {
			b := branch
			errGroup.Go(func() error {
				transporter, err := e.manager.GetTransporter(common.Net(b.Protocol))
				if err != nil {
					return fmt.Errorf("[Executor]branchid:%vget transport error:%v", b.BranchId, err)
				}
				var (
					reqOpts []common.ReqOpt
				)
				if b.Timeout > 0 {
					reqOpts = append(reqOpts, common.WithTimeOut(time.Duration(b.Timeout)*time.Second))
				}
				req := common.NewReq([]byte(b.ReqData), []byte(b.ReqHeader), reqOpts...)
				req.AddEasyCarHeaders(b.GID, b.BranchId)

				var (
					branchState = consts.BranchSucceed
					errmsg      string
				)
				bCtx, span := Tracer(ctx, "reqRM")
				span.SetAttributes(
					attribute.String("protocol", string(transporter.GetType())),
					attribute.String("reqUrl", b.Url),
					attribute.String("branchId", b.BranchId),
				)
				r := retry.New(3, retry.WithMaxBackOffTime(1*time.Second))

				err = r.Run(func() error {
					_, err = transporter.Request(groupCtx, b.Url, req)
					return err
				})

				if err != nil {
					if errors.Is(err, retry.ErrOverMaximumAttempt) {
						logging.Warnf("over maximum attempt")
					}
					logging.Error(fmt.Sprintf("[Executor] Request branch:%vrequest error:%v", b, err))
					errmsg = err.Error()
					span.SetStatus(codes.Error, errmsg)
					branchState = consts.BranchFailState
				}
				b.State = branchState

				if _, erro := dao.GetTransaction().UpdateBranchStateByGid(bCtx, b.BranchId,
					b.State, errmsg); erro != nil {
					logging.Error(fmt.Sprintf("[Executor]update branch state error:%v\n", erro))
				}
				span.End()
				return err
			})
		}
		if err := errGroup.Wait(); err != nil {
			return err
		}
	}
	return nil
}
