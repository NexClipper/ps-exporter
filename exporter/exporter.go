package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

var labels = []string{"label1", "label2", "label3"}
var procStats = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "ps_exporter",
	Help: "Exported process stats",
}, labels)

// Run starts the prometheus exporter
func Run(host string, port int) {
	log.Printf("Exporter running on port %s:%d\n", host, port)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}

// Export exports process metrics to the gauge
func Export(values []int) {
	// TODO: Implement
}
