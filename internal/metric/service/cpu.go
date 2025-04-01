package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func calculateTotalTime(cpu *value_object.Cpu) (int64, error) {
	userTime, ok := cpu.UserTime.Value.(int64)
	if !ok {
		return 0, errors.New("`userTime` type assertion failed")
	}

	systemTime, ok := cpu.SystemTime.Value.(int64)
	if !ok {
		return 0, errors.New("`systemType` type assertion failed")
	}

	return userTime + systemTime, nil
}

func (ms *MetricService) FindCpu(c context.Context) (*value_object.Cpu, error) {
	cpu, err := ms.repo.GetCpu(c)
	if err != nil {
		log.Printf("Error finding CPU. Reason: %v", err)
		return nil, err
	}

	totalTime, err := calculateTotalTime(cpu)
	if err != nil {
		log.Printf("Error calculating totalTime. Reason: %v", err)
		return nil, err
	}

	// NOTE: This isn't calculated in repository because it's not part of MongoDB stats
	cpu.TotalTime = domain.Metric{
		Value:     totalTime,
		Timestamp: time.Now(),
	}
	log.Println("Found CPU")
	return cpu, nil
}
