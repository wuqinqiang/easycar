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

func (g *Protocol) Request(ctx context.Context, optFns ...common.OptsFn) (*common.Resp, error) {
	opts := new(common.Opts)
	for _, optFn := range optFns {
		optFn(opts)
	}
	parse, err := g.getParse(g.uri)
	if err != nil {
		return nil, err
	}
	server, method, err := parse.Get()
	if err != nil {
		return nil, err
	}
	conn, err := g.getConn(server)
	if err != nil {
		return nil, err
	}

	var (
		respM []byte
	)
	if err = conn.Invoke(ctx, method, opts.Body, &respM); err != nil {
		return nil, err
	}
	fmt.Println("数据:", string(respM))
	return nil, nil
}

func (g *Protocol) getConn(uri string) (*grpc.ClientConn, error) {
	codecOpt := grpc.ForceCodec(rawCodec{})
	opts := grpc.WithDefaultCallOptions(codecOpt)
	// todo add more options
	return grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()), opts)
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

func (cb rawCodec) Name() string { return "dtm_raw" }
