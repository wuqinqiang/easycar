package httpsrv

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/log"

	"github.com/wuqinqiang/easycar/tools"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/wuqinqiang/easycar/proto"
)

type HttpSrv struct {
	port       int
	grpcPort   int
	timeout    time.Duration
	httpServer *http.Server
	once       sync.Once
	//tls     *tls.Config
}

func New(port, grpcPort int, opts ...Opt) *HttpSrv {
	h := &HttpSrv{
		port:     port,
		grpcPort: grpcPort,
		once:     sync.Once{},
		timeout:  10 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (srv *HttpSrv) Run(ctx context.Context) (err error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%d", srv.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(color.HiRedString("grpc DialContext:err:%v", err))
		return err
	}

	gwmux := runtime.NewServeMux()
	if err = proto.RegisterEasyCarHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	srv.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", srv.port),
		Handler: gwmux,
	}
	tools.GoSafe(func() {
		if err = srv.httpServer.ListenAndServe(); err != nil {
			return
		}
	})
	log.Info(fmt.Sprintf("[HttpSrv] http port:%d", srv.port))
	return nil
}

func (srv *HttpSrv) Stop(ctx context.Context) (err error) {
	srv.once.Do(func() {
		err = srv.httpServer.Close()
	})
	log.Info("[HttpSrv]Stopped")
	return
}
