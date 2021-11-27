package client

import (
	"context"

	"github.com/wuqinqiang/easycar/pkg/common"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

var _ TransactionInterface = &TCC{}

type TCCOption func(tcc *TCC)

var (
	DefaultProtoCol = "http"
)

type RMFunc func(rmFunc *RM) ([]common.BranchData, error)

type TCC struct {
	protoCol string
}

func NewTCC(options ...TCCOption) *TCC {
	tcc := &TCC{protoCol: DefaultProtoCol}
	for _, option := range options {
		option(tcc)
	}
	return tcc
}

func (tcc *TCC) GetTransactionName() common.TransactionName {
	return "tcc"
}

func (tcc *TCC) GetProtocol() string {
	return tcc.protoCol
}

// WeGo begin a transaction for tcc
func (tcc *TCC) WeGo(ctx context.Context, serverAddress string, rmFunc RMFunc) error {
	rm := NewRM(serverAddress)
	_, err := rm.Start(tcc)
	if err != nil {
		// todo log
		return err
	}
	// step1:register all branch that
	// step2: call branch try url
	branchList, err := rmFunc(rm)
	if err != nil {
		// todo log
		return err
	}
	// register all branch that participated in this transaction
	err = rm.RegisterBranch(ctx, branchList)
	if err != nil {
		return err
	}

	defer func() {
		action := utils.IF(err != nil, "/abort", "/commit").(string)
		var (
			reportReq  common.ReportStateData
			reportResp common.RespBase
		)
		reportReq.SetGId(rm.gId)
		err = common.RestyClient.PostJson(rm.serverAddress+action, reportReq, &reportResp)
		if err != nil {
			// todo log....
			return
		}
	}()
	// then we should call every branch for prepare resource, then abort the transaction  when any branch response err,
	// if err ==nil util last branch call over, it means that the transaction is success
	return tcc.PrepareBranch(ctx, branchList)
}

func (tcc *TCC) PrepareBranch(ctx context.Context, branchList []common.BranchData) (err error) {
	if len(branchList) == 0 {
		return nil
	}
	for _, branch := range branchList {
		var (
			resp common.RespBase
		)
		err = common.RestyClient.PostJson(branch.GetUrl(), branch.GetReqData(),
			&resp, common.SetTimeOut(5))
		if err != nil {
			return err
		}
	}
	return
}

// ask myself some questions
// 1.how to easy to use?
//
