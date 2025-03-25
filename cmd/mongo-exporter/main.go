package main

import (
	"fmt"

	router "github.com/SrVariable/mongo-exporter/api"
	"github.com/SrVariable/mongo-exporter/internal/database"
)

func main() {
	db := database.GetInstance()
	if err := db.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Disconnect()

	router.Run(":8080")
}
