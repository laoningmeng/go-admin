package handler

import (
	"context"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
)

type Base struct {
	sys.UnimplementedBaseServer
}

func (b *Base) Login(c context.Context, r *sys.LoginRequest) (*sys.LoginResponse, error) {
	// todo 1. 根据用户明查找是否存在，不存在直接返回，否则返回
	return &sys.LoginResponse{Token: "my_token"}, nil
}

func (b *Base) Info(c context.Context, r *sys.InfoRequest) (*sys.InfoResponse, error) {

	return nil, nil
}
