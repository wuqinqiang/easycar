package resolver

import (
	"context"
	"strings"
	"time"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	"google.golang.org/grpc/resolver"
)

type Builder struct {
	discoverer   registry.Discovery
	watchTimeout time.Duration
}

func Init() {
	resolver.Register(&Builder{})
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())
	watchCtx, watchCancel := context.WithTimeout(ctx, b.watchTimeout)
	defer watchCancel()
	var (
		err error
	)
	if b.discoverer == nil {
		if b.discoverer, err = registry.GetDiscovery(target.URL.Scheme); err != nil {
			cancel()
			return nil, err
		}
	}

	watch, err := b.discoverer.Watch(watchCtx, strings.TrimPrefix(target.URL.Path, "/"))
	if err != nil {
		cancel()
		return nil, err
	}
	r := NewDefaultResolver(ctx, cancel, cc, watch)
	tools.GoSafe(func() {
		r.watch()
	})
	return r, nil
}

func (b Builder) Scheme() string {
	return "easycar"
}
