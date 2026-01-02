package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/oganes5796/shops-list/internal/orders/api"
	"github.com/oganes5796/shops-list/internal/orders/client/db"
	"github.com/oganes5796/shops-list/internal/orders/config"
	"github.com/oganes5796/shops-list/internal/orders/repository/repo"
	"github.com/oganes5796/shops-list/internal/orders/server"
	"github.com/oganes5796/shops-list/internal/orders/service/serv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	cfg := config.NewCfgDB()
	pool, err := db.NewPostgresDB(ctx, cfg)
	if err != nil {
		slog.Error("Error connecting to PostgreSQL", "error", err)
		os.Exit(1)
	}
	defer pool.Close(ctx)
	slog.Info("Successfully connected to PostgreSQL")

	repos := repo.NewRepository(pool)
	services := serv.NewService(repos)
	handlers := api.NewImplementation(services)

	srv := &server.Server{}
	go func() {
		if err := srv.Run(
			os.Getenv("HOST"),
			os.Getenv("PORT_ORDERS"),
			handlers.InitRoutes(),
		); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("error occurred while running http server", "error", err)
		}
	}()
	slog.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("App shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error occurred on server shutting down", "error", err)
	}

	slog.Info("App exited")
}
