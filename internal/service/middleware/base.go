package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var whiteList = []string{
	"/grpc.health.v1.Health/Check", // 健康检测
	"/base/Login",                  // 登录路由
}

// checkWhiteList  检测白名单
func checkWhiteList(fullMethod string) bool {
	for _, ele := range whiteList {
		if fullMethod == ele {
			return true
		}
	}
	return false
}

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	beforeErr := Before(ctx, req, info)
	if beforeErr != nil {
		return nil, beforeErr
	}
	resp, err = handler(ctx, req)
	After()
	return
}

func Before(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) error {
	if ok := checkWhiteList(info.FullMethod); ok {
		return nil
	}
	m, haveMata := metadata.FromIncomingContext(ctx)
	if !haveMata {
		return errors.New("缺少参数")
	}
	if _, ok := m["token"]; !ok {
		return errors.New("缺少参数")
	}
	// 根据token 查询是否存在，这里没有查询数据，写的固定值
	if m["token"][0] != "abcd" {
		return errors.New("没有权限")
	}
	return nil
}

func After() {

}
