package core

import (
	"context"
	"errors"

	"github.com/wuqinqiang/easycar/core/executor"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/mode"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/dao"
)

var (
	ErrGlobalNotExist = errors.New("global not exist")
)

type Coordinator struct {
	// resty timeout
	dao dao.TransactionDao
}

func NewCoordinator(dao dao.TransactionDao) *Coordinator {
	c := &Coordinator{
		dao: dao,
	}
	return c
}

func (c *Coordinator) Begin(ctx context.Context) (string, error) {
	gid := GetGid()

	g := entity.NewGlobal(gid)
	g.SetState(consts.Begin)
	err := c.dao.CreateGlobal(ctx, g)
	return gid, err
}

func (c *Coordinator) Register(ctx context.Context, gId string, branches entity.BranchList) error {
	global, err := c.dao.GetGlobal(ctx, gId)
	if err != nil {
		return err
	}
	if global.IsEmpty() {
		return ErrGlobalNotExist
	}
	return c.dao.CreateBatches(ctx, branches)
}

func (c *Coordinator) Commit(ctx context.Context, global entity.Global) error {
	branches, err := c.dao.GetBranchList(ctx, global.GetGId())
	if err != nil {
		return err
	}
	err = executor.NewCommitExecutor(branches).Execute(ctx)
	return err
}

func (c *Coordinator) Rollback(ctx context.Context, global entity.Global) error {
	branches, err := c.dao.GetBranchList(ctx, global.GetGId())
	if err != nil {
		return err
	}
	if err = executor.NewRollbackExecutor(branches).Execute(ctx); err != nil {
		return err
	}
	return err
}

func (c *Coordinator) GetGlobal(ctx context.Context, gid string) (entity.Global, error) {
	return c.dao.GetGlobal(ctx, gid)
}

func (c *Coordinator) GetBranchList(ctx context.Context, gid string) (list []*entity.Branch, err error) {
	return c.dao.GetBranchList(ctx, gid)
}

func (c *Coordinator) GetMode(branch entity.Branch) Mode {
	switch branch.TranType {
	case consts.SAGA:
		return mode.NewSaga()
	case consts.TCC:
		return mode.NewTcc()
	}
	panic("not support")
}
