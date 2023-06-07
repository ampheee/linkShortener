package utilities

import (
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
)

func ParseFlags(storageType, migration *string, c *config.Config) {
	log := logger.GetLogger()
	switch *storageType {
	case "In-memory_Redis":
		c.StorageType = "In-memory_Redis"
	default:
		log.Warn().Msg("Unknown storage type. PostgreSQL set as default")
		c.StorageType = "PostgreSQL"
	}
	c.Migration = *migration
}
