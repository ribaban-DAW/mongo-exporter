package mongo

import (
	"context"

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
