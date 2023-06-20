package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EmailApiKey      string
	BootstrapServers string
	GroupID          string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("cannot load env file %v", err)
	}

	return &Config{
		EmailApiKey:      os.Getenv("EMAIL_API_KEY"),
		BootstrapServers: os.Getenv("BOOTSTRAP_SERVERS"),
		GroupID:          os.Getenv("EMAIL_GROUP_ID"),
	}
}
