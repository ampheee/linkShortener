package utilities

import (
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
)

func ParseFlagsFromCLI(storageType, migration *string, c *config.Config) {
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
