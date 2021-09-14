run:
	@go run cmd/server/main.go
	
test_api:
	@go test ./test/...

test_usecase:
	@go test ./internal/usecase/...

test_entity:
	@go test ./internal/entity/...

test_all:
	@go test ./...
