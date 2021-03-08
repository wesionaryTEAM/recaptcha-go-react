package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

//load env initially
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
