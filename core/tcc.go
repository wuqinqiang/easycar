package core

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/pkg/entity"

	"github.com/wuqinqiang/easycar/pkg/utils"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type TCCServer struct {
	*entity.Global
}

func (t *TCCServer) ProcessBranchList(ctx context.Context, branchList []*entity.Branch) error {
	if len(branchList) == 0 {
		return nil
	}
	globalState := t.GetState()
	if globalState == Succeed || globalState == Failed {
		return nil
	}
	action := utils.IF(t.GetState() == consts.Submitted, consts.Confirm, consts.Cancel).(string)
	for _, branch := range branchList {
		if branch.GetBranchType() != consts.BranchAction(action) {
			continue
		}

		if !branch.CanHandle() {
			continue
		}
		if t.IsGrpc() {
			// todo  handler grpc
			return nil
		}
		var (
			resp common.RespBase
		)
		err := common.RestyClient.PostJson(branch.GetUrl()+action, branch.GetReqData(), resp)
		if err != nil {
			// todo log
			return err
		}
	}
	return nil
}
