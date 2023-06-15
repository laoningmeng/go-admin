package handler

import (
	"context"
	"errors"
	"github.com/laoningmeng/go-admin/internal/pkg/encrypt"
	"github.com/laoningmeng/go-admin/internal/service/global"
	"github.com/laoningmeng/go-admin/internal/service/model"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"strconv"
)

type Base struct {
	sys.UnimplementedBaseServer
}

func (b *Base) Login(c context.Context, r *sys.LoginRequest) (*sys.LoginResponse, error) {
	// todo 1. 根据用户明查找是否存在，不存在直接返回，否则返回
	user := model.User{
		Name:     r.Username,
		Password: encrypt.Md5(r.Password),
	}
	u, err := user.FindOne()
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 进行密码加密处理
	token, err := encrypt.GenerateToken(map[string]string{"username": u.Name, "id": strconv.FormatUint(u.ID, 10)})
	if err != nil {
		global.Logger.Fatalf("Login：生成token失败 失败信息：%v", err)
	}
	return &sys.LoginResponse{Token: token}, err

}

func (b *Base) Info(c context.Context, r *sys.InfoRequest) (*sys.InfoResponse, error) {

	return nil, nil
}
