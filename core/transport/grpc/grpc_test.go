package grpc

import (
	"context"
	"sync"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/wuqinqiang/easycar/core/transport/common"

	proto2 "github.com/wuqinqiang/easycar/proto"
)

func TestProtocol_Request(t *testing.T) {
	s := &Transport{m: &sync.Map{}}
	a := common.NewReq(nil, nil)
	_, err := s.Request(context.Background(), "127.0.0.1:8089/proto.EasyCar/Begin", a)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProtocol_Request2(t *testing.T) {
	s := &Transport{m: &sync.Map{}}
	req := proto2.StartReq{GId: "11"}
	reqByte, err := proto.Marshal(&req)
	if err != nil {
		t.Fatal(err)
	}
	a := common.NewReq(reqByte, nil)
	_, err = s.Request(context.Background(), "127.0.0.1:8089/proto.EasyCar/Phase1", a)
	if err != nil {
		t.Fatal(err)
	}

}
