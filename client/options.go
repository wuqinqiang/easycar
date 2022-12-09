package client

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
)

// DefaultOptions default for a Client Options
var DefaultOptions = &Options{
	connTimeout: 15 * time.Second,
}

type Option func(options *Options)

type Options struct {
	// connTimeout conn easycar timeout
	connTimeout time.Duration // tls
	tls         *tls.Config
	dailOpts    []grpc.DialOption
}

func WithConnTimeout(seconds time.Duration) Option {
	return func(options *Options) {
		if seconds > 0 {
			options.connTimeout = seconds
		}
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
