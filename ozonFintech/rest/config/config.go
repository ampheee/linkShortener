package config

import (
	"github.com/spf13/viper"
	"os"
	"ozonFintech/pkg/logger"
)

type Config struct {
	PostgreSQLDB struct {
		User    string
		Pass    string
		Host    string
		Port    string
		Dbname  string
		SSLMode string
	}
	RedisDB struct {
		Addr  string
		Pass  string
		DBNum string
	}
	StorageType string
	Migration   string
}

func LoadConfigFromYaml() *viper.Viper {
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

func ParseConfigFromYaml(v *viper.Viper) Config {
	log := logger.GetLogger()
	var conf Config
	err := v.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse config.")
	}
	log.Info().Msg("Config parsed successfully.")
	return conf
}

func ParseConfigFromEnv() Config {
	log := logger.GetLogger()
	c := Config{
		PostgreSQLDB: struct {
			User    string
			Pass    string
			Host    string
			Port    string
			Dbname  string
			SSLMode string
		}{
			User:    os.Getenv("POSTGRES_USER"),
			Pass:    os.Getenv("POSTGRES_PASS"),
			Host:    os.Getenv("POSTGRES_HOST"),
			Port:    os.Getenv("POSTGRES_PORT"),
			Dbname:  os.Getenv("POSTGRES_DB_NAME"),
			SSLMode: os.Getenv("POSTGRES_SSL_MODE"),
		}, RedisDB: struct {
			Addr  string
			Pass  string
			DBNum string
		}{
			Addr:  os.Getenv("REDIS_ADDR"),
			Pass:  os.Getenv("REDIS_PASS"),
			DBNum: os.Getenv("REDIS_DB_NUM"),
		},
	}
	log.Info().Msg("Config parsed successfully.")
	return c
}
