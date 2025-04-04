package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Cpu struct {
	UserTime   domain.Metric[int64] `json:"userTime"`
	SystemTime domain.Metric[int64] `json:"systemTime"`
	TotalTime  domain.Metric[int64] `json:"totalTime"`
}
