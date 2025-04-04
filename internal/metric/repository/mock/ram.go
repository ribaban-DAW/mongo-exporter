package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (mr *MockRepository) GetRam(c context.Context) (*value_object.Ram, error) {
	return mr.Ram, nil
}
