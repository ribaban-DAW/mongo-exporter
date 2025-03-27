package metric

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MetricRepository interface {
	GetOpCounters(c context.Context) ([]Metric, error)
	GetOpCounterByName(c context.Context, name string) (*Metric, error)
	GetCpuUsage(c context.Context) ([]Metric, error)
	GetRamUsage(c context.Context) (*Metric, error)
}

type DatabaseRepository struct {
	Client *mongo.Client
}

func NewDatabaseRepository(client *mongo.Client) *DatabaseRepository {
	return &DatabaseRepository{
		Client: client,
	}
}

func (dr *DatabaseRepository) getServerStatus(c context.Context) (bson.M, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	return result, err
}

func (dr *DatabaseRepository) GetOpCounters(c context.Context) ([]Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	var metrics []Metric

	opcounters, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	for metricName, metricValue := range opcounters {
		metric := Metric{
			Name:      metricName,
			Value:     fmt.Sprintf("%d", metricValue),
			Timestamp: time.Now(),
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (dr *DatabaseRepository) GetOpCounterByName(c context.Context, name string) (*Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	opcounters, ok := serverStatus["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}

	value, ok := opcounters[name]
	if !ok {
		return nil, errors.New("metric not found")
	}

	metric := Metric{
		Name:      name,
		Value:     fmt.Sprintf("%d", value),
		Timestamp: time.Now(),
	}
	return &metric, nil
}

func (dr *DatabaseRepository) GetCpuUsage(c context.Context) ([]Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	var metrics []Metric

	extraInfo, ok := serverStatus["extra_info"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	userTime := extraInfo["user_time_us"]
	metrics = append(
		metrics,
		Metric{
			Name:      "userTime",
			Value:     fmt.Sprintf("%d", userTime),
			Timestamp: time.Now(),
		},
	)
	systemTime := extraInfo["system_time_us"]
	metrics = append(
		metrics,
		Metric{
			Name:      "systemTime",
			Value:     fmt.Sprintf("%d", systemTime),
			Timestamp: time.Now(),
		},
	)
	return metrics, nil
}

func (dr *DatabaseRepository) GetRamUsage(c context.Context) (*Metric, error) {
	serverStatus, err := dr.getServerStatus(c)
	if err != nil {
		return nil, err
	}

	mem, ok := serverStatus["mem"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	resident := mem["resident"]

	metric := Metric{
		Name:      "resident",
		Value:     fmt.Sprintf("%d", resident),
		Timestamp: time.Now(),
	}
	return &metric, nil
}
