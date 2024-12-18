package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func NewConfigExample() *Config {
	return &Config{
		Port: 8080,
	}
}

func NewConfigFromEnv() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Fatal error while getting config from env: PORT:%s", os.Getenv("PORT"))
		return NewConfigExample()
	}
	return &Config{
		Port: port,
	}
}
