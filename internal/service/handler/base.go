package handler

import (
	"context"
	"errors"
	"fmt"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Base struct {
	sys.UnimplementedBaseServer
}

func CheckIsGrpcHealthCheck(fullMethod string) bool {
	if fullMethod == "/grpc.health.v1.Health/Check" {
		return true
	}
	return false
}

func BeforeLogin(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// todo 检测token ,并查询是否具有权限
	if ok := CheckIsGrpcHealthCheck(info.FullMethod); ok {
		// 健康检测
		return handler(ctx, req)
	}

	m, haveMata := metadata.FromIncomingContext(ctx)
	if !haveMata {
		return nil, errors.New("缺少参数")
	}
	if _, ok := m["token"]; !ok {
		return nil, errors.New("缺少参数")
	}
	// 根据token 查询是否存在，这里没有查询数据，写的固定值
	if m["token"][0] != "abcd" {
		return nil, errors.New("没有权限")
	}
	fmt.Println("before login")
	return handler(ctx, req)
}

func (b *Base) Login(c context.Context, r *sys.LoginRequest) (*sys.LoginResponse, error) {
	// todo 1. 根据用户明查找是否存在，不存在直接返回，否则返回
	return &sys.LoginResponse{Token: "my_token"}, nil
}

func (b *Base) Info(c context.Context, r *sys.InfoRequest) (*sys.InfoResponse, error) {

	return nil, nil
}
