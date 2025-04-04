package domain

import "time"

type Metric[T any] struct {
	Value     T         `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
