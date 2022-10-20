package grpcsrv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/core/coordinator"

	elog "github.com/wuqinqiang/easycar/logging"

	"google.golang.org/grpc/reflection"

	"github.com/wuqinqiang/easycar/proto"

	"google.golang.org/grpc"
)

var (
	_              proto.EasyCarServer = (*GrpcSrv)(nil)
	ErrGIdNotExist                     = errors.New("gid is not exist")
)

type GrpcSrv struct {
	proto.UnimplementedEasyCarServer
	coordinator *coordinator.Coordinator

	lis net.Listener

	timeout    time.Duration
	listenOn   string
	once       sync.Once
	grpcOpts   []grpc.ServerOption
	grpcServer *grpc.Server
}

func New(listenOn string, coordinator *coordinator.Coordinator, opts ...Opt) (*GrpcSrv, error) {
	srv := &GrpcSrv{
		coordinator: coordinator,
		timeout:     10 * time.Second,
		listenOn:    listenOn,
		once:        sync.Once{},
	}

	for _, opt := range opts {
		opt(srv)
	}
	maxSize := 5 * 1024 * 1024 //5M:max Recv msg size
	srv.grpcOpts = append(srv.grpcOpts, grpc.MaxRecvMsgSize(maxSize))

	var (
		err error
	)
	srv.lis, err = net.Listen("tcp", listenOn)
	return srv, err
}

func (s *GrpcSrv) Run(ctx context.Context) error {
	s.grpcServer = grpc.NewServer(s.grpcOpts...)

	proto.RegisterEasyCarServer(s.grpcServer, s)
	// for reflection
	reflection.Register(s.grpcServer)
	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			log.Fatal(err)
		}
	}()
	elog.Info(fmt.Sprintf("[Grpc] grpc listen:%s", s.listenOn))
	return nil
}

func (s *GrpcSrv) Stop(ctx context.Context) (err error) {
	s.once.Do(func() {
		s.grpcServer.GracefulStop()
	})
	if err = s.coordinator.Close(ctx); err != nil {
		return
	}
	elog.Info("[GrpcSrv] stopped")
	return
}

func (s *GrpcSrv) Endpoint() *url.URL {
	return &url.URL{
		Scheme: "grpc",
		Host:   s.listenOn,
	}
}
