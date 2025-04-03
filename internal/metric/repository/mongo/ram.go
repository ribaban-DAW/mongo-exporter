package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetRam(c context.Context) (*value_object.Ram, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	mem, ok := serverStatus["mem"].(bson.M)
	if !ok {
		return nil, errors.New("`mem` type assertion failed")
	}

	resident, ok := mem["resident"].(int32)
	if !ok {
		return nil, errors.New("`resident` type assertion failed")
	}
	virtual, ok := mem["virtual"].(int32)
	if !ok {
		return nil, errors.New("`virtual` type assertion failed")
	}

	ram := value_object.Ram{
		Resident: domain.Metric[int32]{
			Value:     resident,
			Timestamp: time.Now(),
		},
		Virtual: domain.Metric[int32]{
			Value:     virtual,
			Timestamp: time.Now(),
		},
	}

	return &ram, nil
}
