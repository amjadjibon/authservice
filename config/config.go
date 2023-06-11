package config

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	Host     string `env:"HOST" envDefault:"0.0.0.0"`
	Port     string `env:"PORT" envDefault:"8080"`
	DSN      string `env:"DSN"`
	Store    string `env:"STORE" envDefault:"postgres"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
