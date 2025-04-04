package value_object

import (
	"github.com/SrVariable/mongo-exporter/internal/metric/domain"
)

type Ram struct {
	Resident domain.Metric[int32] `json:"resident"`
	Virtual  domain.Metric[int32] `json:"virtual"`
}
