package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/oganes5796/shops-list/internal/client/db"
	"github.com/oganes5796/shops-list/internal/config"

	apiShop "github.com/oganes5796/shops-list/internal/api/shop"
	repoShop "github.com/oganes5796/shops-list/internal/repository/shop"
	srvShop "github.com/oganes5796/shops-list/internal/server/shop"
	serviceShop "github.com/oganes5796/shops-list/internal/service/shop"
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

	repos := repoShop.NewRepository(pool)
	services := serviceShop.NewService(repos)
	handlers := apiShop.NewImplementation(services)

	srv := &srvShop.Server{}
	go func() {
		if err := srv.Run(os.Getenv("HOST"), os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			slog.Error("error occurred while running http server", "error", err)
			os.Exit(1)
		}
	}()
	slog.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	slog.Info("App shutting down")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error occurred on server shutting down", "error", err)
	}
	slog.Info("App exited")

}
