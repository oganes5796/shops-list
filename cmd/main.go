package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/oganes5796/shops-list/internal/client/db"
	"github.com/oganes5796/shops-list/internal/config"

	reposhop "github.com/oganes5796/shops-list/internal/repository/shop"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	cfg := config.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		SSLMode:  "disable",
	}
	pool, err := db.NewPostgresDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close(ctx)
	slog.Info("Successfully connected to PostgreSQL")

	_ = reposhop.NewRepository(pool)

}
