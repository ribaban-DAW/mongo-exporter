package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *database) Connect() error {
	uri, err := getUri()
	if err != nil {
		return err
	}

	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(db.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Error trying to connect to MongoDB. Reason: %v", err)
		return err
	}
	log.Println("Trying to connect to MongoDB")

	if err := client.Ping(db.Ctx, nil); err != nil {
		log.Printf("Error connecting to MongoDB. Reason: %v", err)
		return err
	}
	log.Println("Connected to MongoDB")

	db.Client = client
	return nil
}
