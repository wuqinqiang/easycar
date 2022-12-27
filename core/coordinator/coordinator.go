package coordinator

import (
	"context"
	"fmt"
	"time"

	"github.com/wuqinqiang/easycar/core/notify"

	"github.com/wuqinqiang/easycar/logging"
	"github.com/wuqinqiang/easycar/tracing"

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
	executor            Executor
	closeFn             func(ctx context.Context) error
	notify              notify.Notify
}

func NewCoordinator(dao dao.TransactionDao, executor Executor, notify notify.Notify, automaticExecution2 bool) *Coordinator {
	c := &Coordinator{
		dao:                 dao,
		automaticExecution2: automaticExecution2,
		executor:            executor,
		closeFn:             executor.Close,
		notify:              notify,
	}
	return c
}

func (c *Coordinator) Begin(ctx context.Context) (string, error) {
	gid := entity.GetGid()
	g := entity.NewGlobal(gid)
	g.SetState(consts.Init)
	now := time.Now().Unix()
	g.CreateTime = now
	g.UpdateTime = now
	err := c.dao.CreateGlobal(ctx, g)
	return gid, err
}

func (c *Coordinator) Close(ctx context.Context) error {
	c.notify.Stop()
	return c.closeFn(ctx)
}

func (c *Coordinator) Register(ctx context.Context, branches entity.BranchList) error {
	if len(branches) == 0 {
		return nil
	}
	return c.dao.CreateBatches(ctx, branches)
}

func (c *Coordinator) Start(ctx context.Context, global *entity.Global) error {
	if err := c.Phase1(ctx, global); err != nil {
		return err
	}
	if c.automaticExecution2 {
		logging.Infof("[Coordinator] Phase2 start gid:%v", global.GID)
		tools.GoSafe(func() {
			if err := c.Phase2(context.Background(), global); err != nil {
				logging.Errorf("[Start] Phase2:err:%v", err)
				return
			}
		})
	}
	return nil
}

func (c *Coordinator) Phase1(ctx context.Context, global *entity.Global) (err error) {
	_, span := tracing.Tracer(ctx, "Phase1", "gid", global.GetGId())
	defer span.End()

	phase1State := consts.Phase1Success
	defer func() {
		if err != nil {
			phase1State = consts.Phase1Failed
			c.notify.Notify(notify.NewContext(global.GetGId(), err))
		}
		global.State = phase1State
		logging.Infof("[Coordinator] phase1 end gid:%v state:%v", global.GetGId(), global.State)
		if erro := c.UpdateGlobalState(ctx, global.GetGId(), phase1State); erro != nil {
			logging.Errorf("[Coordinator]Phase1 UpdateGlobalState:%v", erro)
		}
	}()
	err = c.executor.Phase1(ctx, global)
	return
}

func (c *Coordinator) Phase2(ctx context.Context, global *entity.Global) (err error) {

	isGotoRollback := global.GotoRollback()

	_, span := tracing.Tracer(ctx, "Phase2"+tools.IF(isGotoRollback, "Rollback", "Commit").(string),
		"gid", global.GetGId())
	defer span.End()

	var (
		processingStateVal, overStateVal interface{}
	)
	processingStateVal = tools.IF(isGotoRollback, consts.Phase2Rollbacking, consts.Phase2Committing)
	if _, err = c.dao.UpdateGlobalStateByGid(ctx, global.GetGId(),
		processingStateVal.(consts.GlobalState)); err != nil {
		return
	}

	overStateVal = tools.IF(isGotoRollback, consts.Rollbacked, consts.Committed)

	defer func() {
		if err != nil {
			overStateVal = tools.IF(isGotoRollback, consts.Phase2RollbackFailed, consts.Phase2CommitFailed)
			c.notify.Notify(notify.NewContext(global.GetGId(), err))
		}
		logging.Infof("[Coordinator] Phase2 end gid %v,state:%v", global.GID, overStateVal)
		_, erro := c.dao.UpdateGlobalStateByGid(ctx, global.GetGId(), overStateVal.(consts.GlobalState))
		if erro != nil {
			logging.Errorf("[Phase2]UpdateGlobalStateByGid gid:%v err:%v", global.GetGId(), erro)
		}
	}()
	err = c.executor.Phase2(ctx, global)
	return
}

func (c *Coordinator) Commit(ctx context.Context, global *entity.Global) error {
	if c.automaticExecution2 {
		logging.Warnf("[Commit] gid:%v,warn:%v", global.GetGId(), ErrAutomatic)
		return nil
	}
	return c.Phase2(ctx, global)
}

func (c *Coordinator) Rollback(ctx context.Context, global *entity.Global) error {
	if c.automaticExecution2 {
		logging.Warnf("[Rollback] gid:%v,warn:%v", global.GetGId(), ErrAutomatic)
		return nil
	}
	return c.Phase2(ctx, global)
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
