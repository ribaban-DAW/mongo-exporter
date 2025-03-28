package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (dr *DatabaseRepository) GetOpCounters(c context.Context) ([]domain.Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	var metrics []domain.Metric

	opcounters, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	for metricName, metricValue := range opcounters {
		metric := domain.Metric{
			Name:      metricName,
			Value:     fmt.Sprintf("%d", metricValue),
			Timestamp: time.Now(),
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (dr *DatabaseRepository) GetOpCounterByName(c context.Context, name string) (*domain.Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	opcounters, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}

	value, ok := opcounters[name]
	if !ok {
		return nil, errors.New("metric not found")
	}

	metric := domain.Metric{
		Name:      name,
		Value:     fmt.Sprintf("%d", value),
		Timestamp: time.Now(),
	}
	return &metric, nil
}
