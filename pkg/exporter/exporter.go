package exporter

import (
	"errors"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	sync.Mutex

	totalScrapes prometheus.Counter

	mux *http.ServeMux

	registry *prometheus.Registry
}

// ServeHTTP implements http.Handler
func (exporter *Exporter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	exporter.mux.ServeHTTP(writer, request)
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
}
