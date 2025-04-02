package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (mr *MockRepository) GetOpCounters(c context.Context) (*value_object.OpCounters, error) {
	return mr.OpCounters, nil
}
