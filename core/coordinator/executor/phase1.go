package executor

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
)

type Phase1 struct {
	list entity.BranchList
	*Executor
}

func Phase1Executor(branchList entity.BranchList) *Phase1 {
	return &Phase1{list: branchList, Executor: &defaultExecutor}
}

func (e *Phase1) Execute(ctx context.Context) error {
	if len(e.list) == 0 {
		return nil
	}
	return e.execute(ctx, e.list, func(branch *entity.Branch) bool {
		return branch.IsTccTry() || branch.IsSAGANormal()
	})
}
