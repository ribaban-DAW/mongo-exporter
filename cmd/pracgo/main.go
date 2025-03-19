package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	defaultHost = "localhost"
	defaultPort = "8080"
	defaultUri = defaultHost + ":" + defaultPort
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world")
}

func main() {
	router := gin.Default()
	router.GET("/", helloWorld)

	if err := router.Run(defaultUri); err != nil {
		log.Fatal(err)
		return
	}
}
