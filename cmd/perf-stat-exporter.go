package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/larssonoliver/perf-stat-exporter/pkg/exporter"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	pid             = flag.Int("pid", 0, "PID of process to observe.")
	parseIntervalMs = flag.Uint("interval", 10_000, "perf-stat scraping interval in ms.")
)

func main() {
	flag.Parse()
	if *pid < 1 {
		fmt.Println("A valid -pid is required, exiting...")
		os.Exit(1)
	}

	registry := prometheus.NewRegistry()
	exporter, _ := exporter.NewPerfExporter(registry, uint(*parseIntervalMs), *pid)

	fmt.Println("Serving metrics at :8080/metrics")
	fmt.Println(http.ListenAndServe(":8080", exporter))
	os.Exit(-1)
}
