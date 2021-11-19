package service

import (
	"context"

	"github.com/wuqinqiang/easycar/core/service/entity"
)

type TCC struct {
	*entity.Global
}

func (t *TCC) HandleBranches(ctx context.Context, branchList []*entity.Branch) error {
	if len(branchList) == 0 {
		return nil
	}
	globalState := t.GetState()
	if globalState == entity.SucceedState || globalState == entity.FailedState {
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
