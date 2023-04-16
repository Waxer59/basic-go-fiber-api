package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}

	return os.Getenv(key)
}
