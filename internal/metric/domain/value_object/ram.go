package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Ram struct {
	Resident domain.Metric `json:"resident"`
	Virtual  domain.Metric `json:"virtual"`
}
