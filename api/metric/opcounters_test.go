package metric

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/SrVariable/mongo-exporter/internal/metric/mock"
)

func TestGetOpCounters(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCountersHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics
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

func TestGetOpCounterByName_Insert(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "insert"})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[0]
	var got domain.Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetOpCounterByName_Delete(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "delete"})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[1]
	var got domain.Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetOpCounterByName_Query(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "query"})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[2]
	var got domain.Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetOpCounterByName_Update(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "update"})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[3]
	var got domain.Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetOpCounterByName_Empty(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: ""})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "metric not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}

func TestGetOpCounterByName_NoExist(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "foo"})

	metrics := []domain.Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := mock.NewMockRepository(metrics)
	service := service.NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	mock.GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "metric not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}
