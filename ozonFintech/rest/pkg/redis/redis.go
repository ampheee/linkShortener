package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
)

func GetClient(ctx context.Context, c config.Config) (*redis.Client, error) {
	log := logger.GetLogger()
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisDB.Addr,  // адрес сервера Redis
		Password: c.RedisDB.Pass,  // пароль, если он установлен
		DB:       c.RedisDB.DBNum, // номер базы данных Redis
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("Redis is unable to connect.")
		return nil, err
	}
	log.Info().Msg("Redis connected successfully.")
	return rdb, nil
}
