package grpc

import (
	"context"
	"fmt"

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

func (g *Protocol) GetType() common.NetType {
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
	conn, err := g.getConn(ctx, server)
	if err != nil {
		return nil, err
	}
	var (
		respM []byte
	)
	err = conn.Invoke(ctx, method, req.Body, &respM)
	return nil, err
}

func (g *Protocol) getConn(ctx context.Context, uri string) (*grpc.ClientConn, error) {
	codecOpt := grpc.ForceCodec(rawCodec{})
	opts := grpc.WithDefaultCallOptions(codecOpt)
	return grpc.DialContext(ctx, uri, grpc.WithTransportCredentials(insecure.NewCredentials()), opts)
}

func (g *Protocol) getParse(server string) (Parser, error) {
	return NewDefault(server), nil
}

type rawCodec struct{}

func (cb rawCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (cb rawCodec) Unmarshal(data []byte, v interface{}) error {
	ba, ok := v.(*[]byte)
	if !ok {
		return fmt.Errorf("please pass in *[]byte")
	}
	*ba = append(*ba, data...)

	return nil
}

func (cb rawCodec) Name() string { return "easycar" }
