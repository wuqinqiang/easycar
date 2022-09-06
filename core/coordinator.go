package core

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/executor"
	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/dao"
)

var (
	ErrAutomatic = fmt.Errorf("transaction phase2 has been processed automatically")
)

type Coordinator struct {
	dao                 dao.TransactionDao
	automaticExecution2 bool
}

func NewCoordinator(dao dao.TransactionDao, automaticExecution2 bool) *Coordinator {
	c := &Coordinator{
		dao:                 dao,
		automaticExecution2: automaticExecution2,
	}
	return c
}

func (c *Coordinator) Begin(ctx context.Context) (string, error) {
	gid := entity.GetGid()
	g := entity.NewGlobal(gid)
	g.SetState(consts.Init)
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
	logging.Infof("[Coordinator] phase1 end", "gid", global.GetGId(), "state", global.State)

	if err = c.UpdateGlobalState(ctx, global.GetGId(), phase1State); err != nil {
		return err
	}

	if c.automaticExecution2 {
		logging.Infof("[Coordinator] Phase2 start", "gid", global.GID)
		tools.GoSafe(func() {
			if err = c.Phase2(context.Background(), global, branches); err != nil {
				logging.Error(fmt.Sprintf("[Start] Phase2:err:%v", err))
				return
			}
		})
	}
	global.State = phase1State
	return nil
}

func (c *Coordinator) Phase1(ctx context.Context, branchList entity.BranchList) error {
	return executor.Phase1Executor(branchList).Execute(ctx)
}

func (c *Coordinator) Phase2(ctx context.Context, global *entity.Global, branches entity.BranchList) (err error) {
	var (
		processingStateVal, overStateVal interface{}
	)
	processingStateVal = tools.IF(global.Phase1Failed(), consts.Phase2Rollbacking, consts.Phase2Committing)
	if _, err = c.dao.UpdateGlobalStateByGid(ctx, global.GetGId(),
		processingStateVal.(consts.GlobalState)); err != nil {
		return
	}

	overStateVal = tools.IF(global.Phase1Failed(), consts.Rollbacked, consts.Committed)

	defer func() {
		if err != nil {
			overStateVal = tools.IF(global.Phase1Failed(), consts.Phase2RollbackFailed, consts.Phase2CommitFailed)
		}
		logging.Infof("[Coordinator] Phase2 end", "gid", global.GID, "state", overStateVal)
		_, erro := c.dao.UpdateGlobalStateByGid(ctx, global.GetGId(), overStateVal.(consts.GlobalState))
		if erro != nil {
			fmt.Printf("[Phase2]UpdateGlobalStateByGid gid:%v err:%v", global.GetGId(), erro)
		}
	}()
	err = executor.NewPhase2Executor(global, branches).Execute(ctx)
	return
}

func (c *Coordinator) Commit(ctx context.Context, global *entity.Global, branches entity.BranchList) error {
	if c.automaticExecution2 {
		return ErrAutomatic
	}
	return c.Phase2(ctx, global, branches)
}

func (c *Coordinator) Rollback(ctx context.Context, global *entity.Global, branches entity.BranchList) error {
	if c.automaticExecution2 {
		return ErrAutomatic
	}
	return c.Phase2(ctx, global, branches)
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

func (c *Coordinator) GetAutomaticExecution2() bool {
	return c.automaticExecution2
}
