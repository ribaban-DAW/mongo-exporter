package service

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (ms *metricService) FindRamUsage(c context.Context) (*domain.Metric, error) {
	metric, err := ms.repo.GetRamUsage(c)
	if err != nil {
		return nil, err
	}
	metric.Name = "usage"
	metric.Value = metric.Value + " MiB"
	return metric, nil
}
