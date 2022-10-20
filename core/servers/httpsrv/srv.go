package httpsrv

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/tools"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/fatih/color"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/wuqinqiang/easycar/proto"
)

type HttpSrv struct {
	listenOn     string
	grpcListenOn string
	timeout      time.Duration
	httpServer   *http.Server
	once         sync.Once
	//tls     *tls.Config
}

func New(httpListenOn, grpcListenOn string, opts ...Opt) *HttpSrv {
	h := &HttpSrv{
		listenOn:     httpListenOn,
		grpcListenOn: grpcListenOn,
		once:         sync.Once{},
		timeout:      10 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (srv *HttpSrv) Run(ctx context.Context) (err error) {
	conn, err := grpc.DialContext(ctx, srv.grpcListenOn, grpc.WithBlock(),
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
		Addr:    srv.listenOn,
		Handler: gwmux,
	}
	tools.GoSafe(func() {
		if err = srv.httpServer.ListenAndServe(); err != nil {
			return
		}
	})
	logging.Info(fmt.Sprintf("[HttpSrv] http listen:%s", srv.listenOn))
	return nil
}

func (srv *HttpSrv) Stop(ctx context.Context) (err error) {
	srv.once.Do(func() {
		err = srv.httpServer.Close()
	})
	logging.Info("[HttpSrv]Stopped")
	return
}

func (srv *HttpSrv) Endpoint() *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   srv.listenOn,
	}
}
