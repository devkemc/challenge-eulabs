package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

type Schema struct {
	Environment string `env:"environment"`
	HttpPort    int    `env:"http_port"`
	DatabaseURI string `env:"database_uri"`
}

var (
	cfg Schema
)

func LoadConfig() *Schema {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil {
		log.Printf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}
	return &cfg
}
func GetConfig() *Schema {
	return &cfg
}
