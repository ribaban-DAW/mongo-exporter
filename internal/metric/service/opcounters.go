package service

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (ms *metricService) FindOpCounters(c context.Context) ([]domain.Metric, error) {
	return ms.repo.GetOpCounters(c)
}

func (ms *metricService) FindOpCounterByName(c context.Context, name string) (*domain.Metric, error) {
	return ms.repo.GetOpCounterByName(c, name)
}
