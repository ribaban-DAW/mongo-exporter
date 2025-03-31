package metric

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	mockrepo "github.com/SrVariable/mongo-exporter/internal/metric/mock/repository"
	mockserv "github.com/SrVariable/mongo-exporter/internal/metric/mock/service"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCpuUsage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	metrics := []domain.Metric{
		{Name: "userTime", Value: "10000", Timestamp: time.Now()},
		{Name: "systemTime", Value: "5000", Timestamp: time.Now()},
	}

	repo := mockrepo.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics/cpu", nil)
	mockserv.GetCpuUsageHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := append(
		metrics,
		domain.Metric{
			Name:      "totalTime",
			Value:     "15000",
			Timestamp: time.Now(),
		},
	)
	var got []domain.Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(want), len(got))
	for i := range metrics {
		assert.Equal(t, want[i].Name, got[i].Name)
		assert.Equal(t, want[i].Value, got[i].Value)
	}
}
