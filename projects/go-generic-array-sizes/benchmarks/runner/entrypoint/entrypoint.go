package entrypoint

import (
	"encoding/json"
	"io"

	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/runner"
)

func Run(packagePath string, output io.Writer) error {
	results, err := runner.NewRunner(packagePath).RunBenchmarks()
	if err != nil {
		return err
	}

	marshalled, err := json.Marshal(results)
	if err != nil {
		return err
	}
	_, err = output.Write(marshalled)
	return err
}
