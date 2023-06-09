FROM golang:alpine
WORKDIR ./rest
COPY . .
RUN go mod tidy
WORKDIR ./app
RUN go build app.go
EXPOSE 8080
ENTRYPOINT ["./app", "-storage=Redis"]