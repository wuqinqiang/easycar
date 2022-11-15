package grpc

import (
	"context"
	"sync"

	"github.com/wuqinqiang/easycar/logging"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/wuqinqiang/easycar/core/transport/common"
)

func NewTransporter() *Transport {
	return &Transport{m: &sync.Map{}}
}

type Transport struct {
	m *sync.Map
}

func (g *Transport) GetType() common.Net {
	return common.Grpc
}

func (g *Transport) Request(ctx context.Context, uri string, req *common.Req) (*common.Resp, error) {
	parse, err := g.getParse(uri)
	if err != nil {
		return nil, err
	}
	server, method, err := parse.Get()
	if err != nil {
		return nil, err
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()

	conn, err := g.getClient(timeoutCtx, server)
	if err != nil {
		return nil, err
	}
	md := metadata.New(req.Headers)
	err = conn.Invoke(metadata.NewOutgoingContext(ctx, md), method, req.Body, &[]byte{})
	return nil, err
}

func (g *Transport) getClient(ctx context.Context, server string) (conn *grpc.ClientConn, err error) {
	val, ok := g.m.Load(server)
	if !ok {
		conn, err = g.getConn(ctx, server)
		if err != nil {
			return
		}
		g.m.Store(server, conn)
		return
	}
	conn = val.(*grpc.ClientConn)
	return
}

func (g *Transport) Close(ctx context.Context) error {
	g.m.Range(func(key, value any) bool {
		conn := value.(*grpc.ClientConn)
		if err := conn.Close(); err != nil {
			logging.Errorf("[grpc] Transport close err:%v", err)
		}
		return true
	})
	logging.Infof("[grpc] all ClientConn closed")
	return nil
}

func (g *Transport) getConn(ctx context.Context, uri string) (*grpc.ClientConn, error) {
	opts := grpc.WithDefaultCallOptions(grpc.ForceCodec(rawCodec{}))
	return grpc.DialContext(ctx, uri, grpc.WithTransportCredentials(insecure.NewCredentials()), opts)
}

func (g *Transport) getParse(server string) (Parser, error) {
	return NewDefault(server), nil
}
