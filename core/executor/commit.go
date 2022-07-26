package executor

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
)

type CommitExecutor struct {
	list entity.BranchList
	*executor
}

func NewCommitExecutor(branchList entity.BranchList) *CommitExecutor {
	return &CommitExecutor{list: branchList, executor: GetExecutor()}
}

func (e *CommitExecutor) Execute(ctx context.Context) error {
	if len(e.list) == 0 {
		return nil
	}
	return e.execute(ctx, e.list, func(branch *entity.Branch) bool {
		return branch.IsTccTry() || branch.IsSAGANormal()
	})
}
