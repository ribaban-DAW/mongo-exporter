package service

import (
	"context"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (ms *MetricService) FindOpCounters(c context.Context) (*value_object.OpCounters, error) {
	opcounters, err := ms.repo.GetOpCounters(c)
	if err != nil {
		log.Printf("Error finding OpCounters. Reason: %v", err)
		return nil, err
	}
	log.Println("Found OpCounters")
	return opcounters, nil
}
