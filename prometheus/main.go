package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"go-study/prometheus/metrics"
	"net/http"
	"time"
)

func setupMetrics(ep string) {
	log.Info("prometheus metrics listening at: ", ep)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(ep, nil) // nolint: errcheck
}

func main() {
	setupMetrics(":9200")
	timer := prometheus.NewTimer(prometheus.ObserverFunc(
		metrics.ProcessDuration.WithLabelValues("feature_and_attributes", "analyzer").Observe))
	time.Sleep(time.Second * 3)
	timer.ObserveDuration()
	select {}
}
