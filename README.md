# perf-stat-exporter

A tool that queries linux processes using `perf stat` and exports the metrics 
in a prometheus format.

## Usage

```bash
./perf-stat-exporter -pid <process pid>
```

Metrics endpoint is then available at `:8080/metrics`.

