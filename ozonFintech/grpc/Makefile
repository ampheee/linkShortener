.PHONY: app-on-postgres app-on-redis grpc_build

grpc_build:
	 protoc --proto_path=api/proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/*proto

app-on-redis:
	docker compose stop
	docker compose up app_redis redis --build

app-on-postgres:
	rm -rf ./data/postgres/*
	docker compose stop
	docker compose up app_postgres postgres --build

test:
	go test -v ./internal/linkService/mockRepo/*.go
	go test -v ./pkg/utilities/Encode_test.go ./pkg/utilities/Encode.go