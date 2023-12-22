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
	"github.com/Spargwy/just-to-do-it/pkg/auth/bcrypt"
	"github.com/Spargwy/just-to-do-it/pkg/auth/jwt"
	"github.com/Spargwy/just-to-do-it/pkg/config"
	"github.com/Spargwy/just-to-do-it/pkg/logger"
	"github.com/sirupsen/logrus"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig(".")

	logger.New(cfg.Logger.Level)
	db, err := postgres.NewPostgres(cfg.Database)
	if err != nil {
		logrus.Fatalf("NewPostgres: %v", err)
	}

	jwt := jwt.New(cfg.Authenticator.JwtPath)

	encrypter := bcrypt.New(cfg.Encrypter.Cost)

	executor := service.New(db, encrypter, jwt)

	server := api.New(&executor, jwt)

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
