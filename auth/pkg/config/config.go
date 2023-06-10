package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"AUTH_PORT"`
}

func Load() *Config {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read .env file %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("cannot unmarshal config into struct: %v", err)
	}

	return &config
}
