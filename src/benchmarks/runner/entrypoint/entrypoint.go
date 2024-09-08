package entrypoint

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/dawidl022/go-const-generics/benchmarks/runner/pgfplots"
	"github.com/dawidl022/go-const-generics/benchmarks/runner/runner"
	"github.com/dawidl022/go-const-generics/benchmarks/runner/stats"
)

func Run(packagePath, benchPattern string, comparison *[2]string, output io.Writer) error {
	results, err := runner.NewRunner(packagePath, benchPattern).RunBenchmarks()
	if err != nil {
		return err
	}

	err = writeTablesToOutputDir(results, benchPattern, comparison)
	if err != nil {
		return err
	}
	return writePgfPlotsToOutputDir(results, benchPattern, comparison)
}

const outputDir = "outputs"

func writeTablesToOutputDir(results runner.Results, benchPattern string, comparison *[2]string) error {
	dirPath := path.Join(outputDir, results.PackageName)
	err := os.MkdirAll(outputPath(results, benchPattern), 0755)
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

	if comparison != nil {
		diff := stats.RelativeSpeedup(results.OfFunc(comparison[0]), results.OfFunc(comparison[1]))
		diffTable := pgfplots.NewRelativeSpeedupTable(diff)
		err = os.WriteFile(
			path.Join(outputPath(results, benchPattern), "speedup.dat"),
			[]byte(diffTable.String()), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func writePgfPlotsToOutputDir(results runner.Results, benchPattern string, comparison *[2]string) error {
	plotLatex, err := pgfplots.NewPgfPlot(results, outputPath(results, benchPattern), comparison != nil)
	if err != nil {
		return err
	}
	return os.WriteFile(
		path.Join(outputPath(results, benchPattern), fmt.Sprintf("%s.tex", results.PackageName)),
		plotLatex.Bytes(), 0644)
}

func outputPath(results runner.Results, benchPattern string) string {
	subDir := lettersOnly(benchPattern)
	if subDir == "" {
		return path.Join(outputDir, results.PackageName)
	}
	return path.Join(outputDir, results.PackageName, subDir)
}

func lettersOnly(s string) string {
	res := ""
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			res += string(char)
		}
	}
	return res
}
