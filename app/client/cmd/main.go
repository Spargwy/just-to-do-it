package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Spargwy/just-to-do-it/app/client/api"
	"github.com/Spargwy/just-to-do-it/app/client/db/postgres"
	"github.com/Spargwy/just-to-do-it/app/client/service"
	"github.com/Spargwy/just-to-do-it/pkg/config"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig(".")

	logger.New(cfg.Logger.Level)
	db, err := postgres.NewPostgres(cfg.Database.Client)
	if err != nil {
		logrus.Fatalf("NewPostgres: %v", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		logrus.Fatalf("Ping DB error: %v", err)
	}

	ct := service.New(db)

	server := api.New(ct)

	go func() {
		if err := server.Start(cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}
