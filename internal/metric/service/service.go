package service

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/repository"
)

type MetricService struct {
	repo repository.MetricRepository
}

func NewMetricService(repo repository.MetricRepository) *MetricService {
	return &MetricService{repo}
}
