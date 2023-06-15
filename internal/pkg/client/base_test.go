package client

import (
	"context"
	"fmt"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"testing"
)

func TestLogin(t *testing.T) {
	client := NewBaseClient()
	login, err := client.Login(context.Background(), &sys.LoginRequest{
		Username: "admin",
		Password: "123456",
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(login)
}
