package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

func (dr *DatabaseRepository) GetRamUsage(c context.Context) (*domain.Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	mem, ok := serverStatus["mem"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	resident := mem["resident"]

	metric := domain.Metric{
		Name:      "resident",
		Value:     fmt.Sprintf("%d", resident),
		Timestamp: time.Now(),
	}
	return &metric, nil
}
