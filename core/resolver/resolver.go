package resolver

import (
	"context"

	"github.com/wuqinqiang/easycar/core/endponit"

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

func NewDefaultResolver(ctx context.Context, cc resolver.ClientConn, w registry.Watcher) *defaultResolver {
	r := &defaultResolver{
		cc:      cc,
		watcher: w,
	}
	r.ctx, r.cancel = context.WithCancel(ctx)
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
			logging.Errorf("[defaultResolver] watch err:%v", err)
			return
		}
		r.updateState(instances)
	}
}

func (r *defaultResolver) updateState(list []*registry.Instance) {

	var (
		state resolver.State
	)
	logging.Infof("[defaultResolver]updateState:%v", list)

	for _, instance := range list {
		e, err := endponit.GetHostByEndpoint(instance.Nodes, "grpc")
		if err != nil {
			logging.Errorf("[updateState]GetHostByEndpoint err:%v", err)
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
		logging.Errorf("[updateState]UpdateState err:%v", err)
		return
	}
}

func (r *defaultResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (r *defaultResolver) Close() {
	r.cancel()
	if err := r.watcher.Stop(); err != nil {
		logging.Errorf("defaultResolver close:%v", err)
	}
}
