package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/wuqinqiang/easycar/logging"

	"golang.org/x/sync/errgroup"
)

type Opt func(core *Core)

type Core struct {
	servers      []Server
	stopCtx      context.Context
	runWaitGroup sync.WaitGroup
	errGroup     *errgroup.Group
	cancel       func()
	once         sync.Once
}

func WithServers(srvs ...Server) Opt {
	return func(core *Core) {
		core.servers = append(core.servers, srvs...)
	}
}

func New(opts ...Opt) *Core {
	core := &Core{
		runWaitGroup: sync.WaitGroup{},
		once:         sync.Once{},
	}
	var (
		ctx context.Context
	)
	ctx, core.cancel = context.WithCancel(context.Background())
	core.errGroup, core.stopCtx = errgroup.WithContext(ctx)

	for _, opt := range opts {
		opt(core)
	}
	return core
}
func (core *Core) Run(ctx context.Context) error {
	for _, server := range core.servers {
		core.runWaitGroup.Add(1)
		srv := server
		core.errGroup.Go(func() error {
			<-core.stopCtx.Done()
			return srv.Stop(ctx)
		})

		core.errGroup.Go(func() error {
			defer core.runWaitGroup.Done()
			return srv.Run(ctx)
		})
	}
	// wait for all service is  start
	core.runWaitGroup.Wait()
	logging.Info("easycar start")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	core.errGroup.Go(func() error {
		select {
		case <-core.stopCtx.Done():
			return core.stopCtx.Err()
		case <-c:
			return core.Stop()
		}
	})
	if err := core.errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (core *Core) Stop() error {
	core.once.Do(func() {
		core.cancel()
	})
	return nil
}
