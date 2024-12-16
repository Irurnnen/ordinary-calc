package config

type Config struct {
	Port int
}

func NewConfigExample() *Config {
	return &Config{
		Port: 8080,
	}
}

// TODO: Config from Env

// TODO: Config from config file

// TODO: Config from arguments
