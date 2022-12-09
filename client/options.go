package client

import (
	"crypto/tls"
	"time"

	"github.com/wuqinqiang/easycar/core/balancer"
	"github.com/wuqinqiang/easycar/core/registry"
	"google.golang.org/grpc"
)

// DefaultOptions default for a Client Options
var DefaultOptions = &Options{
	connTimeout: 15 * time.Second,
}

type Option func(options *Options)

type Options struct {
	// connTimeout conn easycar timeout
	connTimeout time.Duration
	// discovery service discovery
	discovery registry.Discovery
	// tls
	tls *tls.Config
	// tactics of balancer
	tactics  balancer.TacticsName
	dailOpts []grpc.DialOption
}

func WithConnTimeout(seconds time.Duration) Option {
	return func(options *Options) {
		if seconds > 0 {
			options.connTimeout = seconds
		}
	}
}

func WithTactics(name balancer.TacticsName) Option {
	return func(options *Options) {
		options.tactics = name
	}
}

func WithGrpcDailOpts(opts []grpc.DialOption) Option {
	return func(options *Options) {
		options.dailOpts = append(options.dailOpts, opts...)
	}
}

func WithTls(tls *tls.Config) Option {
	return func(options *Options) {
		options.tls = tls
	}
}

func WithDiscovery(discovery registry.Discovery) Option {
	return func(options *Options) {
		options.discovery = discovery
	}
}
