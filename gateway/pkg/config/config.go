package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	UserServiceURL  string
	EmailServiceURL string
	AdServiceURL    string
	AuthServiceURL  string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		Port:            os.Getenv("GW_PORT"),
		UserServiceURL:  os.Getenv("USR_URL"),
		EmailServiceURL: os.Getenv("EMAIL_URL"),
		AuthServiceURL:  os.Getenv("AUTH_URL"),
		AdServiceURL:    os.Getenv("AD_URL"),
	}
}
