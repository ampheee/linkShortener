FROM golang:alpine

RUN go mod download

COPY . .
WORKDIR ./app/

RUN go build app.go

CMD ["./app", "-migration=OnStart"]