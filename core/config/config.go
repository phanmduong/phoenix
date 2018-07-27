package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	Config	map[string]string
}

func NewConfig() *AppConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	return &AppConfig{
		Config: map[string]string{
			"APP_DEBUG": os.Getenv("APP_DEBUG"),
		},
	}
}