package grpcsrv

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

const DefaultListenOn = "127.0.0.1:8089"

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

func New(settings Grpc, coordinator *coordinator.Coordinator) (*GrpcSrv, error) {
	listenOn := DefaultListenOn
	if settings.ListenOn != "" {
		listenOn = settings.ListenOn
	}
	listenOn = tools.FigureOutListen(listenOn)
	srv := &GrpcSrv{
		coordinator: coordinator,
		timeout:     10 * time.Second,
		listenOn:    listenOn,
		once:        sync.Once{},
	}
	// setup tls
	var (
		err error
	)
	if settings.Tls() {
		certificate, err := tls.LoadX509KeyPair(settings.CertFile, settings.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
		tlsConf := &tls.Config{Certificates: []tls.Certificate{certificate}}
		srv.grpcOpts = append(srv.grpcOpts, grpc.Creds(credentials.NewTLS(tlsConf)))
	}
	srv.grpcOpts = append(srv.grpcOpts, grpc.MaxRecvMsgSize(5*1024*1024)) //5M:max Recv msg size
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

func (s *GrpcSrv) Handler(certFile, name string) httpsrv.Handler {
	return func(ctx context.Context) (http.Handler, error) {
		options := []grpc.DialOption{
			grpc.WithBlock(),
		}
		var (
			err error
		)
		creds := insecure.NewCredentials()

		if certFile != "" {
			creds, err = credentials.NewClientTLSFromFile(certFile, name)
			if err != nil {
				return nil, err
			}
		}
		options = append(options, grpc.WithTransportCredentials(creds))
		conn, err := grpc.DialContext(ctx, s.listenOn, options...) //todo replace
		if err != nil {
			fmt.Println(color.HiRedString("grpc DialContext:err:%v", err))
			return nil, err
		}
		mux := runtime.NewServeMux()
		err = proto.RegisterEasyCarHandler(ctx, mux, conn)
		return mux, err
	}
}

func (s *GrpcSrv) Stop(ctx context.Context) (err error) {
	if s.grpcServer == nil {
		return
	}
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
