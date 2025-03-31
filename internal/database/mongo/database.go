package mongo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

var singleInstance *database

func GetInstance() *database {
	if singleInstance != nil {
		return singleInstance
	}
	singleInstance = &database{}
	return singleInstance
}

func getUri() (string, error) {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return "", errors.New("DB_HOST not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return "", errors.New("DB_PORT not set")
	}

	return fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort), nil
}
