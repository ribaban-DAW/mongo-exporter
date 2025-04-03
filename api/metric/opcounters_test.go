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

func TestGetOpCounters(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	opcounters := value_object.OpCounters{
		Delete: domain.Metric[int64]{Value: 250},
		Insert: domain.Metric[int64]{Value: 500},
		Query:  domain.Metric[int64]{Value: 1000},
		Update: domain.Metric[int64]{Value: 5000},
	}

	repo := mockrepo.NewMockRepository(nil, nil, &opcounters, nil)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/v1/metrics/opcounters", nil)
	mockhand.GetOpCountersHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := opcounters
	var got value_object.OpCounters
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.Delete.Value, got.Delete.Value)
	assert.Equal(t, want.Insert.Value, got.Insert.Value)
	assert.Equal(t, want.Query.Value, got.Query.Value)
	assert.Equal(t, want.Update.Value, got.Update.Value)
}
