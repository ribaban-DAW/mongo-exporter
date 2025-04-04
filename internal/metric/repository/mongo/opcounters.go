package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetOpCounters(c context.Context) (*value_object.OpCounters, error) {
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	serverStatus, err := dr.getCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	ocs, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("`opcounters` type assertion failed")
	}

	keys := []string{"insert", "query", "update", "delete"}
	var values = map[string]int64{}
	for _, key := range keys {
		if k, ok := ocs[key].(int64); ok {
			values[key] = k
		} else {
			return nil, fmt.Errorf("`%s` type assertion failed", key)
		}
	}

	opcounters := value_object.OpCounters{
		Insert: domain.Metric[int64]{
			Value:     values["insert"],
			Timestamp: time.Now(),
		},
		Query: domain.Metric[int64]{
			Value:     values["query"],
			Timestamp: time.Now(),
		},
		Update: domain.Metric[int64]{
			Value:     values["update"],
			Timestamp: time.Now(),
		},
		Delete: domain.Metric[int64]{
			Value:     values["delete"],
			Timestamp: time.Now(),
		},
	}

	return &opcounters, nil
}
