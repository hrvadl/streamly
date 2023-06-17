package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DBConnectionString string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port:               os.Getenv("AUTH_PORT"),
		DBConnectionString: os.Getenv("AUTH_DB_CONNECTION_STRING"),
	}
}
