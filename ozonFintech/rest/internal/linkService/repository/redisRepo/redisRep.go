package RedisRepo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"ozonFintech/internal/linkService"
	"ozonFintech/pkg/logger"
)

type Redis struct {
	logger zerolog.Logger
	Client *redis.Client
}

func (r Redis) SelectLink(ctx context.Context, abbreviatedLink string) (string, error) {
	val, err := r.Client.Get(ctx, abbreviatedLink).Result()
	if err != nil {
		r.logger.Warn().Err(err).Msg("something went wrong while select link from redis.")
		return "", err
	}
	if val == "" {
		r.logger.Warn().Msg("no such abbreviated link in redis storage.")
		return "", nil
	}
	return val, nil
}

func (r Redis) InsertLink(ctx context.Context, abbreviatedLink, originalLink string) error {
	err := r.Client.Set(ctx, abbreviatedLink, originalLink, 0).Err()
	if err != nil {
		r.logger.Warn().Msg("something went wrong while insert link into redis.")
		return err
	}
	r.logger.Info().Msg("link successfully inserted into redis db.")
	return nil
}

func NewRedisRep(client *redis.Client) linkService.Repo {
	return &Redis{
		logger: logger.GetLogger(),
		Client: client,
	}
}
