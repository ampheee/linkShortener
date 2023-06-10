grpc_build:
	 protoc --proto_path=api/proto/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. api/proto/*proto

app-on-redis:
	docker compose up app redis --no-deps --build

app-on-postgres:
	docker compose up app postgres --no-deps --build
