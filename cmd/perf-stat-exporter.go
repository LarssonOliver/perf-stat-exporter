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
	pid             = flag.Int("pid", -1, "PID of process to observe, -1 observes all processes.")
	parseIntervalMs = flag.Uint("interval", 10_000, "perf-stat scraping interval in ms.")
)

func main() {
	flag.Parse()

	pids := []int{*pid}
    if *pid != -1 {
        pids = append(pids, -1)
    }

	registry := prometheus.NewRegistry()
	exporter, _ := exporter.NewPerfExporter(registry, uint(*parseIntervalMs), pids...)

	fmt.Println("Serving metrics at :8080/metrics")
	fmt.Println(http.ListenAndServe(":8080", exporter))
	os.Exit(-1)
}
