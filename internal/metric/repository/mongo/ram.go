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

func (dr *DatabaseRepository) GetRam(c context.Context) (*value_object.Ram, error) {
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	serverStatus, err := dr.getCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	mem, ok := serverStatus["mem"].(bson.M)
	if !ok {
		return nil, errors.New("`mem` type assertion failed")
	}

	keys := []string{"resident", "virtual"}
	var values = map[string]int32{}
	for _, key := range keys {
		if k, ok := mem[key].(int32); ok {
			values[key] = k
		} else {
			return nil, fmt.Errorf("`%s` type assertion failed", key)
		}
	}

	ram := value_object.Ram{
		Resident: domain.Metric[int32]{
			Value:     values["resident"],
			Timestamp: time.Now(),
		},
		Virtual: domain.Metric[int32]{
			Value:     values["virtual"],
			Timestamp: time.Now(),
		},
	}

	return &ram, nil
}
