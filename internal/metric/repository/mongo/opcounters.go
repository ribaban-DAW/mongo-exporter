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

	insert, ok := ocs["insert"].(int64)
	if !ok {
		return nil, errors.New("`insert` type assertion failed")
	}
	query, ok := ocs["query"].(int64)
	if !ok {
		return nil, errors.New("`query` type assertion failed")
	}
	update, ok := ocs["update"].(int64)
	if !ok {
		return nil, errors.New("`update` type assertion failed")
	}
	delete, ok := ocs["delete"].(int64)
	if !ok {
		return nil, errors.New("`delete` type assertion failed")
	}

	opcounters := value_object.OpCounters{
		Insert: domain.Metric[int64]{
			Value:     insert,
			Timestamp: time.Now(),
		},
		Query: domain.Metric[int64]{
			Value:     query,
			Timestamp: time.Now(),
		},
		Update: domain.Metric[int64]{
			Value:     update,
			Timestamp: time.Now(),
		},
		Delete: domain.Metric[int64]{
			Value:     delete,
			Timestamp: time.Now(),
		},
	}

	return &opcounters, nil
}
