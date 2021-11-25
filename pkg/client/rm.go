package client

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type RM struct {
	// server Address
	serverAddress string
	// the transaction global id
	gId string
	// todo some timeout config
}

type TransactionInterface interface {
	GetTransactionName() common.TransactionName
	GetProtocol() string
}

func NewRM(tmServer string) *RM {
	return &RM{serverAddress: tmServer}
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
	err = common.RestyClient.PostJson(r.serverAddress+"/begin", req, resp)
	if err != nil {
		// todo common err handler
		return
	}
	r.gId = resp.GetGId()
	return r.gId, nil
}

// RegisterBranch Register transaction branchList to server
func (r *RM) RegisterBranch(ctx context.Context, branchList []common.BranchData) error {
	if len(branchList) == 0 {
		return fmt.Errorf("")
	}
	var (
		resp common.RespBase
	)
	err := common.RestyClient.
		PostJson(r.serverAddress+"/registerBranch", branchList, resp)
	if err != nil {
		// todo handler err for common
		return fmt.Errorf("er")
	}
	return nil
}

// Commit commit the transaction when all branch success
func (r *RM) Commit(ctx context.Context) error {
	var (
		req  common.ReqCommitData
		resp common.RespBase
	)
	req.SetGId(r.gId)
	err := common.RestyClient.PostJson(r.serverAddress+"/commit", req, resp)
	if err != nil {

	}
	return nil
}

// Fail fail when any transaction to err
func (r *RM) Fail(ctx context.Context) error {
	var (
		req  common.ReqFailData
		resp common.RespBase
	)
	req.SetGId(r.gId)
	err := common.RestyClient.PostJson(r.serverAddress+"/fail", req, resp)
	if err != nil {
	}
	return nil
}
