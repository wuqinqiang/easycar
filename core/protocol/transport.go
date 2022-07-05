package protocol

import (
	"context"
	"errors"

	"github.com/wuqinqiang/easycar/core/protocol/grpc"

	"github.com/wuqinqiang/easycar/core/protocol/common"

	"github.com/wuqinqiang/easycar/core/protocol/http"
)

var NotFoundTransport = errors.New("not found protocol")

type NetTransport interface {
	// GetType returns the type of the net protocol
	GetType() common.NetType
	Request(ctx context.Context, optFns ...common.OptsFn) (*common.Resp, error)
}

func GetTransport(net common.NetType, service string) (NetTransport, error) {
	switch net {
	case common.Http:
		return http.NewHttpTransport(service), nil
	case common.Grpc:
		return grpc.NewProtocol(service), nil
	}
	return nil, NotFoundTransport
}
