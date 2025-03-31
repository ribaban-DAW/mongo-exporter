package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type MetricRepository interface {
	GetOpCounters(c context.Context) ([]domain.Metric, error)
	GetOpCounterByName(c context.Context, name string) (*domain.Metric, error)
	GetCpuUsage(c context.Context) ([]domain.Metric, error)
	GetRamUsage(c context.Context) ([]domain.Metric, error)
}
