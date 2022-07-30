package executor

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"
)

type Phase2 struct {
	list   entity.BranchList
	global *entity.Global
	*executor
}

func NewPhase2Executor(global *entity.Global, branchList entity.BranchList) *Phase2 {
	return &Phase2{list: branchList, global: global, executor: GetExecutor()}
}

func (e *Phase2) Execute(ctx context.Context) error {
	if len(e.list) == 0 {
		return nil
	}
	return e.execute(ctx, e.list, func(branch *entity.Branch) bool {
		if e.global.State == consts.Phase1Success {
			if branch.IsTccConfirm() {
				return true
			}
			return false
		}
		// other phase1 failed
		if branch.IsSAGACompensation() || branch.IsTccCancel() {
			return true
		}
		return false
	})
}
