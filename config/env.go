package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppPort string
	DBHost  string
	DBPort  string
}

func NewEnv() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning couldn't load `.env` file")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Println("APP_PORT not set")
		return nil, errors.New("APP_PORT not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Println("DB_HOST not set")
		return nil, errors.New("DB_HOST not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Println("DB_PORT not set")
		return nil, errors.New("DB_PORT not set")
	}

	config := &Env{
		AppPort: appPort,
		DBHost:  dbHost,
		DBPort:  dbPort,
	}
	log.Println("Config loaded")

	return config, nil
}
