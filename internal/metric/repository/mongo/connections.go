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

	keys := []string{"current", "available", "totalCreated", "active"}
	var values = map[string]int32{}
	for _, key := range keys {
		if k, ok := conn[key].(int32); ok {
			values[key] = k
		} else {
			return nil, fmt.Errorf("`%s` type assertion failed", key)
		}
	}

	connections := value_object.Connections{
		Current: domain.Metric[int32]{
			Value:     values["current"],
			Timestamp: time.Now(),
		},
		Available: domain.Metric[int32]{
			Value:     values["available"],
			Timestamp: time.Now(),
		},
		TotalCreated: domain.Metric[int32]{
			Value:     values["totalCreated"],
			Timestamp: time.Now(),
		},
		Active: domain.Metric[int32]{
			Value:     values["active"],
			Timestamp: time.Now(),
		},
	}

	return &connections, nil
}
