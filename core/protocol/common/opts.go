package common

import "time"

var DefaultOps = Opts{timeOut: 5 * time.Second}

type (
	Req struct {
		Body    []byte
		Headers map[string]string
		Opts    *Opts
	}

	Opts struct {
		timeOut time.Duration
	}
	OptsFn func(*Opts)
)

func NewReq(body []byte, headers map[string]string) *Req {
	return &Req{
		Body:    body,
		Headers: headers,
	}
}

func WithTimeOut(t time.Duration) OptsFn {
	return func(netOpts *Opts) {
		netOpts.timeOut = t
	}
}
