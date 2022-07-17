package coordinator

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
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
	list = list.Assign2(req.Branches)
	if err := e.core.Register(ctx, req.GetGId(), list); err != nil {
		return nil, err
	}
	resp := new(proto.RegisterResp)
	return resp, nil
}

func (e EasyCarSrv) Commit(ctx context.Context, req *proto.CommitReq) (*proto.CommitResp, error) {
	if err := e.core.Commit(ctx, req.GetGId()); err != nil {
		return nil, err
	}
	resp := new(proto.CommitResp)
	return resp, nil
}

func (e EasyCarSrv) RollBack(ctx context.Context, req *proto.RollBackReq) (*proto.RollBackResp, error) {
	if err := e.core.Rollback(ctx, req.GetGId()); err != nil {
		return nil, err
	}
	resp := new(proto.RollBackResp)
	return resp, nil
}

func (e EasyCarSrv) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	return nil, nil

}

func (e EasyCarSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	global, err := e.core.GetState(ctx, req.GetGId())
	if err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.State = consts.ConvertStateToGrpc(global.GetState())
	return resp, nil
}
