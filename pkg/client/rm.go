package client

import (
	"fmt"

	"github.com/wuqinqiang/easycar/pkg/common"
	"github.com/wuqinqiang/easycar/pkg/entity"
)

type RMInterface interface {
	BranchCommit()
	BranchRollBack()
	GetTransactionName() common.TransactionName
	GetProtocol() string
	RegisterBranch(gId string)
}

type RM struct {
	processorManager map[common.TransactionName]RMInterface
}

func NewRM() *RM {
	rm := &RM{processorManager: make(map[common.TransactionName]RMInterface)}
	// tcc
	rm.processorManager["tcc"] = NewTCC()
	return rm
}

func (rm *RM) RegisterBranch(gId string, name common.TransactionName, BranchId string) {
	// to do register to tc
}

func (rm *RM) Commit(gId string, branch entity.Branch) error {
	op, ok := rm.processorManager[branch.GetTransactionName()]
	if !ok {
		return fmt.Errorf("not found %s transaction model", branch.GetTransactionName())
	}
	op.BranchCommit()
	return nil
}

func (rm *RM) Rollback(gId string, branch entity.Branch) error {
	op, ok := rm.processorManager[branch.GetTransactionName()]
	if !ok {
		return fmt.Errorf("not found %s transaction model", branch.GetTransactionName())
	}
	op.BranchRollBack()
	return nil
}
