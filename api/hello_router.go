package api

import (
	"github.com/SrVariable/mongo-exporter/internal/hello"

	"github.com/gin-gonic/gin"
)

func addHelloRoutes(rg *gin.RouterGroup) {
	h := rg.Group("/hello")

	h.GET("/", hello.SayHello)
}
