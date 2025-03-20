package main

import (
	router "github.com/SrVariable/mongo-exporter/api"
)

func main() {
	router.Run("localhost:8080")
}
