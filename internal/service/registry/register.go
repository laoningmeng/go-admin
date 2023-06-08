package registry

import (
	"fmt"
	"github.com/google/uuid"
	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"strconv"
	"sync"
)

type Register interface {
	Init()
	HealthCheck(server *grpc.Server)
}

type ConsulServerInfo struct {
	Name string
	Port int
	Addr string
}

type Consul struct {
	Addr          string
	Port          int
	Info          ConsulServerInfo
	Token         string
	TLSSkipVerify bool
	GRPCUseTLS    bool
	once          sync.Once
}

func (c *Consul) HealthCheck(s *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
}

func (c *Consul) Init() {
	cfg := capi.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.Addr, c.Port)
	cfg.Token = c.Token
	client, err := capi.NewClient(cfg)
	if err != nil {
		log.Fatal("registry Error:", err)
	}
	err = client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      fmt.Sprint(uuid.New()),
		Name:    c.Info.Name,
		Port:    c.Info.Port,
		Address: c.Info.Addr,
		Check: &capi.AgentServiceCheck{
			Name:                           c.Info.Name,
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%s:%s", c.Info.Addr, strconv.Itoa(c.Info.Port)),
			DeregisterCriticalServiceAfter: "5s",
			//如果开启TLS,使用以下配置跳过TLS验证
			TLSSkipVerify: c.TLSSkipVerify,
			GRPCUseTLS:    c.GRPCUseTLS,
		},
	})
}
