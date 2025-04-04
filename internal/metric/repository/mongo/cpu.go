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

func (dr *DatabaseRepository) GetCpu(c context.Context) (*value_object.Cpu, error) {
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	serverStatus, err := dr.getCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	extraInfo, ok := serverStatus["extra_info"].(bson.M)
	if !ok {
		return nil, errors.New("`extraInfo` type assertion failed")
	}

	keys := []string{"user_time_us", "system_time_us"}
	var values = map[string]int64{}
	for _, key := range keys {
		if k, ok := extraInfo[key].(int64); ok {
			values[key] = k
		} else {
			return nil, fmt.Errorf("`%s` type assertion failed", key)
		}
	}

	cpu := value_object.Cpu{
		UserTime: domain.Metric[int64]{
			Value:     values["user_time_us"],
			Timestamp: time.Now(),
		},
		SystemTime: domain.Metric[int64]{
			Value:     values["system_time_us"],
			Timestamp: time.Now(),
		},
	}

	return &cpu, nil
}
