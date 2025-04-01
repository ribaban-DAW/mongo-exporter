package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Cpu struct {
	UserTime   domain.Metric `json:"userTime"`
	SystemTime domain.Metric `json:"systemTime"`
	TotalTime  domain.Metric `json:"totalTime"`
}
