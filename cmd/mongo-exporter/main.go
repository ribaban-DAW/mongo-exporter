package main

import (
	"context"
	"time"

	"github.com/SrVariable/mongo-exporter/config"
	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	"github.com/SrVariable/mongo-exporter/router"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		return
	}

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	db := mongo.NewDatabase(context, cancel, env)
	if err := db.Connect(); err != nil {
		return
	}
	defer db.Disconnect()

	router.Run(env)
}
