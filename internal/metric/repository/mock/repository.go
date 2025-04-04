package repository

import (
	vo "github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

// NOTE: Maybe I should create a mock for each value object instead of a single mock for all of them.
type MockRepository struct {
	Collection  *vo.Collection
	Connections *vo.Connections
	Cpu         *vo.Cpu
	OpCounters  *vo.OpCounters
	Ram         *vo.Ram
}

func NewMockRepository(collection *vo.Collection, connections *vo.Connections, cpu *vo.Cpu, opCounters *vo.OpCounters, ram *vo.Ram) *MockRepository {
	return &MockRepository{
		Collection:  collection,
		Connections: connections,
		Cpu:         cpu,
		OpCounters:  opCounters,
		Ram:         ram,
	}
}
