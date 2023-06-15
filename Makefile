server:
	go run ./cmd/grpc.go
env:
	docker-compose -f ./deployment/docker-compose.yml up

init:
	go run ./cmd/init.go

cert:
	chmod +x deployment/generateCert.sh
	deployment/generateCert.sh