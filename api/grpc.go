package api

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ proto.EasyCarServer = (*Api)(nil)

type Api struct {
	proto.UnimplementedEasyCarServer
	grpcOpts   []grpc.ServerOption
	grpcServer *grpc.Server
	core       *core.Coordinator
}

func NewEasyCarApi(core *core.Coordinator) *Api {
	return &Api{
		core: core,
	}
}

func (api *Api) Run() {
	api.grpcServer = grpc.NewServer(api.grpcOpts...)
	proto.RegisterEasyCarServer(api.grpcServer, api)
}

func (api Api) check(ctx context.Context, gid string, fn func(g *entity.Global) error) (g entity.Global, err error) {
	g, err = api.core.GetGlobal(ctx, gid)
	if err != nil {
		return
	}
	if g.IsEmpty() {
		err = core.ErrGlobalNotExist
		return
	}
	if fn != nil {
		err = fn(&g)
	}
	return
}

func (api Api) Begin(ctx context.Context, empty *emptypb.Empty) (*proto.BeginResp, error) {
	gid, err := api.core.Begin(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(proto.BeginResp)
	resp.GId = gid
	return resp, nil
}

func (api Api) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterResp, error) {
	var (
		list entity.BranchList
	)
	list = list.AssignmentByGrpc(req.GetGId(), req.Branches)
	if err := api.core.Register(ctx, req.GetGId(), list); err != nil {
		return nil, err
	}
	resp := new(proto.RegisterResp)
	return resp, nil
}

func (api Api) Start(ctx context.Context, req *proto.StartReq) (*proto.StartResp, error) {
	global, err := api.check(ctx, req.GetGId(), func(g *entity.Global) error {
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
	if branches, err = api.core.GetBranchList(ctx, global.GetGId()); err != nil {
		return nil, err
	}
	global.SetState(consts.Phase1Processing)
	if err = api.core.UpdateGlobalState(ctx, global.GetGId(), global.State); err != nil {
		return nil, err
	}

	if err = api.core.Start(ctx, &global, branches); err != nil {
		return nil, err
	}
	resp := new(proto.StartResp)
	return resp, nil
}

//todo remove
func (api Api) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	return nil, nil
}

func (api Api) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	global, err := api.core.GetGlobal(ctx, req.GetGId())
	if err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.State = consts.ConvertStateToGrpc(global.GetState())
	return resp, nil
}
