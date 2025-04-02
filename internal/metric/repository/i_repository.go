package repository

import (
	"context"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

type MetricRepository interface {
	GetOpCounters(c context.Context) (*value_object.OpCounters, error)
	GetCpu(c context.Context) (*value_object.Cpu, error)
	GetRam(c context.Context) (*value_object.Ram, error)
	GetConnections(c context.Context) (*value_object.Connections, error)
}
