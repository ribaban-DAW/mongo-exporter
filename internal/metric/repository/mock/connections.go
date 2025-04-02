package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (mr *MockRepository) GetConnections(c context.Context) (*value_object.Connections, error) {
	return mr.Connections, nil
}
