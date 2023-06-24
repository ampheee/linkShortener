package utilities

import (
	"flag"
	"grpcService/config"
	"grpcService/pkg/logger"
)

var (
	storageType = flag.String("storage", "PostgreSQL", "Enter an storage type: Redis or "+
		"PostgreSQL.")
	migration = flag.String("migration", "", "Enter an migration step to do. Up/Down/OnStart(default)")
)

func ParseFlagsFromCLI(c *config.Config) {
	flag.Parse()
	log := logger.GetLogger()
	switch *storageType {
	case "Redis":
		log.Info().Msg("storage type is redis.")
		c.StorageType = "Redis"
	case "PostgreSQL":
		log.Info().Msg("storage type set as PostgreSQL (default).")
		c.StorageType = "PostgreSQL"
	default:
		log.Warn().Msg("Unknown storage type. PostgreSQL set as default")
		c.StorageType = "PostgreSQL"
	}
	if c.StorageType == "Redis" {
		c.Migration = ""
	} else {
		c.Migration = *migration
	}
}
