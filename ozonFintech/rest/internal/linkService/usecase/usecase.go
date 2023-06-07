package usecase

import (
	"context"
	"github.com/rs/zerolog"
	"ozonFintech/config"
	"ozonFintech/internal/linkService"
	"ozonFintech/internal/linkService/repository/postgreRepo"
	RedisRepo "ozonFintech/internal/linkService/repository/redisRepo"
	"ozonFintech/pkg/logger"
	"ozonFintech/pkg/postgresql"
)

type LinkUseCase struct {
	Logger zerolog.Logger
	Repo   linkService.Repo
}

func (l LinkUseCase) GetAbbreviatedLink(ctx context.Context, link string) (linkService.LinkDTO, error) {
	panic("Implement me!")
}

func (l LinkUseCase) SaveOriginalLink(ctx context.Context, link string) (linkService.LinkDTO, error) {
	panic("Implement me!")
}

func NewLinkService(ctx context.Context, c config.Config) linkService.Link {
	var repo linkService.Repo
	switch c.StorageType {
	case "In-memory_Redis":
		repo = RedisRepo.NewRedisRep()
	case "PostgreSQL":
		pool := postgresql.GetPool(ctx, c)
		postgresql.InitMigration(c)
		repo = postgreRepo.NewPostgreRep(pool)
	}
	return &LinkUseCase{
		Logger: logger.GetLogger(),
		Repo:   repo,
	}
}
