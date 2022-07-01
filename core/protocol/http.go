package protocol

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/wuqinqiang/easycar/core"
)

var (
	_        core.NetProtocol = (*client)(nil)
	restyCli                  = resty.New()
)

type (
	Opt struct {
		body []byte
		head map[string]string
	}

	OptFn func(opt *Opt)

	client struct {
		*Opt
		uri string
	}
)

func (cli *client) GetType() core.NetType {
	return core.Http
}

func (cli *client) Request(ctx context.Context, item interface{}) error {
	//TODO implement me
	panic("implement me")
}

func WithHead(header map[string]string) OptFn {
	return func(opt *Opt) {
		opt.head = header
	}
}

func WithBody(body []byte) OptFn {
	return func(opt *Opt) {
		opt.body = body
	}
}

func NewClient(uri string, opts ...OptFn) *client {
	opt := new(Opt)
	for _, optFn := range opts {
		optFn(opt)
	}
	// todo check uri
	return &client{
		Opt: opt,
		uri: uri,
	}
}

func (cli *client) Req() (Resp, error) {
	resp, err := restyCli.R().
		SetHeaders(cli.head).
		SetBody(cli.body).
		Post(cli.uri)
	if err != nil {
		return Resp{}, err
	}
	return Resp{
		Code: int64(resp.StatusCode()),
		Body: resp.Body(),
	}, nil
}
