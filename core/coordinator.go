package core

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/executor"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/dao"
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
	g.SetState(consts.Ready)
	err := c.dao.CreateGlobal(ctx, g)
	return gid, err
}

func (c *Coordinator) Register(ctx context.Context, gId string, branches entity.BranchList) error {
	return c.dao.CreateBatches(ctx, branches)
}

func (c *Coordinator) Start(ctx context.Context, global *entity.Global, branches entity.BranchList) error {
	phase1State := consts.Phase1Success
	err := c.Phase1(ctx, branches)
	if err != nil {
		fmt.Printf("[Start] Phase1 err:%v\n", err)
		phase1State = consts.Phase1Failed
	}
	global.State = phase1State
	if err = c.UpdateGlobalState(ctx, global.GetGId(), phase1State); err != nil {
		return err
	}

	tools.GoSafe(func() {
		// todo replace ctx
		ctx2 := context.Background()
		if err = c.Phase2(ctx2, global, branches); err != nil {
			fmt.Printf("[Start] Phase2:err:%v", err)
			return
		}
	})
	global.State = phase1State
	return nil
}

func (c *Coordinator) Phase1(ctx context.Context, branchList entity.BranchList) error {
	return executor.Phase1Executor(branchList).Execute(ctx)
}

func (c *Coordinator) Phase2(ctx context.Context, global *entity.Global, branches entity.BranchList) error {
	return executor.NewPhase2Executor(global, branches).Execute(ctx)
}

func (c *Coordinator) GetGlobal(ctx context.Context, gid string) (entity.Global, error) {
	return c.dao.GetGlobal(ctx, gid)
}

func (c *Coordinator) GetBranchList(ctx context.Context, gid string) (list []*entity.Branch, err error) {
	return c.dao.GetBranches(ctx, gid)
}

func (c *Coordinator) UpdateGlobalState(ctx context.Context, gid string,
	state consts.GlobalState) error {
	_, err := c.dao.UpdateGlobalStateByGid(ctx, gid, state)
	return err
}
