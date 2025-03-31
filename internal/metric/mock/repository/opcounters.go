package repository

import (
	"context"
	"errors"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

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
