package executor

import (
	"context"
	"fmt"
	"sort"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core/protocol"
	"github.com/wuqinqiang/easycar/core/protocol/common"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"
)

type (
	// FilterFn is a function that filters branches
	FilterFn func(branch *entity.Branch) bool

	Executor interface {
		Execute(ctx context.Context) error
	}
)

type executor struct {
	// MustInitExecutor todo add  config
}

var (
	e *executor
)

func init() {
	e = &executor{}
}

func GetExecutor() *executor {
	return e
}

func (e *executor) execute(ctx context.Context, branches entity.BranchList, filterFn FilterFn) error {
	phaseList := make([]entity.BranchList, len(branches))

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
		phaseList[bucketIndex] = append(phaseList[bucketIndex], branch)
	}

	errGroup, groupCtx := errgroup.WithContext(ctx)

	for _, tierList := range phaseList {
		if len(tierList) == 0 {
			continue
		}
		for _, branch := range tierList {
			b := branch
			errGroup.Go(func() error {
				net, err := protocol.GetTransport(common.Net(b.Protocol), b.Url)
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

				if _, err = net.Request(groupCtx, req); err != nil {
					fmt.Printf("[Executor] Request branch:%vrequest error:%v", b, err)
					errmsg = err.Error()
					branchState = consts.BranchFailState
				}
				b.State = branchState

				// todo replace with dao
				if _, erro := dao.GetTransaction().UpdateBranchStateByGid(ctx, b.BranchId,
					b.State, errmsg); err != nil {
					fmt.Printf("[Executor]update branch state error:%v\n", erro)
				}
				return err
			})
		}
		if err := errGroup.Wait(); err != nil {
			return err
		}
	}
	return nil
}
