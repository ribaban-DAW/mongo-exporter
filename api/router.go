package api

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run(uri string) (err error) {
	router := setupRouter()
	return router.Run(uri)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/v1")
	addStudentRoutes(v1)
	return router
}
