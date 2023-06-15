package main

import (
	"github.com/laoningmeng/go-admin/internal/service"
	"github.com/laoningmeng/go-admin/internal/service/global"
	"github.com/laoningmeng/go-admin/internal/service/model"
)

func main() {
	// 数据表创建
	service.LoadConfFromNacos()
	service.MysqlInit()
	err := global.DB.Migrator().AutoMigrate(&model.User{}, &model.Role{}, &model.Rule{})
	if err != nil {
		panic(err)
	}
}
