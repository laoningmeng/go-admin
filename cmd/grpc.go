package main

import (
	"github.com/laoningmeng/go-admin/internal/service"
	"github.com/laoningmeng/go-admin/internal/service/handler"
	"github.com/laoningmeng/go-admin/internal/service/middleware"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"github.com/laoningmeng/go-admin/internal/service/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var consul registry.Register = &registry.Consul{
		Addr:          "127.0.0.1",
		Port:          8500,
		Token:         "70002d33-a580-91f0-7612-9ffbc04ab589",
		TLSSkipVerify: true,
		GRPCUseTLS:    true,
		Info: registry.ConsulServerInfo{
			Name: "go-admin",
			Port: 8890,
			Addr: "192.168.81.119",
		},
	}
	cred, err := credentials.NewServerTLSFromFile("../configs/cert/server.pem", "../configs/cert/server.key")
	if err != nil {
		panic("找不到证书信息")
	}
	opt := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.Interceptor),
		grpc.Creds(cred),
	}
	s := service.NewService(grpc.NewServer(opt...))
	option := []service.CustomerComponent{
		service.Addr("0.0.0.0"),
		service.Name("go-admin"),
		service.Port(8890),
		service.Register(consul),
	}
	s.Init(option)
	sys.RegisterBaseServer(s.Components.Server, &handler.Base{})
	s.Run()
}
