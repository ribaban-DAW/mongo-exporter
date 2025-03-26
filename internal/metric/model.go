package metric

import "time"

type Metric struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
