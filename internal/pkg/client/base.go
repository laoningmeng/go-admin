package client

import (
	"context"
	"fmt"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"os"
)

func LoginInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{
		"token": "abcd",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func NewBaseClient() sys.BaseClient {
	path, _ := os.Getwd()
	fmt.Println(path)
	c, err := credentials.NewClientTLSFromFile("../../../configs/cert/server.crt", "www.lixueduan.com")
	if err != nil {
		panic(err)
	}
	option := []grpc.DialOption{
		grpc.WithTransportCredentials(c),
		grpc.WithUnaryInterceptor(LoginInterceptor),
	}
	dial, err := grpc.Dial("192.168.2.4:8890", option...)
	if err != nil {
		panic(err)
	}
	client := sys.NewBaseClient(dial)
	return client
}
