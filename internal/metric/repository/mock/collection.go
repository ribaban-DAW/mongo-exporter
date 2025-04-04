package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (mr *MockRepository) GetCollection(c context.Context, collName string) (*value_object.Collection, error) {
	return mr.Collection, nil
}
