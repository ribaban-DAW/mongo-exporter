package service

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetCpuUsageHandlerMock(ms service.MetricService, c *gin.Context) {
	m, err := ms.FindCpuUsage(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
