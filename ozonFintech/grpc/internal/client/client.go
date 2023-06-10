package client

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcService/config"
	"grpcService/grpc_domain"
	"grpcService/internal/linkService/usecase"
	"grpcService/pkg/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppContext struct {
	Server *grpc.Server
	Logger zerolog.Logger
	Config config.Config
}

func NewClient(c config.Config) (*AppContext, error) {
	log := logger.GetLogger()
	lService, err := usecase.NewLinkService(context.Background(), c)
	if err != nil {
		log.Warn().Err(err).Msg("Unable to get linkService")
		return nil, err
	}
	s := grpc.NewServer()
	grpc_domain.RegisterLinkServiceServer(s, lService)
	aCtx := &AppContext{
		Server: s,
		Logger: log,
		Config: c,
	}
	log.Info().Msg("new client created successfully.")
	return aCtx, nil
}

func (aCtx *AppContext) Run() {

	Listener, err := net.Listen("tcp", aCtx.Config.GRPCPort)
	if err != nil {

	}
	go func() {
		if err = aCtx.Server.Serve(Listener); err != nil {
			aCtx.Logger.Fatal().Err(err).Msg("failed to serve.")
		}
	}()
	conn, err := grpc.DialContext(
		context.Background(),
		aCtx.Config.GRPCPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		aCtx.Logger.Fatal().Err(err).Msg("failed to dial server.")
	}
	aCtx.Logger.Info().Msg("server dialed.")
	defer conn.Close()
	GMux := runtime.NewServeMux()
	err = grpc_domain.RegisterLinkServiceHandler(context.Background(), GMux, conn)
	if err != nil {
		aCtx.Logger.Fatal().Err(err).Msg("gateway register failed.")
	}
	GServer := &http.Server{
		Handler: GMux,
		Addr:    aCtx.Config.HTTPPort,
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				aCtx.Logger.Fatal().Err(err).Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()
		aCtx.Server.GracefulStop()
		if err := GServer.Shutdown(shutdownCtx); err != nil {
			aCtx.Logger.Fatal().Err(err).Msg("gwserver shutdown error.")
		}
		serverStopCtx()
	}()
	aCtx.Logger.Info().Msg("serving grpc: " + aCtx.Config.HTTPPort)
	if err = GServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			aCtx.Logger.Warn().Err(err).Msg("server closed: ")
			os.Exit(0)
		}
		aCtx.Logger.Fatal().Err(err).Msg("failed to listen and serve.")
	}
}
