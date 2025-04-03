package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Connections struct {
	Current      domain.Metric[int32] `json:"current"`
	Available    domain.Metric[int32] `json:"available"`
	TotalCreated domain.Metric[int32] `json:"totalCreated"`
	Active       domain.Metric[int32] `json:"active"`
}
