package resolver

import (
	"context"
	"fmt"

	"github.com/wuqinqiang/easycar/core/endponit"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	"github.com/wuqinqiang/easycar/logging"
	"google.golang.org/grpc/resolver"
)

type defaultResolver struct {
	cc      resolver.ClientConn
	watcher registry.Watcher
	ctx     context.Context
	cancel  func()
}

func NewDefaultResolver(ctx context.Context, cancel func(), cc resolver.ClientConn, w registry.Watcher) *defaultResolver {
	r := &defaultResolver{
		cc:      cc,
		watcher: w,
		ctx:     ctx,
		cancel:  cancel,
	}
	tools.GoSafe(func() {
		r.watch()
	})
	return r
}

func (r *defaultResolver) watch() {
	for {
		select {
		case <-r.ctx.Done():
			return
		default:
		}
		instances, err := r.watcher.Next()
		if err != nil {
			return
		}
		r.updateState(instances)
	}
}

func (r *defaultResolver) updateState(list []*registry.EasyCarInstance) {

	var (
		state resolver.State
	)

	for _, instance := range list {
		e, err := endponit.GetHostByEndpoint(instance.Node, "grpc")
		if err != nil {
			logging.Error(fmt.Sprintf("[updateState]GetHostByEndpoint err:%v", err))
			continue
		}
		if e == "" {
			continue
		}

		state.Addresses = append(state.Addresses, resolver.Address{
			Addr:       e,
			ServerName: instance.Name,
			//Attributes:         nil,
			//BalancerAttributes: nil,
		})
	}
	if len(state.Addresses) == 0 {
		return
	}

	err := r.cc.UpdateState(state)
	if err != nil {
		logging.Error(fmt.Sprintf("[updateState]UpdateState err:%v", err))
		return
	}
}

func (r *defaultResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r *defaultResolver) Close() {
	r.cancel()
	if err := r.watcher.Stop(); err != nil {
		logging.Error(fmt.Sprintf("defaultResolver close:%v", err))
	}
}