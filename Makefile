run:
	@go run cmd/server/main.go
	
build:
	@go build -o main cmd/server/main.go

test_api:
	@go test ./test/...

test_usecase:
	@go test ./internal/usecase/...

test_entity:
	@go test ./internal/entity/...

test_all:
	@go test ./...

build_docker:
	@docker build -t daopmdean/todoapp -f build/docker/Dockerfile .

build_docker_test:
	@docker build -t todoapp_test -f build/docker/Dockerfile.test .

run_docker:
	@docker run -it -p 5000:5000 todoapp

.PHONY: run build test_api test_usecase test_entity test_all build_docker build_docker_test run_docker
