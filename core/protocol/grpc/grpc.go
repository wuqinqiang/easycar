package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/wuqinqiang/easycar/core/protocol/common"
)

type Protocol struct {
	uri string
}

func NewProtocol(uri string) *Protocol {
	return &Protocol{uri: uri}
}

func (g *Protocol) GetType() common.Net {
	return common.Grpc
}

func (g *Protocol) Request(ctx context.Context, req *common.Req) (*common.Resp, error) {
	parse, err := g.getParse(g.uri)
	if err != nil {
		return nil, err
	}
	server, method, err := parse.Get()
	if err != nil {
		return nil, err
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()
	conn, err := g.getConn(timeoutCtx, server)
	if err != nil {
		return nil, err
	}
	var (
		respM []byte
	)
	md := metadata.New(req.Headers)
	err = conn.Invoke(metadata.NewOutgoingContext(ctx, md), method, req.Body, &respM)
	return nil, err
}

func (g *Protocol) getConn(ctx context.Context, uri string) (*grpc.ClientConn, error) {
	opts := grpc.WithDefaultCallOptions(grpc.ForceCodec(rawCodec{}))
	return grpc.DialContext(ctx, uri, grpc.WithTransportCredentials(insecure.NewCredentials()), opts)
}

func (g *Protocol) getParse(server string) (Parser, error) {
	return NewDefault(server), nil
}
