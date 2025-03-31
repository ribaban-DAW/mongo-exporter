package main

import (
	"fmt"

	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	"github.com/SrVariable/mongo-exporter/router"
)

// TODO: Abstract the database more so I only have to db.Connect() and db.Disconect()
func main() {
	db := mongo.GetInstance()
	if err := db.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Disconnect()

	router.Run(":8080")
}
