package client

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type RM struct {
	// server Address
	serverAddress string
	// todo some timeout config
}

type TransactionInterface interface {
	GetTransactionName() common.TransactionName
	GetProtocol() string
}

func NewRM(serverAddress string) *RM {
	return &RM{serverAddress: serverAddress}
}

// Start start for a transaction
func (r *RM) Start(tran TransactionInterface) (gId string, err error) {
	if r.serverAddress == "" {
		return "", fmt.Errorf("no serverAddress url")
	}
	req := new(common.ReqGlobalData)
	req.SetTransactionName(tran.GetTransactionName())
	req.SetProtocol(tran.GetProtocol())

	resp := new(common.RespGlobalData)
	err = common.RestyClient.PostJson(r.serverAddress, req, resp)
	if err != nil {
		// todo common err handler
		return
	}
	return resp.GetGId(), nil
}

// RegisterBranch Register transaction branchList to server
func (r *RM) RegisterBranch(ctx context.Context, branchList []common.BranchData) {

}
