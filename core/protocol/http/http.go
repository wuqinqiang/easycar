package http

import (
	"context"

	"github.com/wuqinqiang/easycar/core/protocol/common"

	"github.com/go-resty/resty/v2"
)

var (
	restyCli = resty.New()
)

type (
	Transport struct {
		uri string
	}
)

func NewHttpTransport(uri string) *Transport {
	return &Transport{
		uri: uri,
	}
}

func (cli *Transport) GetType() common.NetType {
	return common.Http
}

func (cli *Transport) Request(ctx context.Context, optFns ...common.OptsFn) (resp *common.Resp, err error) {
	opts := new(common.Opts)
	for _, optFn := range optFns {
		optFn(opts)
	}
	resp, err = cli.req(ctx, opts.Body, opts.Headers)
	return
}

func (cli *Transport) req(ctx context.Context, body []byte, headers map[string]string) (*common.Resp, error) {
	resp, err := restyCli.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(body).
		Post(cli.uri)
	if err != nil {
		return nil, err
	}
	return &common.Resp{
		Code: int64(resp.StatusCode()),
		Body: resp.Body(),
	}, nil
}
