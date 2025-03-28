package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseRepository struct {
	Client *mongo.Client
}

func NewDatabaseRepository(client *mongo.Client) *DatabaseRepository {
	return &DatabaseRepository{
		Client: client,
	}
}
