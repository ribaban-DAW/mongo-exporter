package metric

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetOpCountersHandler(ms *service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindOpCounters(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
