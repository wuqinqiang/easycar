package executor

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/wuqinqiang/easycar/tools/retry"

	"github.com/pkg/errors"

	"github.com/wuqinqiang/easycar/core/protocol"
	"github.com/wuqinqiang/easycar/core/protocol/common"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/tools/fx"

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

	var (
		wg sync.WaitGroup
	)

	for _, list := range phaseList {
		var (
			err = errors.New("[execute]")
		)
		branches := list
		wg.Add(1)
		tools.GoSafe(func() {
			defer wg.Done()
			fx.From(func(source chan<- interface{}) {
				for i := range branches {
					source <- branches[i]
				}
			}).Walk(func(item interface{}, pipe chan<- interface{}) {
				b, ok := item.(*entity.Branch)
				if !ok {
					pipe <- fmt.Errorf("[Executor]invalid branch:%+v", item)
					return
				}
				transport, err := protocol.GetTransport(common.NetType(b.Protocol), b.Url)
				if err != nil {
					pipe <- fmt.Errorf("[Executor]branchid:%vget transport error:%v", b.BranchId, err)
					return
				}
				req := common.NewReq([]byte(b.ReqData), nil)
				r := retry.NewRetry(2, 2, func() error {
					_, err = transport.Request(ctx, req)
					return err
				})
				if err = r.Run(); err != nil {
					// todo update branch status
					pipe <- fmt.Errorf("branch:%vrequest error:%v", b, err)
				}
			}).ForEach(func(item interface{}) {
				erro, ok := item.(error)
				if !ok {
					return
				}
				err = errors.WithMessagef(err, erro.Error())
			})
		})
		wg.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}
