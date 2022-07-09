package grpc

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/wuqinqiang/easycar/core/protocol/common"

	proto2 "github.com/wuqinqiang/easycar/proto"
)

func TestProtocol_Request(t *testing.T) {
	s := NewProtocol("127.0.0.1:8089/proto.EasyCar/Begin")
	_, err := s.Request(context.Background())
	if err != nil {
		panic(err)
	}

}

func TestProtocol_Request2(t *testing.T) {
	s := NewProtocol("127.0.0.1:8089/proto.EasyCar/Commit")

	req := proto2.CommitReq{GId: "11"}
	reqByte, err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	_, err = s.Request(context.Background(), common.WithBody(reqByte))
	if err != nil {
		panic(err)
	}

}
