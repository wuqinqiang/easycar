package core

import (
	"context"
	"errors"

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
	restyTimeout int64 //second
	dao          dao.TransactionDao
}

func NewCoordinator(dao dao.TransactionDao) *Coordinator {
	c := &Coordinator{
		restyTimeout: 60,
		dao:          dao,
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
	return nil
}

func (c *Coordinator) Rollback(ctx context.Context, global entity.Global) error {
	return nil
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
