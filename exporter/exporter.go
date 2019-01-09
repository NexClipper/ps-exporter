package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

// Run starts the prometheus exporter
func Run(host string, port int) {
	procCollector := newProcCollector()
	prometheus.MustRegister(procCollector)
	log.Printf("Exporter running on port %s:%d\n", host, port)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}
