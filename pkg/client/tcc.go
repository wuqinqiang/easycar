package client

import (
	"context"

	"github.com/wuqinqiang/easycar/pkg/common"
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
	// step1:register branch
	// step2: call branch try url
	branchList, err := rmFunc(rm)
	if err != nil {
		// todo log
		return err
	}

	// first register all branch
	err = rm.RegisterBranch(ctx, branchList)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			// to abort
		}

		// to commit
	}()

	// then call every branch try url,we should abort handler when any branch call err,
	// if err ==nil util last branch call over, it means that the transaction is success
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
	return nil
}

// ask myself some questions
// 1.how to easy to use?
//
