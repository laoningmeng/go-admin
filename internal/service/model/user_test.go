package model

import (
	"fmt"
	"github.com/laoningmeng/go-admin/internal/service"
	"testing"
)

func TestFindOne(t *testing.T) {
	service.LoadConfFromNacos()
	service.MysqlInit()

	u := User{Name: "demo"}

	fmt.Println(u)
}
