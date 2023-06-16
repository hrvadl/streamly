package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port: os.Getenv("AUTH_PORT"),
	}
}
