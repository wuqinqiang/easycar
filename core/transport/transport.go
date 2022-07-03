package transport

import (
	"context"
	"errors"
)

type (
	NetType string
)

const (
	Http NetType = "http"
	Grpc NetType = "grpc"
)

var NotFoundTransport = errors.New("not found transport")

type NetTransport interface {
	// GetType returns the type of the net transport
	GetType() NetType
	Request(ctx context.Context, optFns ...OptsFn) (*Resp, error)
}

// todo change uri ?
func GetProtocol(net NetType, uri string) (NetTransport, error) {
	switch net {
	case Http:
		return NewHttpTransport(uri), nil
	case Grpc:
		// todo
	}
	return nil, NotFoundTransport
}
