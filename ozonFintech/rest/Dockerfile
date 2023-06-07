FROM golang:alpine

COPY . .
WORKDIR ./app/

RUN go mod donwload

RUN go build app.go

CMD ["./app", "-migration=OnStart"]