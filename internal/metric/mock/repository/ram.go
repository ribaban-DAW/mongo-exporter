package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (mr *MockRepository) GetRamUsage(c context.Context) ([]domain.Metric, error) {
	return mr.Metrics, nil
}
