package metric

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockhand "github.com/SrVariable/mongo-exporter/api/metric/mock"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	mockrepo "github.com/SrVariable/mongo-exporter/internal/metric/repository/mock"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetConnections(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	connections := value_object.Connections{
		Available:    domain.Metric[int32]{Value: 10000},
		Current:      domain.Metric[int32]{Value: 50000},
		TotalCreated: domain.Metric[int32]{Value: 100000},
		Active:       domain.Metric[int32]{Value: 300},
	}
	repo := mockrepo.NewMockRepository(nil, &connections, nil, nil, nil)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/v1/metrics/connections", nil)
	mockhand.GetConnectionsHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := connections
	var got value_object.Connections
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.Available.Value, got.Available.Value)
}
