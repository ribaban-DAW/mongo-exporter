package mongo

import (
	"context"

	"github.com/SrVariable/mongo-exporter/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbConfig struct {
	Host string
	Port string
}

type database struct {
	Cancel  context.CancelFunc
	Context context.Context
	Client  *mongo.Client
	Config  dbConfig
}

var singleInstance *database

func NewDatabase(context context.Context, cancel context.CancelFunc, env *config.Env) *database {
	if singleInstance != nil {
		return singleInstance
	}

	config := dbConfig{
		Host: env.DBHost,
		Port: env.DBPort,
	}

	singleInstance = &database{
		Context: context,
		Cancel:  cancel,
		Config:  config,
	}
	return singleInstance
}

func GetInstance() *database {
	if singleInstance == nil {
		panic("Database not initialized")
	}
	return singleInstance
}
