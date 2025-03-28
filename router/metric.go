package router

import (
	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	repo "github.com/SrVariable/mongo-exporter/internal/metric/repository/mongo"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/SrVariable/mongo-exporter/api/metric"
	"github.com/gin-gonic/gin"
)

func addMetricRoutes(rg *gin.RouterGroup) {
	m := rg.Group("/metrics")

	client := mongo.GetInstance().Client
	repo := repo.NewDatabaseRepository(client)
	service := service.NewMetricService(repo)

	m.GET("/", func(c *gin.Context) {
		metric.GetSummaryHandler(service)(c)
	})

	m.GET("/opcounters", func(c *gin.Context) {
		metric.GetOpCountersHandler(service)(c)
	})

	m.GET("/opcounters/:name", func(c *gin.Context) {
		metric.GetOpCounterByNameHandler(service)(c)
	})

	m.GET("/cpu", func(c *gin.Context) {
		metric.GetCpuUsageHandler(service)(c)
	})

	m.GET("/ram", func(c *gin.Context) {
		metric.GetRamUsageHandler(service)(c)
	})
}
