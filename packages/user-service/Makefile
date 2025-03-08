test-user-service:
	gotestsum --format=short -- -failfast -coverprofile=coverage.out ./...

test-unit:
	go test -v -failfast -coverprofile=coverage.out ./...

test-integration:
	go test -v -tags=integration -failfast -coverprofile=coverage.out ./...

coverage:
	go tool cover -html="coverage.out"

dev:
	go run ./cmd/server/main.go

create-migration:
	migrate create -ext sql -dir infra/database/migrations -seq $(name)

migrate-up:
	migrate -path ./infra/database/migrations -database "postgres://postgres:postgres@localhost:5432/customerdb?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./infra/database/migrations -database "postgres://postgres:postgres@localhost:5432/customerdb?sslmode=disable" -verbose down

proto:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/*.proto

.PHONY: test-unit test-integration coverage dev create-migration migrate proto
