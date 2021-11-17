package service

import (
	"context"
	"github.com/wuqinqiang/easycar/internal/dao"
	"github.com/wuqinqiang/easycar/internal/dao/gorm"
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
func (rm *TM) Submit(ctx context.Context, gId string) (err error) {
	var (
		rowsAffected, rowsAffectedBranch int64
	)
	defer func() {
		if err == nil && (rowsAffected == 0 || rowsAffectedBranch == 0) {
			panic("submit something wrong")
		}
	}()

	// todo 严重依赖gorm
	return gorm.GenManager.BeginTransaction(ctx, func(ctx context.Context) error {
		rowsAffected, err = rm.dao.UpdateGlobalStateByGid(ctx, gId, entity.SubmittedState)
		if err != nil {
			return err
		}
		rowsAffectedBranch, err = rm.dao.UpdateBranchStateByGid(ctx, gId, entity.BranchSucceedState)
		if err != nil {
			return err
		}
		return nil
	})
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
