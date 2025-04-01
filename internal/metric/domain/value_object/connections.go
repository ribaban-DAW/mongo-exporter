package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Connections struct {
	Current      domain.Metric `json:"current"`
	Available    domain.Metric `json:"available"`
	TotalCreated domain.Metric `json:"totalCreated"`
	Active       domain.Metric `json:"active"`
}
