package service

import (
	"context"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (ms *metricService) FindRamUsage(c context.Context) ([]domain.Metric, error) {
	metrics, err := ms.repo.GetRamUsage(c)
	if err != nil {
		log.Printf("Error getting RAM usage. Reason: %v", err)
		return nil, err
	}
	log.Printf("Found RAM usage")
	return metrics, nil
}
