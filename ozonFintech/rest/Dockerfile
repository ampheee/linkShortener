FROM golang:alpine

WORKDIR ./rest
COPY . .
RUN go mod download
WORKDIR ./app
RUN go build app.go

EXPOSE 8080

CMD ["./app", "-migration=OnStart"]