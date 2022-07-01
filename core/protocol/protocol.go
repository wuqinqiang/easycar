package protocol

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

var NotFoundProtocol = errors.New("not found protocol")

type NetProtocol interface {
	// GetType returns the type of the net protocol
	GetType() NetType
	Request(ctx context.Context, optFns ...OptsFn) (*Resp, error)
}

// todo change uri ?
func GetProtocol(net NetType, uri string) (NetProtocol, error) {
	switch net {
	case Http:
		return NewHttpProtocol(uri), nil
	case Grpc:
		// todo
	}
	return nil, NotFoundProtocol
}
