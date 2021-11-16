package service

import (
	"context"
	"github.com/wuqinqiang/easycar/internal/dao"
	"github.com/wuqinqiang/easycar/internal/service/entity"
)

type TMInterface interface {
}

type TMOption func(rm *TM)

type TM struct {
	dao dao.TransactionDao
}

func NewTM(dao dao.TransactionDao) *TM {
	return &TM{dao: dao}
}

// Begin  begin a new transaction, return globalId
func (rm *TM) Begin(ctx context.Context, entity *entity.Global) (gId string, err error) {
	_, err = rm.dao.Create(ctx, entity)
	if err != nil {
		return "", BeginTransactionErr
	}
	return entity.GetGId(), nil
}

// Submit summit transaction
func (rm *TM) Submit(gId string) {
	panic("dd")
}

func (rm *TM) RegisterTccBranch(ctx context.Context, gId string, branchList []*entity.Branch) error {
	_, err := rm.dao.First(ctx, gId)
	if err != nil {
		return NotFindTransaction
	}
	err = rm.dao.CreateBatches(ctx, gId, branchList)
	if err != nil {
		return err
	}
	return nil
}

func (rm *TM) BranchRollBack() {

}
