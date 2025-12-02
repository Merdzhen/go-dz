package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("KEY is not set in environment")
	}

	return &Config{
		Key: key,
	}
}
