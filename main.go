package main

import (
	"flag"

	"github.com/nexclipper/ps-exporter/exporter"
)

func main() {
	port := flag.Int("port", 9095, "Port used for the exporter service")
	host := flag.String("host", "localhost", "Host used for the exporter service")
	flag.Parse()
	exporter.Run(*host, *port)
}
