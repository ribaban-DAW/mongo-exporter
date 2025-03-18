package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

const DEFAULT_MONGO_URI = "mongodb://localhost:27017"
const TIMEOUT = 30 * time.Second

func connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = DEFAULT_MONGO_URI
	}
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	fmt.Printf("Connecting to %s\n", uri)

	return client, ctx, cancel, err
}

func close(client *mongo.Client,
			ctx context.Context,
			cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("Connection closed")
	}()
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connection successful")
	return nil
}

func main() {
	client, ctx, cancel, err := connect()
	if err != nil {
		panic(err)
	}

	ping(client, ctx)
	defer close(client, ctx, cancel)
}
