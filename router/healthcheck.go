package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHealthcheckRoutes(rg *gin.RouterGroup) {
	h := rg.Group("/healthcheck")

	h.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	log.Println("Added route /healthcheck")
}
