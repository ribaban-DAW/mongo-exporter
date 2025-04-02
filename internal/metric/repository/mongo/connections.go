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

	conn, ok := serverStatus["connections"].(bson.M)
	if !ok {
		return nil, errors.New("`connections` type assertion failed")
	}

	connections := value_object.Connections{
		Current: domain.Metric{
			Value:     conn["current"],
			Timestamp: time.Now(),
		},
		Available: domain.Metric{
			Value:     conn["available"],
			Timestamp: time.Now(),
		},
		TotalCreated: domain.Metric{
			Value:     conn["totalCreated"],
			Timestamp: time.Now(),
		},
		Active: domain.Metric{
			Value:     conn["active"],
			Timestamp: time.Now(),
		},
	}

	return &connections, nil
}
