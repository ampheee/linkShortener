FROM golang:alpine
WORKDIR /grpc
COPY . .
RUN go mod tidy
EXPOSE 8080
WORKDIR ./app
RUN go build app.go

ENTRYPOINT ["./app", "-storage=PostgreSQL", "-migration=OnStart"]
