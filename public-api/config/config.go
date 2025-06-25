package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}
