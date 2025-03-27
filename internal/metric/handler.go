package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NOTE: I don't know if makes sense
func GetAvailableMetricsHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"endpoints": []string{"cpu", "opcounters", "ram"},
		})
	}
}

func GetOpCountersHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindOpCounters(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}

func GetOpCounterByNameHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		m, err := ms.FindOpCounterByName(c.Request.Context(), name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}

func GetCpuUsageHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindCpuUsage(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}

func GetRamUsageHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindRamUsage(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
