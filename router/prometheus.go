package router

import (
	"log"

	prom "github.com/SrVariable/mongo-exporter/api/metric/prometheus"
	"github.com/SrVariable/mongo-exporter/internal/database/mongo"
	repo "github.com/SrVariable/mongo-exporter/internal/metric/repository/mongo"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func addPrometheusRoute(rg *gin.RouterGroup) {
	m := rg.Group("/prometheus")

	client := mongo.GetInstance().Client
	repo := repo.NewDatabaseRepository(client)
	service := service.NewMetricService(repo)

	prom.RecordCollection(service)
	prom.RecordConnections(service)
	prom.RecordCpu(service)
	prom.RecordRam(service)
	prom.RecordOpCounters(service)

	m.GET("/", gin.WrapH(promhttp.Handler()))
	log.Println("Added route /prometheus")
}
