package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
}

var AppConfig Config

func LoadConfig() {
	AppConfig.ServerPort = getEnv("SERVER_PORT", "8080")
	log.Println("Configuration loaded.")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
