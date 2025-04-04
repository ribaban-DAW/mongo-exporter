package metric

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetSummaryHandler(ms *service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"available": []string{"collection", "connections", "cpu", "opcounters", "ram"},
			},
		)
	}
}
