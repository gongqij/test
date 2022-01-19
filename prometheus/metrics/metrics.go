package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ProcessDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "process_durations",
		Help:       "The time duration of objects processing",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"type", "step"})
)

func init() {
	prometheus.MustRegister(ProcessDuration)
}
