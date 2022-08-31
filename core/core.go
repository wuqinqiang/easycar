package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/entity"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/fatih/color"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
)

var (
	ErrGIdNotExist = errors.New("gid is not exist")
)

var _ proto.EasyCarServer = (*Core)(nil)

type Core struct {
	proto.UnimplementedEasyCarServer
	grpcServer *grpc.Server

	opts        opts
	lis         net.Listener
	once        sync.Once
	coordinator *Coordinator
}

func New(fns ...OptsFn) (s *Core, err error) {
	opts := defaultOpts
	for _, fn := range fns {
		fn(&opts)
	}
	s = &Core{
		opts:        opts,
		once:        sync.Once{},
		coordinator: NewCoordinator(dao.GetTransaction()),
	}
	if s.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", opts.grpcPort)); err != nil {
		return
	}
	return
}
func (core *Core) Run() error {
	// todo add opts
	core.grpcServer = grpc.NewServer()
	proto.RegisterEasyCarServer(core.grpcServer, core)
	go func() {
		if err := core.grpcServer.Serve(core.lis); err != nil {
			log.Fatal(err)
		}
	}()

	reflection.Register(core.grpcServer)

	// http service
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf(":%d", core.opts.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(color.HiRedString("grpc DialContext:err:%v", err))
		return err
	}

	fmt.Println(color.BlueString("easycar grpc port:%d", core.opts.grpcPort))

	gwmux := runtime.NewServeMux()
	if err = proto.RegisterEasyCarHandler(context.Background(), gwmux, conn); err != nil {
		return err
	}
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", core.opts.httpPort),
		Handler: gwmux,
	}
	fmt.Println(color.BlueString("easycar http port:%d", core.opts.httpPort))
	return gwServer.ListenAndServe()
}

func (core *Core) Stop() (err error) {
	core.once.Do(func() {
		err = core.lis.Close()
	})
	return nil
}

func (core *Core) Begin(ctx context.Context, empty *emptypb.Empty) (*proto.BeginResp, error) {
	gid, err := core.coordinator.Begin(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(proto.BeginResp)
	resp.GId = gid
	return resp, nil
}

func (core *Core) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := core.check(ctx, req.GetGId(), func(g *entity.Global) error {
		if !g.AllowRegister() {
			return errors.New("register not allowed")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	branchList := entity.GetBranchList(req.GetGId(), req.Branches)
	if err := core.coordinator.Register(ctx, req.GetGId(), branchList); err != nil {
		return nil, err
	}
	resp := new(proto.RegisterResp)
	return resp, nil
}

func (core *Core) Start(ctx context.Context, req *proto.StartReq) (*proto.StartResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	global, err := core.check(ctx, req.GetGId(), func(g *entity.Global) error {
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
	if branches, err = core.coordinator.GetBranchList(ctx, global.GetGId()); err != nil {
		return nil, err
	}
	global.SetState(consts.Phase1Processing)
	if err = core.coordinator.UpdateGlobalState(ctx, global.GetGId(), global.State); err != nil {
		return nil, err
	}

	if err = core.coordinator.Start(ctx, &global, branches); err != nil {
		return nil, err
	}
	resp := new(proto.StartResp)
	return resp, nil
}

func (core *Core) Abort(ctx context.Context, req *proto.AbortReq) (*proto.AbortResp, error) {
	return nil, nil
}

func (core *Core) GetState(ctx context.Context, req *proto.GetStateReq) (*proto.GetStateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	global, err := core.coordinator.GetGlobal(ctx, req.GetGId())
	if err != nil {
		return nil, err
	}
	resp := new(proto.GetStateResp)
	resp.State = consts.ConvertStateToGrpc(global.GetState())
	return resp, nil
}

func (core *Core) check(ctx context.Context, gid string, fn func(g *entity.Global) error) (g entity.Global, err error) {
	g, err = core.coordinator.GetGlobal(ctx, gid)
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
