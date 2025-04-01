package service

import (
	"context"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (ms *MetricService) FindOpCounters(c context.Context) ([]domain.Metric, error) {
	ocs, err := ms.repo.GetOpCounters(c)
	if err != nil {
		log.Printf("Error getting opcounters. Reason: %v", err)
		return nil, err
	}
	log.Println("Found opcounters")
	return ocs, nil
}

func (ms *MetricService) FindOpCounterByName(c context.Context, name string) (*domain.Metric, error) {
	oc, err := ms.repo.GetOpCounterByName(c, name)
	if err != nil {
		log.Printf("Error getting opcounter by name. Reason: %v", err)
		return nil, err
	}
	log.Printf("Found opcounter %s", oc.Name)
	return oc, nil
}
