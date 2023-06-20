package config

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DNS            string
	UserServiceURL string
	TokenIssuer    string
	TokenAudience  string
	JwtKey         string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Print(err.Error())
	}

	return &Config{
		Port:           os.Getenv("AUTH_PORT"),
		DNS:            os.Getenv("AUTH_DB_DNS"),
		UserServiceURL: os.Getenv("USR_URL"),
		TokenIssuer:    os.Getenv("AUTH_ISSUER"),
		TokenAudience:  os.Getenv("AUTH_AUDIENCE"),
		JwtKey:         base64.StdEncoding.EncodeToString([]byte(os.Getenv("AUTH_JWT_KEY"))),
	}
}
