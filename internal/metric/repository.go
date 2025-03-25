package metric

import (
	"context"
	"errors"
	"fmt"

	"github.com/SrVariable/mongo-exporter/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MetricRepository interface {
	GetMetrics(c context.Context) ([]Metric, error)
	GetMetricByName(c context.Context, name string) (*Metric, error)
}

type DatabaseRepository struct {
	Client *mongo.Client
}

func NewDatabaseRepository(client *mongo.Client) *DatabaseRepository {
	return &DatabaseRepository{
		Client: client,
	}
}

// It currently shows `opcounters`, but can be modified
// depending on the specific metrics we want to get
func (dr *DatabaseRepository) GetMetrics(c context.Context) ([]Metric, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	if err != nil {
		return nil, err
	}

	var metrics []Metric

	opcounters, ok := result["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	for metricName, metricValue := range opcounters {
		metric := Metric{
			Name:      metricName,
			Value:     fmt.Sprintf("%d", metricValue),
			Timestamp: utils.GetTime(),
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (dr *DatabaseRepository) GetMetricByName(c context.Context, name string) (*Metric, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	if err != nil {
		return nil, err
	}

	opcounters, ok := result["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}

	value, ok := opcounters[name]
	if !ok {
		return nil, errors.New("metric not found")
	}

	metric := Metric{
		Name:      name,
		Value:     fmt.Sprintf("%d", value),
		Timestamp: utils.GetTime(),
	}
	return &metric, nil
}

// MockRepository for testing
type MockRepository struct {
	Metrics []Metric
}

func NewMockRepository(metrics []Metric) *MockRepository {
	return &MockRepository{
		Metrics: metrics,
	}
}

func (mr *MockRepository) GetMetrics(c context.Context) ([]Metric, error) {
	return mr.Metrics, nil
}

func (mr *MockRepository) GetMetricByName(c context.Context, name string) (*Metric, error) {
	for _, metric := range mr.Metrics {
		if metric.Name == name {
			return &metric, nil
		}
	}
	return nil, errors.New("metric not found")
}
