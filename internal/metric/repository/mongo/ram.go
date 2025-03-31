package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetRamUsage(c context.Context) ([]domain.Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	var metrics []domain.Metric

	mem, ok := serverStatus["mem"].(bson.M)
	if !ok {
		return nil, errors.New("`mem` type assertion failed")
	}
	resident := mem["resident"]
	metrics = append(
		metrics,
		domain.Metric{
			Name:      "resident",
			Value:     fmt.Sprintf("%d", resident),
			Timestamp: time.Now(),
		},
	)
	virtual := mem["virtual"]
	metrics = append(
		metrics,
		domain.Metric{
			Name:      "virtual",
			Value:     fmt.Sprintf("%d", virtual),
			Timestamp: time.Now(),
		},
	)
	return metrics, nil
}
