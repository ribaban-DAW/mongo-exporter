package service

import (
	"context"
	"log"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (ms *MetricService) FindCpu(c context.Context) (*value_object.Cpu, error) {
	cpu, err := ms.repo.GetCpu(c)
	if err != nil {
		log.Printf("Error finding CPU. Reason: %v", err)
		return nil, err
	}

	cpu.TotalTime = domain.Metric[int64]{
		Value:     cpu.UserTime.Value + cpu.SystemTime.Value,
		Timestamp: time.Now(),
	}
	log.Println("Found CPU")
	return cpu, nil
}
