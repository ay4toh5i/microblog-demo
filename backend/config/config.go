package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
	DbDatabase string `env:"DB_DATABASE"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	HttpPort   int    `env:"HTTP_PORT"`
}

func Init() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
