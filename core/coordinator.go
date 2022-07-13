package core

import (
	"context"
	"errors"

	"github.com/wuqinqiang/easycar/core/protocol"
	"github.com/wuqinqiang/easycar/core/protocol/common"

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

func (c *Coordinator) handler(ctx context.Context, gid string,
	fn func(b *entity.Branch) error) error {
	branches, err := c.dao.GetBranchList(ctx, gid)
	if err != nil {
		return err
	}
	for i := range branches {
		//todo warp err
		err = fn(branches[i])
		continue
		// 处理分支
	}
	return nil
}

func (c *Coordinator) Commit(ctx context.Context, gid string) error {
	global, err := c.dao.GetGlobal(ctx, gid)
	if err != nil {
		return err
	}
	if global.IsEmpty() {
		return ErrGlobalNotExist
	}

	if !global.CanCommit() {
		return errors.New("global state error")
	}

	return c.handler(ctx, gid, func(b *entity.Branch) error {
		if b.IsSucceed() {
			return nil
		}
		transport, err := protocol.GetTransport(common.NetType(b.Protocol), b.Url)
		if err != nil {
			return err
		}
		// todo replace []byte(b.ReqData)
		req := common.NewReq([]byte(b.ReqData), nil)
		if _, err = transport.Request(ctx, req); err != nil {
			return err
		}
		return nil
	})
}

func (c *Coordinator) Rollback(ctx context.Context, gid string) error {
	global, err := c.dao.GetGlobal(ctx, gid)
	if err != nil {
		return err
	}
	if global.IsEmpty() {
		return ErrGlobalNotExist
	}
	if !global.CanRollback() {
		// todo error
		return errors.New("global state error")
	}

	return c.handler(ctx, gid, func(b *entity.Branch) error {
		return nil
	})
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
