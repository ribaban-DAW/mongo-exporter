package mongo

import (
	"context"
	"errors"
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

	userTime, ok := extraInfo["user_time_us"].(int64)
	if !ok {
		return nil, errors.New("`user_time_us` type assertion failed")
	}
	systemTime, ok := extraInfo["system_time_us"].(int64)
	if !ok {
		return nil, errors.New("`system_time_us` type assertion failed")
	}

	cpu := value_object.Cpu{
		UserTime: domain.Metric[int64]{
			Value:     userTime,
			Timestamp: time.Now(),
		},
		SystemTime: domain.Metric[int64]{
			Value:     systemTime,
			Timestamp: time.Now(),
		},
	}

	return &cpu, nil
}
