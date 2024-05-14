package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Host string `envconfig:"HTTP_HOST" default:"0.0.0.0"`
		Port int    `envconfig:"HTTP_PORT" default:"8080"`
	}
	Storage struct {
		User     string `envconfig:"POSTGRES_USER"`
		Password string `envconfig:"POSTGRES_PASSWORD"`
		Host     string `envconfig:"POSTGRES_HOST"`
		Port     string `envconfig:"POSTGRES_PORT"`
		DBName   string `envconfig:"POSTGRES_DB"`
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("unmarshall config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) Address() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Storage.Host,
		c.Storage.Port,
		c.Storage.User,
		c.Storage.Password,
		c.Storage.DBName,
	)
}
