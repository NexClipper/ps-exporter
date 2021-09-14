package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// ProcCollector collects process informtion
type ProcCollector struct {
	memPMetric *prometheus.Desc
	cpuPMetric *prometheus.Desc
	cpuMetric  *prometheus.Desc
}

var labels = []string{
	"pid", "ppid", "command", "user",
}

func newProcCollector() *ProcCollector {
	return &ProcCollector{
		memPMetric: prometheus.NewDesc("ps_mem_util_percent", "Memory utilization percentage", labels, nil),
		cpuPMetric: prometheus.NewDesc("ps_cpu_util_percent", "CPU utilization percentage", labels, nil),
		cpuMetric:  prometheus.NewDesc("ps_cpu_short_usage", "Short-term CPU usage for scheduling", labels, nil),
	}
}

// Descibe adds the descriptor
func (collector *ProcCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.cpuMetric
	ch <- collector.cpuPMetric
	ch <- collector.memPMetric
}

// Collect collects prometheus metrics
func (collector *ProcCollector) Collect(ch chan<- prometheus.Metric) {
	procStats := GetProcessStats()
	for _, proc := range procStats {
		ch <- prometheus.MustNewConstMetric(
			collector.cpuMetric,
			prometheus.GaugeValue,
			proc.cpu,
			proc.pid,
			proc.ppid,
			proc.command,
			proc.user,
			// proc.start,
		)
		ch <- prometheus.MustNewConstMetric(
			collector.cpuPMetric,
			prometheus.GaugeValue,
			proc.cpuP,
			proc.pid,
			proc.ppid,
			proc.command,
			proc.user,
			// proc.start,
		)
		ch <- prometheus.MustNewConstMetric(
			collector.memPMetric,
			prometheus.GaugeValue,
			proc.memP,
			proc.pid,
			proc.ppid,
			proc.command,
			proc.user,
			// proc.start,
		)
	}
}
