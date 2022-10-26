package common

import (
	"encoding/json"
	"time"
)

var (
	defaultTimeOut = 8 * time.Second
)

type (
	Option func(req *Req)
	Req    struct {
		Body    []byte
		Headers map[string]string
		Timeout time.Duration
	}
)

func NewReq(body, headers []byte, opts ...Option) *Req {
	h := make(map[string]string)
	if len(headers) > 0 {
		_ = json.Unmarshal(headers, &h)
	}
	req := &Req{
		Body:    body,
		Headers: h,
		Timeout: defaultTimeOut,
	}
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func (r *Req) AddEasyCarHeaders(gId, branchId string) {
	r.Headers["easyCarGId"] = gId
	r.Headers["easyCarBranchId"] = branchId
	//  should to add request id?
}

func WithTimeOut(t time.Duration) Option {
	return func(req *Req) {
		if t == 0 {
			return
		}
		req.Timeout = t
	}
}

func ReplaceTimeOut(t time.Duration) {
	if t > 0 {
		defaultTimeOut = t
	}
}
