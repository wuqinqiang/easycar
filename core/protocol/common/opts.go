package common

import (
	"encoding/json"
	"time"
)

type (
	ReqOpt func(req *Req)
	Req    struct {
		Body    []byte
		Headers map[string]string
		timeOut time.Duration
	}
)

func NewReq(body []byte, headers []byte, opts ...ReqOpt) *Req {
	h := make(map[string]string)
	if len(headers) > 0 {
		_ = json.Unmarshal(headers, &h)
	}
	req := &Req{
		Body:    body,
		Headers: h,
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
