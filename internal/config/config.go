package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env               string        `env:"ENV"`
	ServerAddr        string        `env:"SERVER_ADDRESS"`
	ServerTimeout     time.Duration `env:"SERVER_TIMEOUT"`
	ServerIdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT"`
	StorageHost       string        `env:"STORAGE_HOST"`
	StoragePort       int           `env:"STORAGE_PORT"`
	StorageUser       string        `env:"STORAGE_USER"`
	StoragePassword   string        `env:"STORAGE_PASSWORD"`
	StorageDB         string        `env:"STORAGE_DATABASE"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}

	return &cfg
}
