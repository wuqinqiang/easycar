package common

import "time"

type (
	ReqOpt func(req *Req)
	Req    struct {
		Body    []byte
		Headers map[string]string
		timeOut time.Duration
	}
)

func NewReq(body []byte, headers map[string]string, opts ...ReqOpt) *Req {
	req := &Req{
		Body:    body,
		Headers: headers,
		timeOut: 8 * time.Second,
	}
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func WithTimeOut(t time.Duration) ReqOpt {
	return func(req *Req) {
		if t == 0 {
			return
		}
		req.timeOut = t
	}
}
