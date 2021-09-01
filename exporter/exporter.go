package exporter

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run starts the prometheus exporter
func Run(host string, port int) {
	procCollector := newProcCollector()
	prometheus.MustRegister(procCollector)
	log.Printf("Exporter running on port %s:%d\n", host, port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil); err != nil {
		panic(err)
	}
}
