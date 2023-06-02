package main

import (
	"github.com/laoningmeng/go-admin/internal/service"
	"github.com/laoningmeng/go-admin/internal/service/registry"
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
	service := service.NewService(
		service.Addr("192.168.2.15"),
		service.Name("go-admin"),
		service.Port(8890),
		service.Register(consul),
	)
	service.Run()
}
