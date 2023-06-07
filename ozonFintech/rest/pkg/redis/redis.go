package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
)

func GetConn(ctx context.Context, c config.Config) *redis.Client {
	log := logger.GetLogger()
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisDB.Addr, // адрес сервера Redis
		Password: "",             // пароль, если он установлен
		DB:       0,              // номер базы данных Redis
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("Redis is unable to connect.")
	}
	log.Info().Msg("Redis connected successfully.")
	return rdb
}
