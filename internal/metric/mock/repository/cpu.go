package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (mr *MockRepository) GetCpu(c context.Context) (*value_object.Cpu, error) {
	return mr.Cpu, nil
}
