package core

import (
	"context"

	entity2 "github.com/wuqinqiang/easycar/core/entity"
)

type TCC struct {
	*entity2.Global
}

func (t *TCC) HandleBranches(ctx context.Context, branchList []*entity2.Branch) error {
	if len(branchList) == 0 {
		return nil
	}
	globalState := t.GetState()
	if globalState == entity2.SucceedState || globalState == entity2.FailedState {
		return nil
	}

	for i := range branchList {
		if !branchList[i].CanHandle() {
			continue
		}
		if t.IsGrpc() {
			// todo  handler grpc
			return nil
		}
		// todo http
	}
	return nil
}
