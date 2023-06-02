package service

import (
	"fmt"
	"github.com/laoningmeng/go-admin/internal/service/registry"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type Components struct {
	Addr     string
	Name     string
	Port     int
	Server   *grpc.Server
	Register registry.Register
}

type CustomerComponent func(components *Components)

type Service struct {
	Components         Components
	CustomerComponents []CustomerComponent
	once               sync.Once
}

func NewService(component ...CustomerComponent) Service {
	s := grpc.NewServer()
	return Service{
		Components: Components{
			Server:   s,
			Addr:     "127.0.0.1",
			Port:     8890,
			Name:     "go-admin",
			Register: nil,
		},
		CustomerComponents: component,
	}
}

func (s *Service) Run() {
	for _, o := range s.CustomerComponents {
		o(&s.Components)
	}
	s.once.Do(func() {
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Components.Addr, s.Components.Port))
		if err != nil {
			panic(err)
		}
		err = s.Components.Server.Serve(listen)
		if err != nil {
			panic(err)
		}
	})

}
