package executor

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
)

type Phase2 struct {
	list entity.BranchList
	*executor
}

func NewPhase2Executor(branchList entity.BranchList) *Phase2 {
	return &Phase2{list: branchList, executor: GetExecutor()}
}

func (e *Phase2) Execute(ctx context.Context) error {
	if len(e.list) == 0 {
		return nil
	}
	return e.execute(ctx, e.list, func(branch *entity.Branch) bool {
		return branch.IsTccTry()
	})
}
