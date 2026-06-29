package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DatabaseUrl        string
	RedisUrl           string
	GoogleClientID     string
	GoogleClientSecret string
	GoolgeCallbackUrl  string
	CloudinaryUrl      string
}

var Configuration *Config

// @Summary 		Load Configurations
// @Description		Getting and storing the Configurations from ENV
// @Tags			Config
// @Accept 			*/*
// @Produce 		plain
func FetchConfig() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load env values", "error", err)
		os.Exit(1)
	}
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		slog.Error("No Database url to connect the database to server")
		os.Exit(2)
	}
	return &Config{
		Port:               os.Getenv("PORT"),
		DatabaseUrl:        databaseUrl,
		RedisUrl:           os.Getenv("REDIS_URL"),
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoolgeCallbackUrl:  os.Getenv("GOOGLE_CALLBACK_URL"),
		CloudinaryUrl:      os.Getenv("CLOUDINARY_URL"),
	}
}
