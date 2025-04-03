package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type OpCounters struct {
	Insert domain.Metric[int64] `json:"insert"`
	Delete domain.Metric[int64] `json:"delete"`
	Query  domain.Metric[int64] `json:"query"`
	Update domain.Metric[int64] `json:"update"`
}
