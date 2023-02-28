package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
