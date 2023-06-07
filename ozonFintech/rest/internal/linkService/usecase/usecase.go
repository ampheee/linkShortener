package usecase

import (
	"context"
	"github.com/rs/zerolog"
	"ozonFintech/config"
	"ozonFintech/internal/linkService"
	"ozonFintech/internal/linkService/repository/postgreRepo"
	RedisRepo "ozonFintech/internal/linkService/repository/redisRepo"
	"ozonFintech/internal/utilities"
	"ozonFintech/pkg/logger"
	"ozonFintech/pkg/postgresql"
)

type LinkUseCase struct {
	Logger zerolog.Logger
	Repo   linkService.Repo
}

func (l LinkUseCase) GetOriginalByAbbreviated(ctx context.Context, link string) (string, error) {
	originalLink, err := l.Repo.SelectLink(ctx, link)
	if err != nil {
		l.Logger.Warn().Err(err).Msg("unable to get link by abbreviated.")
		return "", err
	}
	if originalLink == "" {
		l.Logger.Warn().Msg("no originalLink by abbreviated.")
	}
	return originalLink, nil
}

func (l LinkUseCase) SaveOriginalLink(ctx context.Context, link string) (string, error) {
	abbreviatedLink, err := utilities.Encrypt(link)
	if err != nil {
		return "", err
	}
	if l.Repo.InsertLink(ctx, abbreviatedLink, link) != nil {
		return "", err
	}
	return abbreviatedLink, nil
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
