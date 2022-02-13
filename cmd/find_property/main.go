package main

import (
	"context"
	"flag"
	"github.com/find_property/internal/config"
	v1 "github.com/find_property/internal/delivery/v1"
	"github.com/find_property/internal/repository"
	"github.com/find_property/internal/server"
	"github.com/find_property/internal/service"
	"github.com/find_property/pkg/database/postgres"
	"github.com/find_property/pkg/logger"
	"github.com/find_property/pkg/worker"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Swagger .....
// @version 0.1.1
// @description Api server for ....

// @host localhost:8080
// @BasePath /
func main() {
	configPath := new(string)

	flag.StringVar(configPath, "config-path", "config/config-local.yaml", "specify path to yaml")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with os.Open config"))
	}

	cfg := config.Config{}
	if err := yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Decode config"))
	}

	if err = logger.NewLogger(cfg.Telegram); err != nil {
		logger.LogFatal(err)
	}

	newWorker := worker.New()
	if err = newWorker.StartWorker(); err != nil {
		logger.LogFatal(err)
	}

	postgresClient, err := postgres.NewPostgres(cfg.PostgresDsn)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewPostgres"))
	}

	db, err := postgresClient.Database()
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Gorm"))
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(repos, newWorker)

	endpoints := v1.NewHandler(services)

	srv := server.NewServer(&cfg, endpoints)

	go func() {

		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
		}

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	newWorker.Clear()

	if err = srv.Shutdown(ctx); err != nil {
		logger.LogFatal(errors.Wrap(err, "failed to stop server"))
	}

	if err = postgresClient.Close(); err != nil {
		logger.LogFatal(errors.Wrap(err, "failed to stop db"))
	}
}
