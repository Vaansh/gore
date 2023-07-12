package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Getenv(k string, mustGet bool) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	v := os.Getenv(k)
	if mustGet && v == "" {
		log.Fatalf("Fatal Error: %s environment variable not set.\n", k)
	}
	return v
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
