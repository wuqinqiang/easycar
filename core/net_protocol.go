package core

import (
	"context"
	"errors"
)

type NetType string

const (
	Http NetType = "http"
	Grpc NetType = "grpc"
	// todo more
)

var NotFoundProtocol = errors.New("not found protocol")

type NetProtocol interface {
	// GetType returns the type of the net protocol
	GetType() NetType
	// todo replace item to real request param
	Request(ctx context.Context, item interface{}) error
}

var protocolMap map[NetType]NetProtocol

func init() {
	protocolMap = make(map[NetType]NetProtocol)
}

func RegisterProtocol(net NetType, netProtocol NetProtocol) {
	protocolMap[net] = netProtocol
}

func GetProtocol(net NetType) (NetProtocol, error) {
	protocol, ok := protocolMap[net]
	if !ok {
		return nil, NotFoundProtocol
	}
	return protocol, nil
}
