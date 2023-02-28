package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnvFile() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	loadEnvFile()

	return os.Getenv(key)
}
