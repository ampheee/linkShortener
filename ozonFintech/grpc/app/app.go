package main

import (
	"grpcService/config"
	"grpcService/internal/client"
	"grpcService/pkg/utilities"
)

func main() {
	Config := config.ParseConfigFromEnv()
	utilities.ParseFlagsFromCLI(&Config)
	//Config := config.ParseConfigFromYaml(config.LoadConfigFromYaml())
	Client, err := client.NewClient(Config)
	if err != nil {
		return
	}
	Client.Run()

}
