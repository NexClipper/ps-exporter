# ps-exporter

Process status exporter for [Prometheus](https://prometheus.io/).

## Metrics

Exports the following metrics for all processes:

* `memp`: Memory utilization percentage (`ps -eo %mem`)
* `cpup`: CPU utilization percentage (`ps -eo %cpu`)
* `cpu`: Short-term CPU usage for scheduling (`ps -eo cpu`)

All metrics are annotated with the following labels:

* `pid`
* `ppid`
* `command`
* `user`
* `started`

## Usage

Start the exporter using: `./ps-exporter --port 9095 --host localhost`
