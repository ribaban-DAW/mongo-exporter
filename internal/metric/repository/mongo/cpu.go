package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetCpuUsage(c context.Context) ([]domain.Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	var metrics []domain.Metric

	extraInfo, ok := serverStatus["extra_info"].(bson.M)
	if !ok {
		return nil, errors.New("`extraInfo` type assertion failed")
	}
	userTime := extraInfo["user_time_us"]
	metrics = append(
		metrics,
		domain.Metric{
			Name:      "userTime",
			Value:     fmt.Sprintf("%d", userTime),
			Timestamp: time.Now(),
		},
	)
	systemTime := extraInfo["system_time_us"]
	metrics = append(
		metrics,
		domain.Metric{
			Name:      "systemTime",
			Value:     fmt.Sprintf("%d", systemTime),
			Timestamp: time.Now(),
		},
	)
	return metrics, nil
}
