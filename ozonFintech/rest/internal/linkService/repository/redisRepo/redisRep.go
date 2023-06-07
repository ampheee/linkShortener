package RedisRepo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"ozonFintech/internal/linkService"
	"ozonFintech/pkg/logger"
)

type Redis struct {
	logger zerolog.Logger
	pool   *pgxpool.Pool
}

func (r Redis) SelectLink(ctx context.Context, originalLink string) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) InsertLink(ctx context.Context, abbreviatedLink string) error {
	//TODO implement me
	panic("implement me")
}

func NewRedisRep() linkService.Repo {
	return &Redis{
		logger: logger.GetLogger(),
		pool:   nil,
	}
}
