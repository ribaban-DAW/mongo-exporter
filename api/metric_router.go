package api

import (
	"github.com/SrVariable/mongo-exporter/internal/database"
	"github.com/SrVariable/mongo-exporter/internal/metric"
	"github.com/gin-gonic/gin"
)

func addMetricRoutes(rg *gin.RouterGroup) {
	m := rg.Group("/metrics")

	client := database.GetInstance().Client
	repo := metric.NewDatabaseRepository(client)
	service := metric.NewMetricService(repo)

	m.GET("/", func(cg *gin.Context) {
		metric.GetMetricsHandler(service)(cg)
	})

	m.GET("/:name", func(cg *gin.Context) {
		metric.GetMetricByNameHandler(service)(cg)
	})
}
