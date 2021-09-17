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

build_docker_image:
	@docker build -t todoapp -f build/docker/Dockerfile .

run_docker:
	@docker run -it -p 5000:5000 todoapp