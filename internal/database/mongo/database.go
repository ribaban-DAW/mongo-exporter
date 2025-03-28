package mongo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (db *database) Connect() error {
	uri, err := getUri()
	if err != nil {
		return err
	}

	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(db.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err := client.Ping(db.Ctx, nil); err != nil {
		return err
	}

	db.Client = client
	return nil
}

func (db *database) Disconnect() {
	defer db.Cancel()
	defer func() {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			panic(err)
		}
	}()
}
