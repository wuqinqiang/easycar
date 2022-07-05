package grpc

import (
	"context"
	"testing"
)

func TestProtocol_Request(t *testing.T) {
	s := NewProtocol("127.0.0.1:8089/proto.EasyCar/Begin")
	_, err := s.Request(context.Background())
	if err != nil {
		panic(err)
	}

}
