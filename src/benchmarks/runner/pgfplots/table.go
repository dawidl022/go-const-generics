package pgfplots

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/dawidl022/go-const-generics/benchmarks/runner/runner"
)

type Table []string

func (t Table) String() string {
	return strings.Join(t, "\n")
}

func NewTable(results runner.BenchmarkResults) Table {
	table := Table{"size\tnsPerOp\tallocsPerOp"}

	sortedKeys := maps.Keys(results.Metrics)
	slices.Sort(sortedKeys)

	for _, key := range sortedKeys {
		table = append(table, fmt.Sprintf("%d\t%f\t%d", key, results.Metrics[key].NsPerOp, results.Metrics[key].AllocsPerOp))
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
