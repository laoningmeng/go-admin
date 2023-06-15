package registry

import (
	"google.golang.org/grpc"
)

type Register interface {
	Init()
	HealthCheck(server *grpc.Server)
}
