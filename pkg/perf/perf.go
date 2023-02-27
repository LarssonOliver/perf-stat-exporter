package perf

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const fieldSeparator = ","

type PerfStats struct {
	Pid               int
	TimeoutMs           uint
	TaskClockMilliSec float64
	ContextSwitches   uint64
	CpuMigrations     uint64
	PageFaults        uint64
	CacheMisses       uint64
	CacheReferences   uint64
	Cycles            uint64
	Instructions      uint64
	Branches          uint64
	BranchMisses      uint64
}

func PerfStatProcess(pid int, timeoutMs uint) (*PerfStats, error) {
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
		fmt.Sprintf("--pid=%d", pid),
		fmt.Sprintf("--timeout=%d", timeoutMs),
		fmt.Sprintf("--field-separator=%s", fieldSeparator),
		fmt.Sprintf("--event=%s", strings.Join(statFields, ",")),
	}

	cmd := exec.Command("perf", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

    perf := parsePerfStatOutput(stderr.String())
    perf.TimeoutMs = timeoutMs
    perf.Pid = pid

    return perf, nil
}

func parsePerfStatOutput(output string) *PerfStats {
	lines := strings.Split(output, "\n")

	result := &PerfStats{}

	for _, line := range lines {
        if line == "" {
            continue
        }

		fields := strings.Split(line, fieldSeparator)

		value, err := strconv.ParseUint(fields[0], 10, 0)
		if err != nil {
			value = 0
		}

		switch fields[2] {
		case "task-clock":
			fVal, err := strconv.ParseFloat(fields[0], 64)
			if err != nil {
				fVal = 0
			}
			result.TaskClockMilliSec = fVal
		case "context-switches":
			result.ContextSwitches = value
		case "cpu-migrations":
			result.CpuMigrations = value
		case "page-faults":
			result.PageFaults = value
		case "cache-misses":
			result.CacheMisses = value
		case "cache-references":
			result.CacheReferences = value
		case "cycles":
			result.Cycles = value
		case "instructions":
			result.Instructions = value
		case "branches":
			result.Branches = value
		case "branch-misses":
			result.BranchMisses = value
		}
	}

	return result, nil
}
