package transport

import "context"

var (
	_ NetTransport = (*GrpcTransport)(nil)
)

type GrpcTransport struct {
}

func (g GrpcTransport) GetType() NetType {
	//TODO implement me
	panic("implement me")
}

func (g GrpcTransport) Request(ctx context.Context, optFns ...OptsFn) (*Resp, error) {
	//TODO implement me
	panic("implement me")
}
