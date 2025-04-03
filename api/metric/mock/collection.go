package service

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetCollectionHandlerMock(ms *service.MetricService, c *gin.Context) {
	collection, err := ms.FindCollection(c.Request.Context(), "foo", "bar")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, collection)
}
