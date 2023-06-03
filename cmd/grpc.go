package main

import (
	"github.com/laoningmeng/go-admin/internal/service"
	"github.com/laoningmeng/go-admin/internal/service/handler"
	sys "github.com/laoningmeng/go-admin/internal/service/proto"
	"github.com/laoningmeng/go-admin/internal/service/registry"
	"google.golang.org/grpc"
)

func main() {
	var consul registry.Register = &registry.Consul{
		Addr: "192.168.2.15",
		Port: 8500,
		Info: registry.ConsulServerInfo{
			Name: "go-admin",
			Port: 8890,
			Addr: "192.168.2.15",
		},
	}
	opt := []grpc.ServerOption{
		grpc.UnaryInterceptor(handler.BeforeLogin),
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
