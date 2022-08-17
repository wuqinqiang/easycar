package core

import (
	"time"
)

var defaultOpts = opts{
	timeout: 3 * time.Second,
}

type (
	OptsFn func(opts *opts)
)

type opts struct {
	port    int
	timeout time.Duration
}

func WithPort(port int) OptsFn {
	return func(opts *opts) {
		opts.port = port
	}
}
