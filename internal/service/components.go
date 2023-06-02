package service

import (
	"github.com/laoningmeng/go-admin/internal/service/logger"
	"github.com/laoningmeng/go-admin/internal/service/registry"
	"google.golang.org/grpc"
)

func Server(s *grpc.Server) CustomerComponent {
	return func(components *Components) {
		components.Server = s
	}
}

func Addr(v string) CustomerComponent {
	return func(components *Components) {
		components.Addr = v
	}
}
func Name(v string) CustomerComponent {
	return func(components *Components) {
		components.Name = v
	}
}

func Port(p int) CustomerComponent {
	return func(components *Components) {
		components.Port = p
	}
}

func Register(r registry.Register) CustomerComponent {
	return func(components *Components) {
		r.Init()
		r.HealthCheck(components.Server)
		components.Register = r
	}
}
func Logger(l logger.Logger) CustomerComponent {
	return func(components *Components) {

	}
}
