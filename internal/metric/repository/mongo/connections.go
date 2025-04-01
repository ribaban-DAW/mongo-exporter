package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetConnections(c context.Context) (*value_object.Connections, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	result, ok := serverStatus["connections"].(bson.M)
	if !ok {
		return nil, errors.New("`connections` type assertion failed")
	}

	connection := value_object.Connections{
		Current: domain.Metric{
			Value:     result["current"],
			Timestamp: time.Now(),
		},
		Available: domain.Metric{
			Value:     result["available"],
			Timestamp: time.Now(),
		},
		TotalCreated: domain.Metric{
			Value:     result["totalCreated"],
			Timestamp: time.Now(),
		},
		Active: domain.Metric{
			Value:     result["active"],
			Timestamp: time.Now(),
		},
	}

	return &connection, nil
}
