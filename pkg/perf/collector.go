package perf

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type PerfCollector struct {
	sync.RWMutex

	taskClockCounter       *prometheus.CounterVec
	contextSwitchesCounter *prometheus.CounterVec
	cpuMigrationsCounter   *prometheus.CounterVec
	pageFaultCounter       *prometheus.CounterVec
	cacheMissesCounter     *prometheus.CounterVec
	cacheReferencesCounter *prometheus.CounterVec
	cyclesCounter          *prometheus.CounterVec
	instructionsCounter    *prometheus.CounterVec
	branchesCounter        *prometheus.CounterVec
	branchMissesCounter    *prometheus.CounterVec
}

func NewPerfCollector() *PerfCollector {
	labels := []string{"pid"}
	namespace := "perf"

	pf := &PerfCollector{}

	pf.taskClockCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "task_clock",
		Help:      "Cpu task clock time.",
	}, labels)

	pf.contextSwitchesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "context_switches",
		Help:      "Number of observed context-switches.",
	}, labels)

	pf.cpuMigrationsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "cpu_migrations",
		Help:      "Number of observed cpu-migrations.",
	}, labels)

	pf.pageFaultCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "page_faults",
		Help:      "Number of observed page-faults.",
	}, labels)

	pf.cacheMissesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "cache_misses",
		Help:      "Number of observed cache-misses.",
	}, labels)

	pf.cacheReferencesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "cache_references",
		Help:      "Number of observed cache-references.",
	}, labels)

	pf.cyclesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "cycles",
		Help:      "Number of observed cycles.",
	}, labels)

	pf.instructionsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "instructions",
		Help:      "Number of observed instructions.",
	}, labels)

	pf.branchesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "branches",
		Help:      "Number of observed branches.",
	}, labels)

	pf.branchMissesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "branch_misses",
		Help:      "Number of observed branch misses.",
	}, labels)

	return pf
}

func (pf *PerfCollector) Describe(descChan chan<- *prometheus.Desc) {
	pf.RLock()
	defer pf.RUnlock()

	pf.taskClockCounter.Describe(descChan)
	pf.contextSwitchesCounter.Describe(descChan)
	pf.cpuMigrationsCounter.Describe(descChan)
	pf.pageFaultCounter.Describe(descChan)
	pf.cacheMissesCounter.Describe(descChan)
	pf.cacheReferencesCounter.Describe(descChan)
	pf.cyclesCounter.Describe(descChan)
	pf.instructionsCounter.Describe(descChan)
	pf.branchesCounter.Describe(descChan)
	pf.branchMissesCounter.Describe(descChan)
}

func (pf *PerfCollector) Collect(metricChan chan<- prometheus.Metric) {
	pf.RLock()
	defer pf.RUnlock()

	pf.taskClockCounter.Collect(metricChan)
	pf.contextSwitchesCounter.Collect(metricChan)
	pf.cpuMigrationsCounter.Collect(metricChan)
	pf.pageFaultCounter.Collect(metricChan)
	pf.cacheMissesCounter.Collect(metricChan)
	pf.cacheReferencesCounter.Collect(metricChan)
	pf.cyclesCounter.Collect(metricChan)
	pf.instructionsCounter.Collect(metricChan)
	pf.branchesCounter.Collect(metricChan)
	pf.branchMissesCounter.Collect(metricChan)
}
