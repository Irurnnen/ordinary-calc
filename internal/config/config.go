package config

import (
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
		// TODO: Add logs
		return NewConfigExample()
	}
	return &Config{
		Port: port,
	}
}
