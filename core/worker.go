package core

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/pkg/errors"

	"github.com/wuqinqiang/easycar/core/protocol"
	"github.com/wuqinqiang/easycar/core/protocol/common"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/tools/fx"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"
)

type Executor struct {
	GID           string
	PhaseBranches []entity.BranchList
}

func NewWorker(gid string, branches entity.BranchList) *Executor {
	w := &Executor{
		GID:           gid,
		PhaseBranches: make([]entity.BranchList, len(branches)),
	}

	sort.Slice(branches, func(i, j int) bool {
		return branches[i].Level < branches[j].Level
	})
	var previousLevel consts.Level = 1
	bucketIndex := 0
	for i, branch := range branches {
		// todo 修改成二阶段需要
		if !branch.IsSAGANormal() && !branch.IsTccTry() {
			continue
		}
		if i == 0 {
			previousLevel = branch.Level
		}
		if branch.Level > previousLevel {
			bucketIndex += 1
			previousLevel = branch.Level
		}
		w.PhaseBranches[bucketIndex] = append(w.PhaseBranches[bucketIndex], branch)
	}
	return w
}

func (w *Executor) Commit(ctx context.Context) error {
	if len(w.PhaseBranches) == 0 {
		return nil
	}
	var (
		wg sync.WaitGroup
	)
	for _, list := range w.PhaseBranches {
		var (
			err error
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
					pipe <- fmt.Errorf("invalid branch:%+v", item)
					return
				}
				transport, err := protocol.GetTransport(common.NetType(b.Protocol), b.Url)
				if err != nil {
					pipe <- fmt.Errorf("branchid:%vget transport error:%v", b.BranchId, err)
					return
				}
				// todo replace []byte(b.ReqData)
				req := common.NewReq([]byte(b.ReqData), nil)
				if _, err = transport.Request(ctx, req); err != nil {
					// todo update branch status
					pipe <- fmt.Errorf("branch:%vrequest error:%v", b, err)
					return
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
