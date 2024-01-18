package pgfplots

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/runner"
)

type Table []string

func (t Table) String() string {
	return strings.Join(t, "\n")
}

func NewTable(results runner.BenchmarkResults) Table {
	table := Table{"size\tnsPerOp"}

	sortedKeys := maps.Keys(results.Metrics)
	slices.Sort(sortedKeys)

	for _, key := range sortedKeys {
		table = append(table, fmt.Sprintf("%d\t%f", key, results.Metrics[key].NsPerOp))
	}
	return table
}

func NewRelativeSpeedupTable(multipliers map[int]float64) Table {
	table := Table{"size\trelativeSpeedup"}

	sortedKeys := maps.Keys(multipliers)
	slices.Sort(sortedKeys)

	for _, key := range sortedKeys {
		table = append(table, fmt.Sprintf("%d\t%f", key, multipliers[key]))
	}
	return table
}
