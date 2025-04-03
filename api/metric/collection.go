package metric

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetCollectionHandler(ms *service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbName := c.DefaultQuery("dbName", "")
		collName := c.DefaultQuery("collName", "")
		m, err := ms.FindCollection(c.Request.Context(), dbName, collName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
