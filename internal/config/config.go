package config

import (
	"flag"
	"os"
)

type Config struct {
	FlagRunAddr string
}

func New() *Config {
	config := &Config{}
	var flagRunAddr string
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}

	config.FlagRunAddr = flagRunAddr
	return config
}
