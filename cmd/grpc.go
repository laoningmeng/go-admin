package main

import (
	"github.com/laoningmeng/go-admin/internal/service"
	"github.com/laoningmeng/go-admin/internal/service/global"
	"github.com/laoningmeng/go-admin/internal/service/handler"
	"github.com/laoningmeng/go-admin/internal/service/middleware"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	consul := service.NewConsul()
	service.LoadConfFromNacos()
	service.MysqlInit()
	cred, err := credentials.NewServerTLSFromFile(global.ServerConf.CertFilePath, global.ServerConf.CertKeyPath)
	if err != nil {
		panic("找不到证书信息")
	}
	opt := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.Interceptor),
		grpc.Creds(cred),
	}
	s := service.NewService(grpc.NewServer(opt...))
	option := []service.CustomerComponent{
		service.Addr(global.ServerConf.Addr),
		service.Name(global.ServerConf.Name),
		service.Port(global.ServerConf.Port),
		service.Register(consul),
	}
	s.Init(option)
	sys.RegisterBaseServer(s.Components.Server, &handler.Base{})
	s.Run()
}
