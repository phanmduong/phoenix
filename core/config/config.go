package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type AppConfig struct {
	Config map[string]string
}

var instance AppConfig
var once sync.Once

func GetAppConfig() AppConfig {
	once.Do(func() {
		instance = NewConfig()
	})
	return instance
}

func NewConfig() AppConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	return AppConfig{
		Config: map[string]string{
			"DB_NAME" : os.Getenv("DB_NAME"),
			"DB_HOST" : os.Getenv("DB_HOST"),
			"DB_USER": os.Getenv("DB_USER"),
			"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
			"DB_PORT": os.Getenv("DB_PORT"),
			"APP_DEBUG": os.Getenv("APP_DEBUG"),
			"SECRET_JWT": os.Getenv("SECRET_JWT"),
		},
	}
}
