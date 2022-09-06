package grpcsrv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
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
	port       int
	once       sync.Once
	grpcOpts   []grpc.ServerOption
	grpcServer *grpc.Server
}

func New(port int, coordinator *coordinator.Coordinator, opts ...Opt) (*GrpcSrv, error) {
	srv := &GrpcSrv{
		coordinator: coordinator,
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
	maxSize := 5 * 1024 * 1024 //5M:max Recv msg size
	s.grpcServer = grpc.NewServer(grpc.MaxRecvMsgSize(maxSize))
	proto.RegisterEasyCarServer(s.grpcServer, s)
	// for reflection
	reflection.Register(s.grpcServer)
	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			log.Fatal(err)
		}
	}()
	elog.Info(fmt.Sprintf("[Grpc] grpc port:%d", s.port))
	return nil
}

func (s *GrpcSrv) Stop(ctx context.Context) (err error) {
	s.once.Do(func() {
		err = s.lis.Close()
	})
	elog.Info("grpc stopped")
	return
}
