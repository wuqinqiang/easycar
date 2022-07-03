package main

import (
	"net"

	grpc2 "github.com/wuqinqiang/easycar/services/grpc"

	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterEasyCarServer(server, &grpc2.EasyCarSrv{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
