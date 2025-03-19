package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DEFAULT_MONGO_URI = "mongodb://localhost:27017"
const TIMEOUT = 30 * time.Second

func connectDatabase() (
	*mongo.Client,
	context.Context,
	context.CancelFunc,
	error,
) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = DEFAULT_MONGO_URI
	}
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	fmt.Printf("Connecting to %s\n", uri)
	return client, ctx, cancel, err
}

func closeConnection(
	client *mongo.Client,
	ctx context.Context,
	cancel context.CancelFunc,
) {
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

func insertOne(
	client *mongo.Client,
	ctx context.Context,
	database, col string,
	doc interface{},
) (
	*mongo.InsertOneResult,
	error,
) {
	collection := client.Database(database).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("Insertion successful")
	}
	return result, err
}

func insertMany(
	client *mongo.Client,
	ctx context.Context,
	database, col string,
	docs  []interface{},
) (
	*mongo.InsertManyResult,
	error,
) {
	collection := client.Database(database).Collection(col)
	result, err := collection.InsertMany(ctx, docs)
	if err != nil {
		fmt.Println("Insertion successful")
	}
	return result, err
}

func query(client *mongo.Client,
	ctx context.Context,
	database, col string,
	query, field interface{},
) (*mongo.Cursor,
	error,
) {
	collection := client.Database(database).Collection(col)
	result, err := collection.Find(ctx, query, options.Find().SetProjection(field))
	return result, err
}

func main() {
	client, ctx, cancel, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	ping(client, ctx)
	defer closeConnection(client, ctx, cancel)

	// CREATE

	document := bson.D {
		{"id", 1},
		{"name", "Foo"},
		{"age", 22},
	}
	resultInsertOne, err := insertOne(client, ctx, "test", "student", document)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultInsertOne.InsertedID)

	documents := []interface{} {
		bson.D {
			{"id", 2},
			{"name", "Bar"},
			{"age", 23},
		},
		bson.D {
			{"id", 3},
			{"name", "Baz"},
			{"age", 20},
		},
	}
	resultInsertMany, err := insertMany(client, ctx, "test", "student", documents)
	if err != nil {
		panic(err)
	}
	for id := range resultInsertMany.InsertedIDs {
		fmt.Println(resultInsertMany.InsertedIDs[id])
	}

	// READ

	filter := bson.D {
		{"age", 20},
	}
	option := bson.D{
		{"_id", 0},
	}
	cursor, err := query(client, ctx, "test", "student", filter, option)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err := cursor.All(ctx, &results); err != nil  {
		panic(err)
	}

	fmt.Println("Query result age == 20:")
	for _, doc := range results {
		fmt.Println(doc)
	}

	filter = bson.D {
		{"age", bson.D{{"$gt", 20}}},
	}
	option = bson.D{
		{"_id", 0},
	}
	cursor, err = query(client, ctx, "test", "student", filter, option)
	if err != nil {
		panic(err)
	}
	if err := cursor.All(ctx, &results); err != nil  {
		panic(err)
	}

	fmt.Println("Query result age > 20:")
	for _, doc := range results {
		fmt.Println(doc)
	}

	filter = bson.D {
		{"name", bson.D{{"$regex", "r$"}}},
	}
	option = bson.D{
		{"_id", 0},
	}
	cursor, err = query(client, ctx, "test", "student", filter, option)
	if err != nil {
		panic(err)
	}
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	fmt.Println("Query result name ends with 'r'")
	for _, doc := range results {
		fmt.Println(doc)
	}

	filter = bson.D {
		{"name", bson.D{{"$regex", "^B"}}},
	}
	option = bson.D{
		{"_id", 0},
	}
	cursor, err = query(client, ctx, "test", "student", filter, option)
	if err != nil {
		panic(err)
	}
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	fmt.Println("Query result name starts with 'B'")
	for _, doc := range results {
		fmt.Println(doc)
	}
}
