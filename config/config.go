package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT                  int
	DISCOVERY_SERVICE_PORT     string
	PODCAST_SERVICE_PORT       string
	COLLABORATION_SERVICE_PORT string
	USER_SERVICE_PORT          string
	DB_HOST                    string
	DB_PORT                    string
	DB_USER                    string
	DB_NAME                    string
	DB_PASSWORD                string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToInt(coalesce("HTTP_PORT", 8080))
	config.DISCOVERY_SERVICE_PORT = cast.ToString(coalesce("DISCOVERY_SERVICE_PORT", 50051))
	config.USER_SERVICE_PORT = cast.ToString(coalesce("USER_SERVICE_PORT", 50052))
	config.PODCAST_SERVICE_PORT = cast.ToString(coalesce("PODCAST_SERVICE_PORT", 50053))
	config.COLLABORATION_SERVICE_PORT = cast.ToString(coalesce("COLLABORATION_SERVICE_PORT", 50054))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "listenup_discovery_service"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))

	return &config
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
