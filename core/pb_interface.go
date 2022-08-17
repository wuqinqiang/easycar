package core

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ proto.EasyCarServer = (*EasyCarSrv)(nil)

type EasyCarSrv struct {
	proto.UnimplementedEasyCarServer
	core *Coordinator
}

func NewEasyCarSrv(core *Coordinator) *EasyCarSrv {
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
		if !g.IsReady() {
			return fmt.Errorf("global state:%v can not start", g.GetState())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var (
		branches entity.BranchList
	)
	if branches, err = e.core.GetBranchList(ctx, global.GetGId()); err != nil {
		return nil, err
	}
	global.SetState(consts.Phase1Processing)
	if err = e.core.UpdateGlobalState(ctx, global.GetGId(), global.State); err != nil {
		return nil, err
	}

	if err = e.core.Start(ctx, &global, branches); err != nil {
		return nil, err
	}
	resp := new(proto.StartResp)
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
