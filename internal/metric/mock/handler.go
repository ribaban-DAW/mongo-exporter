package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
)

func GetOpCountersHandlerMock(ms service.MetricService, c *gin.Context) {
	m, err := ms.FindOpCounters(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}

// Same as GetMetricByNameHandler and used for testing purposes.
// It receives the context as parameter and doesn't return anything.
func GetOpCounterByNameHandlerMock(ms service.MetricService, c *gin.Context) {
	name := c.Param("name")
	m, err := ms.FindOpCounterByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
