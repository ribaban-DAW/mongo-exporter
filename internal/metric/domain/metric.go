package domain

import "time"

type Metric struct {
	Value     any       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
