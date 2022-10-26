package httpsrv

import (
	"context"
	"net/http"
)

type Option func(srv *HttpSrv)

type HandleOptions struct {
	certFile string
	certName string
}

func WithCertFile(certFile string) OptsFn {
	return func(options *HandleOptions) {
		options.certFile = certFile
	}
}

func WithCertName(name string) OptsFn {
	return func(options *HandleOptions) {
		options.certName = name
	}
}

type OptsFn func(options *HandleOptions)

type HandlerFn func(ctx context.Context, fns ...OptsFn) (http.Handler, error)
