/*
Copyright 2025 Oliver Larsson

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

	registry := prometheus.NewRegistry()
	exporter, _ := exporter.NewPerfExporter(registry, uint(*parseIntervalMs), *pid)

	fmt.Println("Serving metrics at :8080/metrics")
	fmt.Println(http.ListenAndServe(":8080", exporter))
	os.Exit(-1)
}
