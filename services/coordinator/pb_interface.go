package coordinator

import (
	"context"
	"errors"
	"fmt"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrGlobalNotExist = errors.New("global not exist")
)

var _ proto.EasyCarServer = (*EasyCarSrv)(nil)

type EasyCarSrv struct {
	proto.UnimplementedEasyCarServer
	core *core.Coordinator
}

func NewCoordinator(core *core.Coordinator) *EasyCarSrv {
	return &EasyCarSrv{
		core: core,
	}
}

func (e EasyCarSrv) check(ctx context.Context, gid string, fn func(g *entity.Global) error) (g entity.Global, err error) {
	g, err = e.core.GetGlobal(ctx, gid)
	if err != nil {
		return
	}
	if g.IsEmpty() {
		err = ErrGlobalNotExist
		return
	}
	if fn != nil {
		err = fn(&g)
	}
	return
}

func (e EasyCarSrv) Begin(ctx context.Context, empty *emptypb.Empty) (*proto.BeginResp, error) {
	gid, err := e.core.Begin(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(proto.BeginResp)
	resp.GId = gid
	return resp, nil
}

func (e EasyCarSrv) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterResp, error) {
	var (
		list entity.BranchList
	)
	list = list.AssignmentByGrpc(req.GetGId(), req.Branches)
	if err := e.core.Register(ctx, req.GetGId(), list); err != nil {
		return nil, err
	}
	resp := new(proto.RegisterResp)
	return resp, nil
}

func (e EasyCarSrv) Start(ctx context.Context, req *proto.StartReq) (*proto.StartResp, error) {
	global, err := e.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.CanCommit() {
			return fmt.Errorf("global state:%v can not commit", g.GetState())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if !global.IsBegin() {
		return nil, fmt.Errorf("global state:%v can not commit", global.GetState())
	}

	if err := e.core.Commit(ctx, global); err != nil {
		return nil, err
	}
	resp := new(proto.StartResp)
	resp.State = consts.ConvertStateToGrpc(consts.Phase1Success)
	return resp, nil
}

func (e EasyCarSrv) Commit(ctx context.Context, req *proto.CommitReq) (*proto.CommitResp, error) {
	global, err := e.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.CanCommit() {
			return fmt.Errorf("global state:%v can not commit", g.GetState())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if err := e.core.Commit(ctx, global); err != nil {
		return nil, err
	}
	resp := new(proto.CommitResp)
	resp.State = consts.ConvertStateToGrpc(consts.Phase1Success)
	return resp, nil
}

func (e EasyCarSrv) RollBack(ctx context.Context, req *proto.RollBackReq) (*proto.RollBackResp, error) {
	global, err := e.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.CanPhase2() {
			return fmt.Errorf("global state:%v can not rollback", g.GetState())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err := e.core.Rollback(ctx, global); err != nil {
		return nil, err
	}
	resp := new(proto.RollBackResp)
	resp.State = consts.ConvertStateToGrpc(consts.Phase2Success)
	return resp, nil
}

//todo remove
func (e EasyCarSrv) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	return nil, nil
}

func (e EasyCarSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	global, err := e.core.GetGlobal(ctx, req.GetGId())
	if err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.State = consts.ConvertStateToGrpc(global.GetState())
	return resp, nil
}
