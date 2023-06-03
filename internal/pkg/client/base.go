package client

import (
	"context"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	option := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(LoginInterceptor),
	}
	dial, err := grpc.Dial("127.0.0.1:8890", option...)
	if err != nil {
		panic(err)
	}
	client := sys.NewBaseClient(dial)
	return client
}
