package executor

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
)

type RollbackExecutor struct {
	list entity.BranchList
	*executor
}

func NewRollbackExecutor(branchList entity.BranchList) *RollbackExecutor {
	return &RollbackExecutor{list: branchList, executor: GetExecutor()}
}

func (e *RollbackExecutor) Execute(ctx context.Context) error {
	if len(e.list) == 0 {
		return nil
	}
	return e.execute(ctx, e.list, func(branch *entity.Branch) bool {
		return branch.IsTccTry()
	})
}
