package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/wuqinqiang/easycar/core/endponit"

	"github.com/wuqinqiang/easycar/core/registry"

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
	registry     registry.Registry
	instance     *registry.EasyCarInstance
}

func WithServers(srvs ...Server) Opt {
	return func(core *Core) {
		core.servers = append(core.servers, srvs...)
	}
}

func WithRegistry(r registry.Registry) Opt {
	return func(core *Core) {
		core.registry = r
	}
}

func New(opts ...Opt) *Core {
	core := &Core{
		runWaitGroup: sync.WaitGroup{},
		once:         sync.Once{},
	}
	for _, opt := range opts {
		opt(core)
	}
	return core
}
func (core *Core) Run(ctx context.Context) error {
	var (
		c1 context.Context
	)
	c1, core.cancel = context.WithCancel(ctx)
	core.errGroup, core.stopCtx = errgroup.WithContext(c1)

	core.instance = registry.NewEasyCarInstance()

	for _, server := range core.servers {
		core.runWaitGroup.Add(1)

		if e, ok := server.(endponit.Endpoint); ok {
			core.instance.Node = append(core.instance.Node, e.Endpoint().String())
		}
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

	if core.registry != nil {
		if err := core.registry.Register(c1, core.instance); err != nil {
			return err
		}
	}

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

func (core *Core) Stop() (err error) {
	if core.cancel == nil {
		return nil
	}
	core.once.Do(func() {
		if core.registry != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err = core.registry.DeRegister(ctx, core.instance)
		}
		core.cancel()
	})
	return
}
