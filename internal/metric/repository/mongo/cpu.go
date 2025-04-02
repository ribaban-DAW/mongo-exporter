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
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	extraInfo, ok := serverStatus["extra_info"].(bson.M)
	if !ok {
		return nil, errors.New("`extraInfo` type assertion failed")
	}

	cpu := value_object.Cpu{
		UserTime: domain.Metric{
			Value:     extraInfo["user_time_us"],
			Timestamp: time.Now(),
		},
		SystemTime: domain.Metric{
			Value:     extraInfo["system_time_us"],
			Timestamp: time.Now(),
		},
	}

	return &cpu, nil
}
