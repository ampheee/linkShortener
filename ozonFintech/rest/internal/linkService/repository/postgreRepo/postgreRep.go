package postgreRepo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"ozonFintech/internal/linkService"
	"ozonFintech/pkg/logger"
)

type Postgres struct {
	logger zerolog.Logger
	pool   *pgxpool.Pool
}

func (p Postgres) SelectLink(ctx context.Context, originalLink string) error {
	panic("implement me!")
}

func (p Postgres) InsertLink(ctx context.Context, abbreviatedLink string) error {
	panic("implement me!")
}

func NewPostgreRep(pool *pgxpool.Pool) linkService.Repo {
	return &Postgres{
		logger: logger.GetLogger(),
		pool:   pool,
	}
}
