package router

import (
	"github.com/SrVariable/mongo-exporter/api/hello"
	"github.com/gin-gonic/gin"
)

func addHelloRoutes(rg *gin.RouterGroup) {
	h := rg.Group("/hello")

	h.GET("/", hello.SayHello)
}
