package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetOpCounters(c context.Context) (*value_object.OpCounters, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	ocs, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("`opcounters` type assertion failed")
	}

	opcounters := value_object.OpCounters{
		Insert: domain.Metric{
			Value:     ocs["insert"],
			Timestamp: time.Now(),
		},
		Query: domain.Metric{
			Value:     ocs["query"],
			Timestamp: time.Now(),
		},
		Update: domain.Metric{
			Value:     ocs["update"],
			Timestamp: time.Now(),
		},
		Delete: domain.Metric{
			Value:     ocs["delete"],
			Timestamp: time.Now(),
		},
	}

	return &opcounters, nil
}
