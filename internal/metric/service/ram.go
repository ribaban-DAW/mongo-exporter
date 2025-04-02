package service

import (
	"context"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (ms *MetricService) FindRam(c context.Context) (*value_object.Ram, error) {
	ram, err := ms.repo.GetRam(c)
	if err != nil {
		log.Printf("Error finding RAM. Reason: %v", err)
		return nil, err
	}
	log.Println("Found RAM")
	return ram, nil
}
