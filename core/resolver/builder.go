package resolver

import (
	"context"
	"time"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	"google.golang.org/grpc/resolver"
)

type Builder struct {
	discoverer   registry.Discovery
	watchTimeout time.Duration
}

func Init(d registry.Discovery) {
	resolver.Register(newBuilder(d))
}

func newBuilder(d registry.Discovery) *Builder {
	return &Builder{
		discoverer:   d,
		watchTimeout: time.Second * 8,
	}
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx := context.Background()
	watch, err := b.discoverer.Watch(ctx, target.URL.Path)
	if err != nil {
		return nil, err
	}
	r := NewDefaultResolver(ctx, cc, watch)
	tools.GoSafe(func() {
		r.watch()
	})
	return r, nil
}

func (b Builder) Scheme() string {
	return "discovery"
}
