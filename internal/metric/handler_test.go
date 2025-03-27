package metric

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMetrics(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCountersHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics
	var got []Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(want), len(got))
	for i := range metrics {
		assert.Equal(t, want[i].Name, got[i].Name)
		assert.Equal(t, want[i].Value, got[i].Value)
	}
}

func TestGetMetricsByName_Insert(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "insert"})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[0]
	var got Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetMetricsByName_Delete(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "delete"})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[1]
	var got Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetMetricsByName_Query(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "query"})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[2]
	var got Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetMetricsByName_Update(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "update"})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusOK, w.Code)

	want := metrics[3]
	var got Metric
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Value, got.Value)
}

func TestGetMetricsByName_Empty(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: ""})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "metric not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}

func TestGetMetricsByName_NoExist(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "foo"})

	metrics := []Metric{
		{Name: "insert", Value: "1", Timestamp: time.Now()},
		{Name: "delete", Value: "2", Timestamp: time.Now()},
		{Name: "query", Value: "3", Timestamp: time.Now()},
		{Name: "update", Value: "4", Timestamp: time.Now()},
	}

	repo := NewMockRepository(metrics)
	service := NewMetricService(repo)

	c.Request, _ = http.NewRequest(http.MethodGet, "/metrics", nil)
	GetOpCounterByNameHandlerMock(service, c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	want := gin.H{"message": "metric not found"}
	var got gin.H
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}
