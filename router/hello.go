package router

import (
	"log"

	"github.com/SrVariable/mongo-exporter/api/hello"
	"github.com/gin-gonic/gin"
)

func addHelloRoutes(rg *gin.RouterGroup) {
	h := rg.Group("/hello")

	h.GET("/", hello.SayHello)
	log.Println("Added route /hello")
}
