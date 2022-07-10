package grpc

import (
	"context"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/wuqinqiang/easycar/core/protocol/common"

	proto2 "github.com/wuqinqiang/easycar/proto"
)

func TestProtocol_Request(t *testing.T) {
	s := NewProtocol("127.0.0.1:8089/proto.EasyCar/Begin")
	a := common.NewReq(nil, nil)
	_, err := s.Request(context.Background(), a)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProtocol_Request2(t *testing.T) {
	s := NewProtocol("127.0.0.1:8089/proto.EasyCar/Commit")

	req := proto2.CommitReq{GId: "11"}
	reqByte, err := proto.Marshal(&req)
	if err != nil {
		t.Fatal(err)
	}
	a := common.NewReq(reqByte, nil)
	_, err = s.Request(context.Background(), a)
	if err != nil {
		t.Fatal(err)
	}

}
