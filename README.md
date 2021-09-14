# ps-exporter

Process status exporter for [Prometheus](https://prometheus.io/).

## Metrics

Exports the following metrics for all processes:

* `ps_mem_util_percent`: (Gauge) Memory utilization percentage (`ps -eo pmem`)
* `ps_cpu_util_percent`: (Gauge) CPU utilization percentage (`ps -eo pcpu`)
* `ps_cpu_short_usage`: (Gauge) Short-term CPU usage for scheduling (`ps -eo cpu`)

All metrics are annotated with the following labels:

* `pid`
* `ppid`
* `command`
* `user`
* `started`

## Usage

Start the exporter using: `./ps-exporter --port 9095 --host localhost`

### Docker
```
docker run -d \
  -p 9095:9095 \
  --pid="host" \
  -v "/proc:/host/proc:ro,rslave" \
  nexclipper/ps-exporter:latest
```

### Docker-Compose

```
version: '3.8'

services:
  ps-exporter:
    image: nexclipper/ps-exporter:latest
    container_name: ps-exporter
    restart: unless-stopped
    pid: host
    volumes:
      - type: bind
        source: /proc
        target: /host/proc
        read_only: true
          bind:
            propagation: rslave
    ports: 
      - 9095:9095
```