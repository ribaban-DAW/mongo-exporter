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

func TestGetCpu(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	cpu := value_object.Cpu{
		UserTime:   domain.Metric[int64]{Value: 10000},
		SystemTime: domain.Metric[int64]{Value: 1000},
	}

	repo := mockrepo.NewMockRepository(nil, nil, &cpu, nil, nil)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/v1/metrics/cpu", nil)
	mockhand.GetCpuHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := cpu
	var got value_object.Cpu
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.UserTime.Value, got.UserTime.Value)
	assert.Equal(t, want.SystemTime.Value, got.SystemTime.Value)
	assert.Equal(t, want.TotalTime.Value, got.TotalTime.Value)
}
