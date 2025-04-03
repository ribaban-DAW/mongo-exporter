package service

import (
	"context"
	"errors"
	"log"

	"github.com/SrVariable/mongo-exporter/internal/metric/domain/value_object"
)

func (ms *MetricService) FindCollection(c context.Context, dbName, collName string) (*value_object.Collection, error) {
	if dbName == "" || collName == "" {
		log.Println("dbName or collName is empty")
		return nil, errors.New("dbName or collName is empty")
	}
	fullName := dbName + "." + collName
	collection, err := ms.repo.GetCollection(c, fullName)
	if err != nil {
		log.Printf("Error finding Collection %s. Reason: %v", fullName, err)
		return nil, err
	}
	log.Printf("Found Collection %s", fullName)
	return collection, nil
}
