package repository

import (
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
