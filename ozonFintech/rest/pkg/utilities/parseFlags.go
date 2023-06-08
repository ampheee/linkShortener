package utilities

import (
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
)

func ParseFlags(storageType, migration *string, c *config.Config) {
	log := logger.GetLogger()
	switch *storageType {
	case "Redis":
		c.StorageType = "Redis"
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
