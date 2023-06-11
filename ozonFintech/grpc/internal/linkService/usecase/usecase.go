package usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"grpcService/config"
	"grpcService/grpc_domain"
	"grpcService/internal/linkService"
	"grpcService/internal/linkService/repository/postgreRepo"
	RedisRepo "grpcService/internal/linkService/repository/redisRepo"
	"grpcService/pkg/logger"
	"grpcService/pkg/postgresql"
	"grpcService/pkg/redisDB"
	"grpcService/pkg/utilities"
)

type LinkUseCase struct {
	grpc_domain.UnimplementedLinkServiceServer
	Logger zerolog.Logger
	Repo   linkService.Repo
}

func (l LinkUseCase) Get(ctx context.Context, r *grpc_domain.GetLinkRequest) (*grpc_domain.GetLinkResponse, error) {
	l.Logger.Info().Msg(r.String())
	originalLink, err := l.Repo.SelectLink(ctx, r.AbbreviatedLink)
	if err != nil {
		l.Logger.Warn().Err(err).Msg("unable to get link by abbreviated.")
		return &grpc_domain.GetLinkResponse{OrigLink: ""}, err
	}
	return &grpc_domain.GetLinkResponse{OrigLink: originalLink}, nil
}

func (l LinkUseCase) Create(ctx context.Context, r *grpc_domain.CreateLinkRequest) (*grpc_domain.CreateLinkResponse, error) {
	if r.OrigLink == "" {
		return &grpc_domain.CreateLinkResponse{}, errors.New("EmptyLink")
	}
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(r.OrigLink))
	l.Logger.Info().Msg(abbreviatedLink)
	if err := l.Repo.InsertLink(ctx, abbreviatedLink, r.OrigLink); err != nil {
		return &grpc_domain.CreateLinkResponse{}, err
	}
	l.Logger.Info().Msg("link saved and abbreviated returned.")
	return &grpc_domain.CreateLinkResponse{AbbreviatedLink: "rus.tam/" + abbreviatedLink}, nil
}

func NewLinkService(ctx context.Context, c config.Config) (grpc_domain.LinkServiceServer, error) {
	logg := logger.GetLogger()
	var repo linkService.Repo
	switch c.StorageType {
	case "Redis":
		client, err := redisDB.GetClient(ctx, c)
		if err != nil {
			logg.Warn().Err(err).Msg("unable to get redis client while call newLinkService.")
			return nil, err
		}
		repo = RedisRepo.NewRedisRep(client)
	case "PostgreSQL":
		pool, err := postgresql.GetPool(ctx, c)
		if err != nil {
			logg.Warn().Err(err).Msg("unable to get postgresql pool while call newLinkService.")
			return nil, err
		}
		postgresql.InitMigration(c)
		repo = postgreRepo.NewPostgreRep(pool)
	}
	return &LinkUseCase{
		Logger: logg,
		Repo:   repo,
	}, nil
}
