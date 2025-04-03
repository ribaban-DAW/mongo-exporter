package router

import (
	"log"

	"github.com/SrVariable/mongo-exporter/config"
	"github.com/gin-gonic/gin"
)

func Run(env *config.Env) (err error) {
	router := setupRouter()
	uri := ":" + env.AppPort
	return router.Run(uri)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	log.Println("Created router")

	v1 := router.Group("/v1")

	addHealthcheckRoutes(v1)
	addHelloRoutes(v1)
	addMetricRoutes(v1)
	addPrometheusRoute(v1)

	return router
}
