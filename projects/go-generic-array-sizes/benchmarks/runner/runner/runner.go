package runner

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

const MaxSize = 1024 * 1024 * 32

type Results struct {
	PackageName string
	Goos        string
	Goarch      string
	Results     []BenchmarkResults
}

func (r Results) OfFunc(funcName string) BenchmarkResults {
	for _, res := range r.Results {
		if res.FuncName == funcName {
			return res
		}
	}
	panic("invalid funcName provided")
}

type BenchmarkResults struct {
	FuncName string
	Metrics  map[int]BenchmarkMetrics
}

type BenchmarkMetrics struct {
	NsPerOp     float64
	AllocsPerOp int
}

//go:embed templates/config.go.tmpl
var configTmpl string

const benchReExpr = `goos: (.*)\ngoarch: (.*)\npkg: .*\n(([A-Za-z0-9]+\W+[0-9]+\W+[0-9\.]+ ns\/op\W+[0-9]+ B\/op\W+[0-9]+ allocs\/op\n)+)PASS`
const resultReExpr = `([A-Za-z0-9]+)\W+([0-9]+)\W+([0-9\.]+) ns\/op\W+([0-9]+) B\/op\W+([0-9]+) allocs\/op`

type Runner struct {
	benchRe      *regexp.Regexp
	resultRe     *regexp.Regexp
	packagePath  string
	benchPattern string
	tmpl         *template.Template
}

func NewRunner(packagePath, benchPattern string) *Runner {
	return &Runner{
		benchRe:      regexp.MustCompile(benchReExpr),
		resultRe:     regexp.MustCompile(resultReExpr),
		packagePath:  packagePath,
		benchPattern: benchPattern,
		tmpl:         newConfigTemplate(),
	}
}

func (r *Runner) configFilePath() string {
	return path.Join(r.packagePath, "config.go")
}

func (r *Runner) RunBenchmarks() (Results, error) {
	// buffer existing config file and defer writing it back
	originalConfFile, err := os.ReadFile(r.configFilePath())
	if err != nil {
		return Results{}, fmt.Errorf("cannot open config.go in package %q: %w", r.packagePath, err)
	}
	defer os.WriteFile(r.configFilePath(), originalConfFile, 0644)

	return r.performBenchmarks()
}

func (r *Runner) performBenchmarks() (Results, error) {
	results, err := r.runEmptyBenchmark()
	if err != nil {
		return Results{}, err
	}

	for size := 1; size <= MaxSize; size <<= 1 {
		err = r.runNextBenchmark(&results, size)
		if err != nil {
			return Results{}, err
		}
	}
	return results, nil
}

func (r *Runner) runEmptyBenchmark() (Results, error) {
	err := r.writeConfigFile(0)
	if err != nil {
		return Results{}, err
	}
	rawOutput, err := r.executeGoTestBench()
	if err != nil {
		return Results{}, err
	}

	return r.newResults(rawOutput)
}

func (r *Runner) runNextBenchmark(res *Results, size int) error {
	err := r.writeConfigFile(size)
	if err != nil {
		return err
	}

	rawOutput, err := r.executeGoTestBench()
	if err != nil {
		return err
	}
	return r.appendResults(rawOutput, res, size)
}

func (r *Runner) writeConfigFile(size int) error {
	conf := newConfig(r.packagePath, size)

	confFile := &bytes.Buffer{}
	err := r.tmpl.Execute(confFile, conf)
	if err != nil {
		return err
	}

	return os.WriteFile(r.configFilePath(), confFile.Bytes(), 0644)
}

func (r *Runner) executeGoTestBench() (*bytes.Buffer, error) {
	cmd := exec.Command("go", "test", fmt.Sprintf("-bench=%s", r.benchPattern))
	cmd.Env = append(cmd.Env, "GOMAXPROCS=1")
	cmd.Env = append(cmd.Env, fmt.Sprintf("HOME=%s", os.Getenv("HOME")))
	cmd.Dir = r.packagePath

	output := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (r *Runner) newResults(output *bytes.Buffer) (Results, error) {
	groups, err := r.matchOutput(output)
	if err != nil {
		return Results{}, err
	}
	res, err := r.parseNewResults(groups[3])
	return Results{
		PackageName: path.Base(r.packagePath),
		Goos:        string(groups[1]),
		Goarch:      string(groups[2]),
		Results:     res,
	}, err
}

func (r *Runner) matchOutput(output *bytes.Buffer) ([][]byte, error) {
	groups := r.benchRe.FindSubmatch(output.Bytes())
	if len(groups) < 4 {
		return nil, fmt.Errorf("unparsable benchmark output:\n%s", output.String())
	}
	return groups, nil
}

func (r *Runner) parseNewResults(rawResults []byte) ([]BenchmarkResults, error) {
	var res []BenchmarkResults
	linedResults := r.parseLines(rawResults)

	for _, line := range linedResults {
		nsPerOp, allocsPerOp, err := r.matchResults(line)
		if err != nil {
			return nil, err
		}

		res = append(res, BenchmarkResults{
			FuncName: r.resultRe.FindStringSubmatch(line)[1],
			Metrics: map[int]BenchmarkMetrics{
				0: {
					NsPerOp:     nsPerOp,
					AllocsPerOp: allocsPerOp,
				},
			},
		})
	}
	return res, nil
}

func (r *Runner) matchResults(line string) (nsPerOp float64, allocsPerOp int, err error) {
	resElements := r.resultRe.FindStringSubmatch(line)
	if len(resElements) < 6 {
		err = fmt.Errorf("unparsable benchmark result line:\n%s", line)
		return
	}

	nsPerOp, err = strconv.ParseFloat(resElements[3], 64)
	if err != nil {
		err = fmt.Errorf("failed to parse benchmark results in ns/op field: %w", err)
		return
	}
	allocsPerOp, err = strconv.Atoi(resElements[5])
	if err != nil {
		err = fmt.Errorf("failed to parse benchmark results in allocs/op field: %w", err)
		return
	}
	return
}

func (r *Runner) parseLines(rawResults []byte) []string {
	return strings.Split(strings.TrimSpace(string(rawResults)), "\n")
}

func (r *Runner) appendResults(output *bytes.Buffer, results *Results, size int) error {
	groups, err := r.matchOutput(output)
	if err != nil {
		return err
	}
	return r.parseResults(groups[3], results, size)
}

func (r *Runner) parseResults(rawResults []byte, res *Results, size int) error {
	linedResults := r.parseLines(rawResults)

	for i, line := range linedResults {
		nsPerOp, allocsPerOp, err := r.matchResults(line)
		if err != nil {
			return err
		}

		res.Results[i].Metrics[size] = BenchmarkMetrics{
			NsPerOp:     nsPerOp,
			AllocsPerOp: allocsPerOp,
		}
	}
	return nil
}

type config struct {
	PackageName string
	Size        int
}

func newConfig(packagePath string, size int) config {
	return config{
		PackageName: path.Base(packagePath),
		Size:        size,
	}
}

func newConfigTemplate() *template.Template {
	tmpl, err := template.New("configTmpl").Parse(configTmpl)
	if err != nil {
		panic(err)
	}
	return tmpl
}
