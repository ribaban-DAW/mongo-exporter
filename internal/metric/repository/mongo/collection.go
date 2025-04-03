package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) GetCollection(c context.Context, collName string) (*value_object.Collection, error) {
	cmd := bson.D{
		{Key: "top", Value: 1},
	}
	top, err := dr.getCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	totals, ok := top["totals"].(bson.M)
	if !ok {
		return nil, errors.New("`mem` type assertion failed")
	}

	coll, ok := totals[collName].(bson.M)
	if !ok {
		return nil, errors.New("`collection` type assertion failed")
	}

	values := []string{"insert", "queries", "update", "remove"}
	var result = map[string]int32{}
	for _, value := range values {
		if k, ok := coll[value].(bson.M); ok {
			if c, ok := k["count"].(int32); ok {
				result[value] = c
			} else {
				return nil, fmt.Errorf("%s `count` type assertion failed", k)
			}
		}
	}

	collection := value_object.Collection{
		Insert: domain.Metric[int32]{
			Value:     result["insert"],
			Timestamp: time.Now(),
		},
		Remove: domain.Metric[int32]{
			Value:     result["remove"],
			Timestamp: time.Now(),
		},
		Update: domain.Metric[int32]{
			Value:     result["update"],
			Timestamp: time.Now(),
		},
		Queries: domain.Metric[int32]{
			Value:     result["queries"],
			Timestamp: time.Now(),
		},
	}
	return &collection, nil
}
