package core

import (
	"time"

	"google.golang.org/grpc"

	"github.com/wuqinqiang/easycar/conf"
)

var defaultOpts = opts{
	timeout:     3 * time.Second,
	grpcTimeout: 5 * time.Second,
}

type (
	OptsFn func(opts *opts)
)

type opts struct {
	port        int
	grpcTimeout time.Duration
	timeout     time.Duration
	conf        *conf.EasyCar

	grpcOpts []grpc.ServerOption
}

func WithPort(port int) OptsFn {
	return func(opts *opts) {
		opts.port = port
	}
}

func WithGrpcOpts(grpcOpts ...grpc.ServerOption) OptsFn {
	return func(opts *opts) {
		opts.grpcOpts = grpcOpts
	}
}

func WithConf(conf *conf.EasyCar) OptsFn {
	return func(opts *opts) {
		opts.conf = conf
	}
}
