package database

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Connect() error
	Disconnect() error
}

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

func (db *database) Connect() error {
	if db.Client != nil {
		return errors.New("client already connected")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return errors.New("MONGO_URI not set")
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

func (db *database) Disconnect() error {
	if db.Client == nil {
		return errors.New("client is not connected")
	}
	defer db.Cancel()
	defer func() {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			panic(err)
		}
	}()

	db.Client = nil
	return nil
}
