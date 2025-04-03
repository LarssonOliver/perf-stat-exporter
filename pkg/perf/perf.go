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

package perf

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

const fieldSeparator = ";"

func (pf *PerfCollector) StartPerfStatProcessBlocking(pid int, parseIntervalMs uint) error {
	statFields := []string{
		"task-clock",
		"context-switches",
		"cpu-migrations",
		"page-faults",
		"cache-misses",
		"cache-references",
		"cycles",
		"instructions",
		"branches",
		"branch-misses",
	}

	args := []string{
		"stat",
		fmt.Sprintf("--field-separator=%s", fieldSeparator),
		fmt.Sprintf("--interval-print=%d", parseIntervalMs),
		fmt.Sprintf("--event=%s", strings.Join(statFields, ",")),
	}

	if pid < 0 {
		args = append(args, "--all-cpus")
	} else {
		args = append(args, fmt.Sprintf("--pid=%d", pid))
	}

	cmd := exec.Command("perf", args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)

	for scanner.Scan() {
		line := scanner.Text()
		pf.parsePerfStatOutput(pid, line)
	}

	return nil
}

func (pf *PerfCollector) parsePerfStatOutput(pid int, perfOutputLine string) {
	perfOutputLine = strings.TrimSpace(perfOutputLine)

	if perfOutputLine == "" {
		return
	}

	fields := strings.Split(perfOutputLine, fieldSeparator)

	value, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		value = 0
	}

	labels := prometheus.Labels{"pid": strconv.FormatInt(int64(pid), 10)}

	if pid < 0 {
		labels["pid"] = "all-cpus"
	}

	pf.Lock()
	defer pf.Unlock()

	switch fields[3] {
	case "task-clock":
		pf.taskClockCounter.With(labels).Add(value)
	case "context-switches":
		pf.contextSwitchesCounter.With(labels).Add(value)
	case "cpu-migrations":
		pf.cpuMigrationsCounter.With(labels).Add(value)
	case "page-faults":
		pf.pageFaultCounter.With(labels).Add(value)
	case "cache-misses":
		pf.cacheMissesCounter.With(labels).Add(value)
	case "cache-references":
		pf.cacheReferencesCounter.With(labels).Add(value)
	case "cycles":
		pf.cyclesCounter.With(labels).Add(value)
	case "instructions":
		pf.instructionsCounter.With(labels).Add(value)
	case "branches":
		pf.branchesCounter.With(labels).Add(value)
	case "branch-misses":
		pf.branchMissesCounter.With(labels).Add(value)
	}
}
