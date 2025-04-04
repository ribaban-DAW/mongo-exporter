package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Collection struct {
	Insert  domain.Metric[int32] `json:"insert"`
	Remove  domain.Metric[int32] `json:"remove"`
	Queries domain.Metric[int32] `json:"query"`
	Update  domain.Metric[int32] `json:"update"`
}
