package transport

import (
	"context"

	"github.com/go-resty/resty/v2"
)

var (
	_        NetTransport = (*HttpTransport)(nil)
	restyCli              = resty.New()
)

type (
	HttpTransport struct {
		uri string
	}
)

func NewHttpTransport(uri string) *HttpTransport {
	return &HttpTransport{
		uri: uri,
	}
}

func (cli *HttpTransport) GetType() NetType {
	return Http
}

func (cli *HttpTransport) Request(ctx context.Context, optFns ...OptsFn) (resp *Resp, err error) {
	opts := new(Opts)
	for _, optFn := range optFns {
		optFn(opts)
	}
	resp, err = cli.req(ctx, opts.body, opts.headers)
	return
}

func (cli *HttpTransport) req(ctx context.Context, body []byte, headers map[string]string) (*Resp, error) {
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
