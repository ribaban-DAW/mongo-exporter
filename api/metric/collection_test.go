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

func TestGetCollection(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	collection := value_object.Collection{
		Remove:  domain.Metric[int32]{Value: 59},
		Insert:  domain.Metric[int32]{Value: 123121},
		Queries: domain.Metric[int32]{Value: 2},
		Update:  domain.Metric[int32]{Value: 3},
	}

	repo := mockrepo.NewMockRepository(&collection, nil, nil, nil, nil)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/v1/metrics/collection", nil)
	mockhand.GetCollectionHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := collection
	var got value_object.Collection
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.Remove.Value, got.Remove.Value)
	assert.Equal(t, want.Insert.Value, got.Insert.Value)
	assert.Equal(t, want.Queries.Value, got.Queries.Value)
	assert.Equal(t, want.Update.Value, got.Update.Value)
}
