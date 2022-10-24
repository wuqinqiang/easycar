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
)

type HandleOptions struct {
	certFile string
	certName string
}

func WithCertFile(certFile string) OptsFn {
	return func(options *HandleOptions) {
		options.certFile = certFile
	}
}

func WithCertName(name string) OptsFn {
	return func(options *HandleOptions) {
		options.certName = name
	}
}

type OptsFn func(options *HandleOptions)

type HandlerFn func(ctx context.Context, fns ...OptsFn) (http.Handler, error)

type HttpSrv struct {
	listenOn   string
	fn         HandlerFn
	timeout    time.Duration
	httpServer *http.Server
	once       sync.Once
	//tls     *tls.Config
}

func New(httpListenOn string, fn HandlerFn, opts ...Opt) *HttpSrv {
	h := &HttpSrv{
		fn:       fn,
		listenOn: httpListenOn,
		once:     sync.Once{},
		timeout:  10 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (srv *HttpSrv) Run(ctx context.Context) error {
	hctx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	h, err := srv.fn(hctx)
	if err != nil {
		return err
	}
	srv.httpServer = &http.Server{
		Addr:    srv.listenOn,
		Handler: h,
	}
	tools.GoSafe(func() {
		if err = srv.httpServer.ListenAndServe(); err != nil {
			logging.Error(fmt.Sprintf("[HttpSrv ]ListenAndServe err:%v", err))
		}
	})
	logging.Info(fmt.Sprintf("[HttpSrv] http listen:%s", srv.listenOn))
	return nil
}

func (srv *HttpSrv) Stop(ctx context.Context) (err error) {
	if srv.httpServer == nil {
		return
	}
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
