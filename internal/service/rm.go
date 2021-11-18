package service

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/internal/dao"
	"github.com/wuqinqiang/easycar/internal/service/entity"
)

type TMInterface interface {
}

type AddProcessorFunc func(global *entity.Global) Processor

type Processor interface {
	HandleBranches(ctx context.Context, branchList []*entity.Branch) error
}

type TM struct {
	dao              dao.TransactionDao
	processorManager map[entity.TransactionName]AddProcessorFunc
}

func NewTM(dao dao.TransactionDao) *TM {
	tm := &TM{dao: dao}
	tm.processorManager["tcc"] = func(global *entity.Global) Processor {
		return &TCC{global}
	}

	// todo saga and more
	return tm
}

func (rm *TM) GetGlobal(ctx context.Context, gId string) (*entity.Global, error) {
	return rm.dao.First(ctx, gId)
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
	rowsAffected, err = rm.dao.UpdateGlobalStateByGid(ctx, gId, entity.SucceedState)
	if err != nil {
		return err
	}
	rowsAffectedBranch, err = rm.dao.UpdateBranchStateByGid(ctx, gId, entity.BranchSucceedState)
	if err != nil {
		return err
	}
	return nil
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

func (rm *TM) Abort(ctx context.Context, gId string) error {
	global, err := rm.dao.First(ctx, gId)
	if err != nil {
		return err
	}
	list, err := rm.dao.GetBranchList(ctx, gId)
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return fmt.Errorf("branch  must not empty")
	}
	err = rm.processorManager[global.GetTransactionName()](global).HandleBranches(ctx, list)

	// todo if handle branch err, must update global state
	if err != nil {
		return err
	}
	return nil
}
