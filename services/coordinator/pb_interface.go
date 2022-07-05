package coordinator

import (
	"context"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ proto.EasyCarServer = (*EasyCarSrv)(nil)

type EasyCarSrv struct {
	proto.UnimplementedEasyCarServer
}

func (e EasyCarSrv) Begin(ctx context.Context, empty *emptypb.Empty) (*proto.BeginResp, error) {
	resp := new(proto.BeginResp)
	resp.Data = new(proto.BeginRespInfo)
	resp.Common = new(proto.BaseResp)
	resp.Common.Code = 0
	resp.Common.Msg = "success"
	return resp, nil
}

func (e EasyCarSrv) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterResp, error) {
	//TODO implement me
	panic("implement me")
}

func (e EasyCarSrv) Commit(ctx context.Context, req *proto.CommitReq) (*proto.CommitResp, error) {
	//TODO implement me
	panic("implement me")
}

func (e EasyCarSrv) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	//TODO implement me
	panic("implement me")
}

func (e EasyCarSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	//TODO implement me
	panic("implement me")
}
