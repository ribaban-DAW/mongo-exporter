package service

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/repository"
)

type MetricService interface {
	FindOpCounters(c context.Context) ([]domain.Metric, error)
	FindOpCounterByName(c context.Context, name string) (*domain.Metric, error)
	FindCpuUsage(c context.Context) (*domain.Metric, error)
	FindRamUsage(c context.Context) (*domain.Metric, error)
}

type metricService struct {
	repo repository.MetricRepository
}

func NewMetricService(repo repository.MetricRepository) *metricService {
	return &metricService{repo}
}
