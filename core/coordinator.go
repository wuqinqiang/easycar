package core

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/mode"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/dao"
)

type Coordinator struct {
	// resty timeout
	restyTimeout int64 //second
	dao          dao.TransactionDao
}

func NewCoordinator() *Coordinator {
	return &Coordinator{
		restyTimeout: 60,
		// todo more
	}
}

func (c *Coordinator) Begin(ctx context.Context) (entity.Global, error) {
	return entity.Global{}, nil
}

func (c *Coordinator) Register(ctx context.Context, gId string, branches []*entity.Branch) error {
	return nil
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
