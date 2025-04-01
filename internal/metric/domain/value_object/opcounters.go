package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type OpCounters struct {
	Insert domain.Metric `json:"insert"`
	Delete domain.Metric `json:"delete"`
	Query  domain.Metric `json:"query"`
	Update domain.Metric `json:"update"`
}
