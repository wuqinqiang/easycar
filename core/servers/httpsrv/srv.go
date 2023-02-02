package httpsrv

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/tools"
)

type HttpSrv struct {
	listenOn   string
	fn         Handler
	timeout    time.Duration
	httpServer *http.Server
	once       sync.Once
	//tls     *tls.Config
}

func New(conf Http, fn Handler) *HttpSrv {
	h := &HttpSrv{
		fn:       fn,
		listenOn: tools.FigureOutListen(conf.ListenOn),
		once:     sync.Once{},
		timeout:  10 * time.Second,
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
			logging.Errorf("[HttpSrv ]ListenAndServe err:%v", err)
		}
	})
	logging.Infof("[HttpSrv] http listen:%s", srv.listenOn)
	return nil
}

func (srv *HttpSrv) Stop(ctx context.Context) (err error) {
	if srv.httpServer == nil {
		return
	}
	srv.once.Do(func() {
		err = srv.httpServer.Close()
	})
	logging.Infof("[HttpSrv]Stopped")
	return
}

func (srv *HttpSrv) Endpoint() *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   srv.listenOn,
	}
}
