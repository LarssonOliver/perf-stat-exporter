# perf-stat-exporter

A tool that queries Linux processes using `perf stat` and exports the metrics 
in a Prometheus readable format.

## Usage

```bash
./bin/perf-stat-exporter -pid <process pid>
```

This requires `perf` to be available in your `PATH`.

Metrics endpoint is then available at `:8080/metrics`.

## Building

Check [go.mod](./go.mod) for the required go version.

```bash
make
```

This produces `./bin/perf-stat-exporter`.

