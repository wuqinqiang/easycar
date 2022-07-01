package protocol

import (
	"context"

	"github.com/go-resty/resty/v2"
)

var (
	_        NetProtocol = (*HttProtocol)(nil)
	restyCli             = resty.New()
)

type (
	HttProtocol struct {
		uri string
	}
)

func NewHttpProtocol(uri string) *HttProtocol {
	return &HttProtocol{
		uri: uri,
	}
}

func (cli *HttProtocol) GetType() NetType {
	return Http
}

func (cli *HttProtocol) Request(ctx context.Context, optFns ...OptsFn) (resp *Resp, err error) {
	opts := new(Opts)
	for _, optFn := range optFns {
		optFn(opts)
	}
	resp, err = cli.req(ctx, opts.body, opts.headers)
	return
}

func (cli *HttProtocol) req(ctx context.Context, body []byte, headers map[string]string) (*Resp, error) {
	resp, err := restyCli.R().
		SetHeaders(headers).
		SetBody(body).
		Post(cli.uri)
	if err != nil {
		return nil, err
	}
	return &Resp{
		Code: int64(resp.StatusCode()),
		Body: resp.Body(),
	}, nil
}
