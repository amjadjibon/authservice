package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	Host     string `env:"HOST" envDefault:"0.0.0.0"`
	Port     string `env:"PORT" envDefault:"8080"`
	DSN      string `env:"DSN"`
	Store    string `env:"STORE" envDefault:"postgres"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`

	// Email
	EmailHost     string `env:"EMAIL_HOST" envDefault:"smtp.gmail.com"`
	EmailPort     string `env:"EMAIL_PORT" envDefault:"587"`
	EmailUsername string `env:"EMAIL_USERNAME"`
	EmailPassword string `env:"EMAIL_PASSWORD"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
