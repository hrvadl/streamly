package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DNS  string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Print(err.Error())
	}

	return &Config{
		Port: os.Getenv("AUTH_PORT"),
		DNS:  os.Getenv("AUTH_DB_DNS"),
	}
}
