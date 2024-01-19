package stats

import (
	"slices"

	"golang.org/x/exp/maps"

	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/runner"
)

func RelativeSpeedup(b1 runner.BenchmarkResults, b2 runner.BenchmarkResults) map[int]float64 {
	sortedKeys := maps.Keys(b1.Metrics)
	slices.Sort(sortedKeys)

	res := make(map[int]float64)

	for _, key := range sortedKeys {
		if b1.Metrics[key].NsPerOp <= b2.Metrics[key].NsPerOp {
			res[key] = (1/float64(b1.Metrics[key].NsPerOp))/(1/float64(b2.Metrics[key].NsPerOp)) - 1.0
		} else {
			res[key] = -((1/float64(b2.Metrics[key].NsPerOp))/(1/float64(b1.Metrics[key].NsPerOp)) - 1.0)
		}
	}
	return res
}
