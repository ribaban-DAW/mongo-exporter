package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
)

func GetCpuUsageHandler(ms service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindCpuUsage(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
