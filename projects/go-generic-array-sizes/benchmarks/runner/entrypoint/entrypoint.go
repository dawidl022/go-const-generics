package entrypoint

import (
	"fmt"
	"io"
	"os"
	"path"

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
	return nil
}
