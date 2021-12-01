package core

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/pkg/common"

	"github.com/wuqinqiang/easycar/core/dao"
	entity2 "github.com/wuqinqiang/easycar/core/entity"
)

type TMInterface interface {
	Begin(ctx context.Context, entity *entity2.Global) (gId string, err error)
}

type AddProcessorFunc func(global *entity2.Global) Processor

type Processor interface {
	ProcessBranchList(ctx context.Context, branchList []*entity2.Branch) error
}

type TM struct {
	dao              dao.TransactionDao
	processorManager map[common.TransactionName]AddProcessorFunc
}

func NewTM(dao dao.TransactionDao) *TM {
	tm := &TM{dao: dao}

	// tcc
	tm.processorManager["tcc"] = func(global *entity2.Global) Processor {
		return &TCC{global}
	}
	go func() {
		// todo
	}()
	// todo saga and more
	return tm
}

func (tm *TM) GetGlobal(ctx context.Context, gId string) (*entity2.Global, error) {
	return tm.dao.First(ctx, gId)
}

// Begin  begin a new transaction, return globalId
func (tm *TM) Begin(ctx context.Context, entity *entity2.Global) (gId string, err error) {
	_, err = tm.dao.Create(ctx, entity)
	if err != nil {
		return "", BeginTransactionErr
	}
	return entity.GetGId(), nil
}

// Submit submit transaction
func (tm *TM) Submit(ctx context.Context, gId string) (err error) {
	var (
		rowsAffected, rowsAffectedBranch int64
	)
	defer func() {
		if err == nil && (rowsAffected == 0 || rowsAffectedBranch == 0) {
			panic("submit something wrong")
		}
	}()

	rowsAffected, err = tm.dao.UpdateGlobalStateByGid(ctx, gId, common.Succeed)
	if err != nil {
		return err
	}
	rowsAffectedBranch, err = tm.dao.UpdateBranchStateByGid(ctx, gId, common.BranchSucceedState)
	if err != nil {
		return err
	}
	return nil
}

func (tm *TM) RegisterTccBranch(ctx context.Context, gId string, branchList []*entity2.Branch) error {
	_, err := tm.dao.First(ctx, gId)
	if err != nil {
		return NotFindTransaction
	}
	err = tm.dao.CreateBatches(ctx, gId, branchList)
	if err != nil {
		return err
	}
	return nil
}

func (tm *TM) Abort(ctx context.Context, gId string) error {
	global, err := tm.dao.First(ctx, gId)
	if err != nil {
		return err
	}
	list, err := tm.dao.GetBranchList(ctx, gId)
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return fmt.Errorf("branch  must not empty")
	}
	err = tm.GetProcessorByName(global.GetTransactionName(), global).ProcessBranchList(ctx, list)
	// todo if handle branch err, must update global state
	if err != nil {
		return err
	}
	return nil
}

func (tm *TM) GetProcessorByName(name common.TransactionName, global *entity2.Global) Processor {
	return tm.processorManager[name](global)
}
