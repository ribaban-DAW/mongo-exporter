package router

import "github.com/gin-gonic/gin"

func Run(uri string) (err error) {
	router := setupRouter()
	return router.Run(uri)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/v1")

	addHealthcheckRoutes(v1)
	addHelloRoutes(v1)
	addMetricRoutes(v1)

	return router
}
