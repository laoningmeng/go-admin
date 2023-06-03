server:
	go run ./cmd/grpc.go
consul:
	docker-compose -f ./deployment/docker-compose.yml up