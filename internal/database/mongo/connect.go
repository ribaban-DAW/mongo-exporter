package mongo

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *database) Connect() error {
	uri := fmt.Sprintf("mongodb://%s:%s", db.Config.Host, db.Config.Port)

	client, err := mongo.Connect(db.Context, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Error trying to connect to MongoDB. Reason: %v", err)
		return err
	}
	log.Println("Trying to connect to MongoDB")

	if err := client.Ping(db.Context, nil); err != nil {
		log.Printf("Error connecting to MongoDB. Reason: %v", err)
		return err
	}
	log.Println("Connected to MongoDB")

	db.Client = client
	return nil
}
