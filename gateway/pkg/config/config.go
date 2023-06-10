package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"GW_PORT"`
	UserServiceURL  string `mapstructure:"USR_URL"`
	EmailServiceURL string `mapstructure:"EMAIL_URL"`
	AdServiceURL    string `mapstructure:"AD_URL"`
}

func Load() *Config {
	var config Config
	viper.AddConfigPath("/")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read .env file %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("cannot unmarshal config into struct: %v", err)
	}

	return &config
}
