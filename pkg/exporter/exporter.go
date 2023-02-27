package exporter

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/larssonoliver/perf-stat-exporter/pkg/perf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	sync.Mutex

	totalScrapes prometheus.Counter

	mux *http.ServeMux

	registry *prometheus.Registry
}

func NewPerfExporter(registry *prometheus.Registry) (*Exporter, error) {
	if registry == nil {
		return nil, errors.New("Parameter 'registry' required.")
	}

	exporter := &Exporter{
		registry: registry,
	}

	exporter.totalScrapes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "perf",
		Name:      "exporter_scrapes_total",
		Help:      "Current total metric scrapes.",
	})

	exporter.registry.MustRegister(exporter)

	exporter.mux = http.NewServeMux()
	exporter.mux.Handle("/metrics", promhttp.HandlerFor(
		exporter.registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError},
	))

	return exporter, nil
}

func (exporter *Exporter) Describe(descChan chan<- *prometheus.Desc) {
	descChan <- exporter.totalScrapes.Desc()
}

func (exporter *Exporter) Collect(metricsChan chan<- prometheus.Metric) {
	exporter.Lock()
	defer exporter.Unlock()

	exporter.totalScrapes.Inc()
	metricsChan <- exporter.totalScrapes

	stats, err := perf.PerfStatProcess(os.Getpid(), 100)
	if err != nil {
        fmt.Printf("Error occurred while collecting perf stat metrics: %v", err)
	}

    fmt.Println(stats)
}

func (exporter *Exporter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	exporter.mux.ServeHTTP(writer, request)
}
