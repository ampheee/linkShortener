package config

import (
	"github.com/spf13/viper"
	"ozonFintech/pkg/logger"
)

type Config struct {
	PostgreSQLDB struct {
		User     string
		Pass     string
		Host     string
		Port     string
		Dbname   string
		SSLMode  string
		MaxConns string
	}
	RedisDB struct {
		Addr string
	}
	StorageType string
	Migration   string
}

func LoadConfig() *viper.Viper {
	log := logger.GetLogger()
	v := viper.New()
	v.AddConfigPath("../config/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load config.")
	}
	log.Info().Msg("Config loaded successfully.")
	return v
}

func ParseConfig(v *viper.Viper) Config {
	log := logger.GetLogger()
	var conf Config
	err := v.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse config.")
	}
	log.Info().Msg("Config parsed successfully.")
	return conf
}
