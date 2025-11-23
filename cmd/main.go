package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"

	handlerNumbers "github.com/avraam311/profiled-number-adder/internal/api/http/handlers/numbers"
	"github.com/avraam311/profiled-number-adder/internal/api/http/server"
	"github.com/avraam311/profiled-number-adder/internal/infra/config"
	"github.com/avraam311/profiled-number-adder/internal/infra/logger"
	serviceNumbers "github.com/avraam311/profiled-number-adder/internal/service/numbers"
)

const (
	configFilePath = "config/local.yaml"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger.Init()
	cfg := config.New()
	if err := cfg.LoadConfigFiles(configFilePath); err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to load config file")
	}

	srvcNums := serviceNumbers.New()
	handNums := handlerNumbers.New(srvcNums)

	router := server.NewRouter(handNums)
	srv := server.NewServer(cfg.GetString("server.port"), router)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Logger.Fatal().Err(err).Msg("failed to run server")
		}
	}()
	logger.Logger.Info().Msg("server is running")

	<-ctx.Done()
	logger.Logger.Info().Msg("shutdown signal received")

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	logger.Logger.Info().Msg("shutting down")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to shutdown server")
	}
	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		logger.Logger.Info().Msg("timeout exceeded, forcing shutdown")
	}
}
