package entrypoint

import (
	"fmt"
	"io"
	"os"
	"path"
	"slices"

	"golang.org/x/exp/maps"

	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/pgfplots"
	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/runner"
)

func Run(packagePath string, output io.Writer) error {
	results, err := runner.NewRunner(packagePath).RunBenchmarks()
	if err != nil {
		return err
	}

	return writeTablesToOutputDir(results)
}

const outputDir = "outputs"

func writeTablesToOutputDir(results runner.Results) error {
	dirPath := path.Join(outputDir, results.PackageName)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		table := pgfplots.NewTable(res)
		err = os.WriteFile(
			path.Join(dirPath, fmt.Sprintf("%s.dat", res.FuncName)),
			[]byte(table.String()), 0644)
		if err != nil {
			return err
		}
	}

	diff := relativeSpeedup(results.Results[0], results.Results[1])
	diffTable := pgfplots.NewRelativeSpeedupTable(diff)
	err = os.WriteFile(
		path.Join(dirPath, "speedup.dat"),
		[]byte(diffTable.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func relativeSpeedup(b1 runner.BenchmarkResults, b2 runner.BenchmarkResults) map[int]float64 {
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
