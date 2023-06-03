package service

import (
	"fmt"
	"github.com/laoningmeng/go-admin/internal/service/logger"
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
	Logger   logger.Logger
}

type CustomerComponent func(components *Components)

type Service struct {
	Components         Components
	CustomerComponents []CustomerComponent
	once               sync.Once
}

func NewService(s *grpc.Server) Service {
	zap := logger.NewLogger()
	return Service{
		Components: Components{
			Server:   s,
			Addr:     "127.0.0.1",
			Port:     8890,
			Name:     "go-admin",
			Register: nil,
			Logger:   zap,
		},
	}
}
func (s *Service) Init(opt []CustomerComponent) {
	s.Components.Logger.Info("start run")
	for _, o := range opt {
		o(&s.Components)
	}
}

func (s *Service) Run() {
	s.once.Do(func() {
		s.Components.Logger.Infof("启动信息：Addr: %s Port:%d", s.Components.Addr, s.Components.Port)
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Components.Addr, s.Components.Port))
		if err != nil {
			panic(err)
		}
		err = s.Components.Server.Serve(listen)
		if err != nil {
			//panic(err)
			s.Components.Logger.Errorf("启动出现错误:%v", err)
		}
	})

}
