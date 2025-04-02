package repository

import (
	vo "github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

type MockRepository struct {
	Connections *vo.Connections
	Cpu         *vo.Cpu
	OpCounters  *vo.OpCounters
	Ram         *vo.Ram
}

func NewMockRepository(connections *vo.Connections, cpu *vo.Cpu, opCounters *vo.OpCounters, ram *vo.Ram) *MockRepository {
	return &MockRepository{
		Connections: connections,
		Cpu:         cpu,
		OpCounters:  opCounters,
		Ram:         ram,
	}
}
