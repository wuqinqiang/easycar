package core

import (
	"context"

	"github.com/wuqinqiang/easycar/pkg/entity"

	"github.com/wuqinqiang/easycar/pkg/utils"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type TCC struct {
	*entity.Global
}

func (t *TCC) ProcessBranchList(ctx context.Context, branchList []*entity.Branch) error {
	if len(branchList) == 0 {
		return nil
	}
	globalState := t.GetState()
	if globalState == common.Succeed || globalState == common.Failed {
		return nil
	}
	action := utils.IF(t.GetState() == common.Submitted, common.Confirm, common.Cancel).(string)
	for _, branch := range branchList {
		if branch.GetBranchType() != common.BranchType(action) {
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
