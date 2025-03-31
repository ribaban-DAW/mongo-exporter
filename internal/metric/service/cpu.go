package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

// NOTE: This will be simplified when I implement Value Objects
func CalculateCpuUsage(metricsStart []domain.Metric) (int64, error) {
	// Type casting
	userTime, err := strconv.ParseInt(metricsStart[0].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert userTimeStart to int")
	}
	systemTime, err := strconv.ParseInt(metricsStart[1].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert systemTimeStart to int")
	}

	totalTime := userTime + systemTime
	return totalTime, nil
}

func (ms *metricService) FindCpuUsage(c context.Context) ([]domain.Metric, error) {
	metrics, err := ms.repo.GetCpuUsage(c)
	if err != nil {
		log.Printf("Error getting CPU usage. Reason: %v", err)
		return nil, err
	}

	totalTime, err := CalculateCpuUsage(metrics)
	if err != nil {
		log.Printf("Error calculating CPU usage. Reason: %v", err)
		return nil, err
	}
	metric := domain.Metric{
		Name:      "totalTime",
		Value:     fmt.Sprintf("%d", totalTime),
		Timestamp: time.Now(),
	}
	metrics = append(metrics, metric)
	log.Println("Found CPU usage")
	return metrics, nil
}
