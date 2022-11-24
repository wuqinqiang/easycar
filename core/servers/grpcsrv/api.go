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

func (s *GrpcSrv) Register(ctx context.Context, req *proto.RegisterReq) (*emptypb.Empty, error) {
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
	if err := s.coordinator.Register(ctx, branchList); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *GrpcSrv) Start(ctx context.Context, req *proto.StartReq) (*emptypb.Empty, error) {
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
	global.SetState(consts.Phase1Preparing)
	if err = s.coordinator.UpdateGlobalState(ctx, global.GetGId(), global.State); err != nil {
		return nil, err
	}

	if err = s.coordinator.Start(ctx, &global); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *GrpcSrv) commonPhase2(ctx context.Context, gid string,
	filterFn func(g *entity.Global) error, executeFn func(context.Context, *entity.Global) error) (err error) {
	var (
		global entity.Global
	)
	if global, err = s.check(ctx, gid, filterFn); err != nil {
		return
	}
	return executeFn(ctx, &global)

}

func (s *GrpcSrv) Commit(ctx context.Context, req *proto.CommitReq) (*emptypb.Empty, error) {
	err := s.commonPhase2(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.GotoCommit() {
			return fmt.Errorf("gid:%v can not commit", req.GetGId())
		}
		return nil
	}, s.coordinator.Commit)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *GrpcSrv) Rollback(ctx context.Context, req *proto.RollBckReq) (*emptypb.Empty, error) {
	err := s.commonPhase2(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.GotoRollback() {
			return fmt.Errorf("gid:%v can not bollback", req.GetGId())
		}
		return nil
	}, s.coordinator.Rollback)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *GrpcSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	global, err := s.check(ctx, req.GetGId(), nil)
	if err != nil {
		return nil, err
	}
	var (
		branches entity.BranchList
	)
	if branches, err = s.coordinator.GetBranchList(ctx, global.GetGId()); err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.GId = global.GetGId()
	resp.EndTime = global.GetEndTime()
	resp.State = consts.ConvertStateToGrpc(global.GetState())

	for i := range branches {
		resp.Branches = append(resp.Branches, &proto.GetStateRespBranch{
			BranchId:   branches[i].BranchId,
			ReqData:    branches[i].ReqData,
			ReqHeader:  branches[i].ReqHeader,
			Uri:        branches[i].Url,
			TranType:   consts.ConvertTranTypeToGrpc(branches[i].TranType),
			Protocol:   branches[i].Protocol,
			Action:     consts.ConvertBranchActionToGrpc(branches[i].Action),
			State:      consts.ConvertBranchStateToGrpc(branches[i].State),
			Level:      int64(branches[i].Level),
			LastErrMsg: branches[i].LastErrMsg,
		})
	}

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
