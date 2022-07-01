package protocol

import "context"

var (
	_ NetProtocol = (*GrpcProtocol)(nil)
)

type GrpcProtocol struct {
}

func (g GrpcProtocol) GetType() NetType {
	//TODO implement me
	panic("implement me")
}

func (g GrpcProtocol) Request(ctx context.Context, optFns ...OptsFn) (*Resp, error) {
	//TODO implement me
	panic("implement me")
}
