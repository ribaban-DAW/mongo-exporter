package api

import (
	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	"github.com/SrVariable/mongo-exporter/internal/metric"
	"github.com/gin-gonic/gin"
)

func addMetricRoutes(rg *gin.RouterGroup) {
	m := rg.Group("/metrics")

	client := mongo.GetInstance().Client
	repo := metric.NewDatabaseRepository(client)
	service := metric.NewMetricService(repo)

	m.GET("/", func(c *gin.Context) {
		metric.GetAvailableMetricsHandler(service)(c)
	})

	m.GET("/opcounters", func(c *gin.Context) {
		metric.GetOpCountersHandler(service)(c)
	})

	m.GET("/:name", func(c *gin.Context) {
		metric.GetOpCounterByNameHandler(service)(c)
	})

	m.GET("/cpu", func(c *gin.Context) {
		metric.GetCpuUsageHandler(service)(c)
	})

	m.GET("/ram", func(c *gin.Context) {
		metric.GetRamUsageHandler(service)(c)
	})
}
