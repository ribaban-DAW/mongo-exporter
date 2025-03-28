package service

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

// NOTE: This will be simplified when I implement Value Objects
func calculateCpuUsage(metricsStart, metricsEnd []domain.Metric, elapsedSecs float64) (float64, error) {
	// Type casting
	userTimeStart, err := strconv.ParseInt(metricsStart[0].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert userTimeStart")
	}
	userTimeEnd, err := strconv.ParseInt(metricsEnd[0].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert userTimeEnd")
	}

	systemTimeStart, err := strconv.ParseInt(metricsStart[1].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert systemTimeStart")
	}
	systemTimeEnd, err := strconv.ParseInt(metricsEnd[1].Value, 10, 64)
	if err != nil {
		return 0, errors.New("couldn't convert systemTimeEnd")
	}

	// Calculation
	deltaTotalTime := (userTimeEnd - userTimeStart) + (systemTimeEnd - systemTimeStart)
	elapsedMicrosecs := elapsedSecs * 1000000.0
	cpuUsage := (float64(deltaTotalTime) / elapsedMicrosecs) / float64(runtime.NumCPU()) * 100.0

	return cpuUsage, nil
}

func (ms *metricService) FindCpuUsage(c context.Context) (*domain.Metric, error) {
	elapsedSecs := 1.0

	metricsStart, err := ms.repo.GetCpuUsage(c)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(elapsedSecs) * time.Second)

	metricsEnd, err := ms.repo.GetCpuUsage(c)
	if err != nil {
		return nil, err
	}

	cpuUsage, err := calculateCpuUsage(metricsStart, metricsEnd, elapsedSecs)
	if err != nil {
		return nil, err
	}
	metric := domain.Metric{
		Name:      "usage",
		Value:     fmt.Sprintf("%.2f%%", cpuUsage),
		Timestamp: time.Now(),
	}
	return &metric, nil
}
