package config

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func NewCfgDB() Config {
	cfg := Config{
		Host:     os.Getenv("POSTGRES_ORDERS_HOST"),
		Port:     os.Getenv("POSTGRES_ORDERS_PORT"),
		User:     os.Getenv("POSTGRES_ORDERS_USER"),
		Password: os.Getenv("POSTGRES_ORDERS_PASSWORD"),
		Database: os.Getenv("POSTGRES_ORDERS_DB"),
		SSLMode:  "disable",
	}
	return cfg
}
