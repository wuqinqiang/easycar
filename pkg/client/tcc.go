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
	RM
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

func (tcc *TCC) RegisterBranch(ctx context.Context) {
	tcc.RM.RegisterBranch(ctx, nil)
}

// WeGo begin a transaction for tcc
func (tcc *TCC) WeGo(ctx context.Context, serverAddress string, rmFunc RMFunc) error {
	rm := NewRM(serverAddress)
	gId, err := rm.Start(tcc)
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
	rm.RegisterBranch(ctx, branchList)

	return nil
}

// ask myself some questions
// 1.how to easy to use?
//
