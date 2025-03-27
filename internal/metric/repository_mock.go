package metric

import (
	"context"
	"errors"
)

type MockRepository struct {
	Metrics []Metric
}

func NewMockRepository(metrics []Metric) *MockRepository {
	return &MockRepository{
		Metrics: metrics,
	}
}

func (mr *MockRepository) GetOpCounters(c context.Context) ([]Metric, error) {
	return mr.Metrics, nil
}

func (mr *MockRepository) GetOpCounterByName(c context.Context, name string) (*Metric, error) {
	for _, metric := range mr.Metrics {
		if metric.Name == name {
			return &metric, nil
		}
	}
	return nil, errors.New("metric not found")
}

func (mr *MockRepository) GetCpuUsage(c context.Context) ([]Metric, error) {
	return nil, errors.New("not implemented yet")
}

func (mr *MockRepository) GetRamUsage(c context.Context) (*Metric, error) {
	return nil, errors.New("not implemented yet")
}
