package prometheus

import (
	"context"
	"log"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	userTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_user_time_us",
		Help: "Mongo Exporter User time in microseconds in the server",
	})

	systemTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_system_time_us",
		Help: "Mongo Exporter System time in microseconds in the server",
	})

	totalTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_total_time_us",
		Help: "Mongo Exporter Total time in microseconds in the server",
	})
)

// TODO: Calculate the usage of CPU based on the totalTime
func RecordCpu(ms *service.MetricService) {
	go func() {
		for {
			cpu, err := ms.FindCpu(context.Background())
			if err != nil {
				log.Printf("Error retrieving CPU metrics: %v", err)
			} else {
				userTime.Set(float64(cpu.UserTime.Value))
				systemTime.Set(float64(cpu.SystemTime.Value))
				totalTime.Set(float64(cpu.TotalTime.Value))
			}
			time.Sleep(2 * time.Second)
		}
	}()
}
