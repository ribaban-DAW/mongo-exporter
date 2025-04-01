package router

import (
	"log"

	"github.com/SrVariable/mongo-exporter/api/metric"
	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	repo "github.com/SrVariable/mongo-exporter/internal/metric/repository/mongo"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
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
	log.Println("Added route /metrics")

	m.GET("/opcounters", func(c *gin.Context) {
		metric.GetOpCountersHandler(service)(c)
	})
	log.Println("Added route /metrics/opcounters")

	m.GET("/cpu", func(c *gin.Context) {
		metric.GetCpuHandler(service)(c)
	})
	log.Println("Added route /metrics/cpu")

	m.GET("/ram", func(c *gin.Context) {
		metric.GetRamHandler(service)(c)
	})
	log.Println("Added route /metrics/ram")

	m.GET("/connections", func(c *gin.Context) {
		metric.GetConnectionsHandler(service)(c)
	})
	log.Println("Added route /metrics/connections")
}
