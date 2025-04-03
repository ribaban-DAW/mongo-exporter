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
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	serverStatus, err := dr.getCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	conn, ok := serverStatus["connections"].(bson.M)
	if !ok {
		return nil, errors.New("`connections` type assertion failed")
	}

	current, ok := conn["current"].(int32)
	if !ok {
		return nil, errors.New("`current` type assertion failed")
	}
	available, ok := conn["available"].(int32)
	if !ok {
		return nil, errors.New("`available` type assertion failed")
	}
	totalCreated, ok := conn["totalCreated"].(int32)
	if !ok {
		return nil, errors.New("`totalCreated` type assertion failed")
	}
	active, ok := conn["active"].(int32)
	if !ok {
		return nil, errors.New("`active` type assertion failed")
	}

	connections := value_object.Connections{
		Current: domain.Metric[int32]{
			Value:     current,
			Timestamp: time.Now(),
		},
		Available: domain.Metric[int32]{
			Value:     available,
			Timestamp: time.Now(),
		},
		TotalCreated: domain.Metric[int32]{
			Value:     totalCreated,
			Timestamp: time.Now(),
		},
		Active: domain.Metric[int32]{
			Value:     active,
			Timestamp: time.Now(),
		},
	}

	return &connections, nil
}
