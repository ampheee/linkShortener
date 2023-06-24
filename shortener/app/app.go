package main

import (
	"grpcService/config"
	"grpcService/internal/server"
	"grpcService/pkg/utilities"
)

func main() {
	Config := config.ParseConfigFromEnv()
	//Config := config.ParseConfigFromYaml(config.LoadConfigFromYaml())
	utilities.ParseFlagsFromCLI(&Config)
	Server, err := server.NewServer(Config)
	if err != nil {
		return
	}
	Server.Run()

}
