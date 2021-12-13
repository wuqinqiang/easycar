package core

import (
	"context"

	entity2 "github.com/wuqinqiang/easycar/pkg/entity"

	"github.com/wuqinqiang/easycar/pkg/common"

	"github.com/wuqinqiang/easycar/core/dao"
)

type CoordinatorInterface interface {
	Begin(ctx context.Context, entity *entity2.Global) (gId string, err error)
	BranchRegister(ctx context.Context, branch *entity2.Branch) (state string, err error)
	Commit(ctx context.Context, gId string) (state string, err error)
	RollBack(ctx context.Context, gId string) (state string, err error)
}

type Coordinator struct {
	// resty timeout
	restyTimeout int64 //second
	dao          dao.TransactionDao
}

func NewCoordinator(dao dao.TransactionDao) *Coordinator {
	d := &Coordinator{dao: dao}
	go func() {
		d.scheduling()
	}()
	return d
}

func (d *Coordinator) scheduling() {
}

func (d *Coordinator) GetGlobal(ctx context.Context, gId string) (*entity2.Global, error) {
	return d.dao.First(ctx, gId)
}

// Begin  begin a new transaction, return globalId
func (d *Coordinator) Begin(ctx context.Context, entity *entity2.Global) (gId string, err error) {
	_, err = d.dao.Create(ctx, entity)
	if err != nil {
		return "", BeginTransactionErr
	}
	return entity.GetGId(), nil
}

// Submit submit transaction
func (d *Coordinator) Submit(ctx context.Context, gId string) (err error) {
	global, err := d.dao.First(ctx, gId)
	if err != nil {
		return err
	}
	if global == nil {
		return NotFindTransaction
	}
	if !global.CanSubmit() {
		return nil
	}
	var (
		rowsAffected, rowsAffectedBranch int64
	)
	defer func() {
		if err == nil && (rowsAffected == 0 || rowsAffectedBranch == 0) {
			panic("submit something wrong")
		}
	}()

	rowsAffected, err = d.dao.UpdateGlobalStateByGid(ctx, gId, common.Succeed)
	if err != nil {
		return err
	}
	// todo notice every branch to commit
	err = d.handlerBranch(gId, true)
	if err != nil {
		return err
	}
	rowsAffectedBranch, err = d.dao.UpdateBranchStateByGid(ctx, gId, common.BranchSucceedState)
	if err != nil {
		return err
	}
	return nil
}

func (d *Coordinator) RegisterTccBranch(ctx context.Context, gId string, branchList []*entity2.Branch) error {
	_, err := d.dao.First(ctx, gId)
	if err != nil {
		return NotFindTransaction
	}
	err = d.dao.CreateBatches(ctx, gId, branchList)
	if err != nil {
		return err
	}
	return nil
}

func (d *Coordinator) RollBack(ctx context.Context, gId string) error {
	global, err := d.dao.First(ctx, gId)
	if err != nil {
		return err
	}
	if global == nil {
		return NotFindTransaction
	}
	if !global.CanRollBack() {
		return nil
	}
	// todo validate something
	_, err = d.dao.UpdateGlobalStateByGid(ctx, gId, common.Failed)
	if err != nil {
		return err
	}
	err = d.handlerBranch(gId, true)
	if err != nil {
		return err
	}
	return nil
}

func (d *Coordinator) handlerBranch(gId string, isCommit bool) error {
	return nil
}
