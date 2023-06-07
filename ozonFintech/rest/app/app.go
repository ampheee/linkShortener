package main

import (
	"context"
	"flag"
	"ozonFintech/config"
	"ozonFintech/internal/client"
	"ozonFintech/internal/utilities"
	"ozonFintech/pkg/logger"
)

var storageType = flag.String("storage", "PostgreSQL", "Enter an storage type: In-memory_Redis or "+
	"PostgreSQL.")

var migration = flag.String("migration", "", "Enter an migration step to do. Up/Down/OnStart(default)")

func main() {
	flag.Parse()
	log := logger.GetLogger()
	ctx := context.Background()
	c := config.ParseConfig(config.LoadConfig())
	utilities.ParseFlags(storageType, migration, &c)
	Client := client.NewClient(ctx, c)
	log.Fatal().Err(Client.App.Listen(":8080")).Msg("Unable")
}
