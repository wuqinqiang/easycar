package core

import (
	"time"

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
}

func WithPort(port int) OptsFn {
	return func(opts *opts) {
		opts.port = port
	}
}

func WithConf(conf *conf.EasyCar) OptsFn {
	return func(opts *opts) {
		opts.conf = conf
	}
}
