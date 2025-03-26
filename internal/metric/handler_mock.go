package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMetricsHandlerMock(ms MetricService, c *gin.Context) {
	m, err := ms.FindMetrics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}

// Same as GetMetricByNameHandler and used for testing purposes.
// It receives the context as parameter and doesn't return anything.
func GetMetricByNameHandlerMock(ms MetricService, c *gin.Context) {
	name := c.Param("name")
	m, err := ms.FindMetricByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
