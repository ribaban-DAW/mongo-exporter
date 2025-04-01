package service

import (
	"context"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (ms *MetricService) FindConnections(c context.Context) (*value_object.Connections, error) {
	connection, err := ms.repo.GetConnections(c)
	if err != nil {
		log.Printf("Error finding Connections metrics. Reason: %v", err)
		return nil, err
	}
	log.Printf("Found Connections")
	return connection, nil
}
