package mock

import (
	"context"
	"errors"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type MockRepository struct {
	Metrics []domain.Metric
}

func NewMockRepository(metrics []domain.Metric) *MockRepository {
	return &MockRepository{
		Metrics: metrics,
	}
}

func (mr *MockRepository) GetOpCounters(c context.Context) ([]domain.Metric, error) {
	return mr.Metrics, nil
}

func (mr *MockRepository) GetOpCounterByName(c context.Context, name string) (*domain.Metric, error) {
	for _, metric := range mr.Metrics {
		if metric.Name == name {
			return &metric, nil
		}
	}
	return nil, errors.New("metric not found")
}

func (mr *MockRepository) GetCpuUsage(c context.Context) ([]domain.Metric, error) {
	return nil, errors.New("not implemented yet")
}

func (mr *MockRepository) GetRamUsage(c context.Context) ([]domain.Metric, error) {
	return nil, errors.New("not implemented yet")
}
