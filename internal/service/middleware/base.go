package middleware

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var whiteList = []string{
	"/grpc.health.v1.Health/Check",
}

// checkWhiteList  检测白名单
func checkWhiteList(fullMethod string) bool {

	return false
}

func BeforeRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// todo 检测token ,并查询是否具有权限
	if ok := checkWhiteList(info.FullMethod); ok {
		//白名单跳过处理
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
