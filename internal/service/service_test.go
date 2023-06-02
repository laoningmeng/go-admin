package service

import (
	"github.com/laoningmeng/go-admin/internal/service/registry"
	"testing"
)

func TestNewService(t *testing.T) {
	var consul registry.Register = &registry.Consul{
		Addr: "192.168.2.15",
		Port: 8500,
		Info: registry.ConsulServerInfo{
			Name: "go-admin",
			Port: 8890,
			Addr: "192.168.2.15",
		},
	}
	service := NewService(Addr("192.168.2.15"), Name("go-admin"), Port(8890), Register(consul))
	service.Run()
}
