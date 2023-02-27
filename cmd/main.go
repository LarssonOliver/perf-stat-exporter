package main

import (
	"log"
	"net/http"

	"github.com/larssonoliver/perf-stat-exporter/pkg/exporter"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	registry := prometheus.NewRegistry()
    exporter, _ := exporter.NewPerfExporter(registry)

    log.Println("Serving metrics at :8080/metrics")
	log.Fatalln(http.ListenAndServe(":8080", exporter))
}
