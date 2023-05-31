package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var opt []grpc.ServerOption
	// load config

	grpcServer := grpc.NewServer(opt...)

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	_ = grpcServer.Serve(listener)
}
