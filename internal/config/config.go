package config

import (
	"cmp"
	"os"
)

type Config struct {
	Port string
}

func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	return cmp.Or(os.Getenv(key), fallback)
}
