package core

import (
	"context"

	"github.com/wuqinqiang/easycar/pkg/apis"
)

var _ apis.EasyCarServer = (*EasyCarGrpcHandler)(nil)

type EasyCarGrpcHandler struct {
	apis.UnimplementedEasyCarServer
}

func (e *EasyCarGrpcHandler) Begin(ctx context.Context, req *apis.BeginReq) (*apis.BeginResp, error) {
	panic("implement me")
}

func (e *EasyCarGrpcHandler) Commit(ctx context.Context, req *apis.CommitReq) (*apis.CommitResp, error) {
	panic("implement me")
}

func (e *EasyCarGrpcHandler) RollBack(ctx context.Context, req *apis.RollBackReq) (*apis.RollBackReq, error) {
	panic("implement me")
}
