package grpcsrv

import (
	"context"
	"errors"
	"fmt"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GrpcSrv) Begin(ctx context.Context, empty *emptypb.Empty) (*proto.BeginResp, error) {
	gid, err := s.coordinator.Begin(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(proto.BeginResp)
	resp.GId = gid
	return resp, nil
}

func (s *GrpcSrv) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.AllowRegister() {
			return errors.New("register not allowed")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	//todo  check register branches are match?

	branchList := entity.GetBranchList(req.GetGId(), req.Branches)
	if err := s.coordinator.Register(ctx, req.GetGId(), branchList); err != nil {
		return nil, err
	}
	resp := new(proto.RegisterResp)
	return resp, nil
}

func (s *GrpcSrv) Start(ctx context.Context, req *proto.StartReq) (*proto.StartResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var (
		global entity.Global
		err    error
	)

	if global, err = s.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.Init() {
			return fmt.Errorf("global state:%v can not start", g.GetState())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	var (
		branches entity.BranchList
	)
	if branches, err = s.coordinator.GetBranchList(ctx, global.GetGId()); err != nil {
		return nil, err
	}
	global.SetState(consts.Phase1Processing)
	if err = s.coordinator.UpdateGlobalState(ctx, global.GetGId(), global.State); err != nil {
		return nil, err
	}

	if err = s.coordinator.Start(ctx, &global, branches); err != nil {
		return nil, err
	}
	resp := new(proto.StartResp)
	return resp, nil
}

func (s *GrpcSrv) commonPhase2(ctx context.Context, gid string,
	filterFn func(g *entity.Global) error, executeFn func(context.Context, *entity.Global, entity.BranchList) error) (err error) {
	var (
		global entity.Global
	)
	if global, err = s.check(ctx, gid, filterFn); err != nil {
		return
	}
	var (
		branches entity.BranchList
	)
	if branches, err = s.coordinator.GetBranchList(ctx, global.GetGId()); err != nil {
		return
	}
	return executeFn(ctx, &global, branches)

}

func (s *GrpcSrv) Commit(ctx context.Context, req *proto.CommitReq) (*proto.CommitResp, error) {
	err := s.commonPhase2(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.IsPhase1Success() {
			return fmt.Errorf("gid:%v can not commit", req.GetGId())
		}
		return nil
	}, s.coordinator.Commit)
	if err != nil {
		return nil, err
	}
	resp := new(proto.CommitResp)
	return resp, nil
}

func (s *GrpcSrv) Rollback(ctx context.Context, req *proto.RollBckReq) (*proto.RollBckResp, error) {
	err := s.commonPhase2(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.IsPhase2Failed() {
			return fmt.Errorf("gid:%v can not bollback", req.GetGId())
		}
		return nil
	}, s.coordinator.Rollback)
	if err != nil {
		return nil, err
	}
	resp := new(proto.RollBckResp)
	return resp, nil
}

func (s *GrpcSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	global, err := s.check(ctx, req.GetGId(), nil)
	if err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.State = consts.ConvertStateToGrpc(global.GetState())
	return resp, nil
}

func (s *GrpcSrv) check(ctx context.Context, gid string, fn func(g *entity.Global) error) (g entity.Global, err error) {
	g, err = s.coordinator.GetGlobal(ctx, gid)
	if err != nil {
		return
	}
	if g.IsEmpty() {
		err = ErrGIdNotExist
		return
	}
	if fn != nil {
		err = fn(&g)
	}
	return
}
