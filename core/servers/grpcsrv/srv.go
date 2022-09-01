package grpcsrv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/fatih/color"
	"google.golang.org/grpc/reflection"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/proto"

	"google.golang.org/grpc"
)

var (
	_              proto.EasyCarServer = (*GrpcSrv)(nil)
	ErrGIdNotExist                     = errors.New("gid is not exist")
)

type GrpcSrv struct {
	proto.UnimplementedEasyCarServer
	coordinator *core.Coordinator

	lis net.Listener

	timeout    time.Duration
	port       int
	once       sync.Once
	grpcOpts   []grpc.ServerOption
	grpcServer *grpc.Server
}

func New(port int, opts ...Opt) (*GrpcSrv, error) {
	srv := &GrpcSrv{
		coordinator: core.NewCoordinator(dao.GetTransaction()),
		timeout:     10 * time.Second,
		port:        port,
		once:        sync.Once{},
	}

	for _, opt := range opts {
		opt(srv)
	}

	var (
		err error
	)
	srv.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	return srv, err
}

func (s *GrpcSrv) Run(ctx context.Context) error {
	s.grpcServer = grpc.NewServer()
	proto.RegisterEasyCarServer(s.grpcServer, s)
	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			log.Fatal(err)
		}
	}()
	// for reflection
	reflection.Register(s.grpcServer)
	fmt.Println(color.BlueString("easycar grpc port:%d", s.port))
	return nil
}

func (s *GrpcSrv) Stop(ctx context.Context) (err error) {
	s.once.Do(func() {
		err = s.lis.Close()
	})
	fmt.Println("grpc over")
	return
}

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

	global, err := s.check(ctx, req.GetGId(), func(g *entity.Global) error {
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

func (s *GrpcSrv) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	return nil, nil
}

func (s *GrpcSrv) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	global, err := s.coordinator.GetGlobal(ctx, req.GetGId())
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
