package runner

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"text/template"
)

type Results struct {
	packageName string
	results     []BenchmarkResults
}

type BenchmarkResults struct {
	funcName string
	metrics  map[int]BenchmarkMetrics
}

type BenchmarkMetrics struct {
	nsPerOp     int
	allocsPerOp int
}

//go:embed templates/config.go.tmpl
var configTmpl string

func RunBenchmarks(packagePath string) (Results, error) {
	conf := newConfig(packagePath)
	tmpl, err := template.New("configTmpl").Parse(configTmpl)
	if err != nil {
		return Results{}, err
	}

	// buffer existing config file and defer writing it back
	configFilePath := path.Join(packagePath, "config.go")
	originalConfFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Results{}, err
	}
	defer os.WriteFile(configFilePath, originalConfFile, 0644)

	// TODO do first iteration with zero size

	for i := 1; i <= 1024*1024*32; i <<= 1 {
		conf.Size = i

		confFile := &bytes.Buffer{}
		err = tmpl.Execute(confFile, conf)
		if err != nil {
			return Results{}, err
		}

		// TODO can const be used instead of numeric literal?
		err = os.WriteFile(configFilePath, confFile.Bytes(), 0644)
		if err != nil {
			return Results{}, err
		}

		cmd := exec.Command("go", "test", "-bench=.")
		cmd.Env = append(cmd.Env, "GOMAXPROCS=1")
		cmd.Env = append(cmd.Env, fmt.Sprintf("HOME=%s", os.Getenv("HOME")))
		cmd.Dir = packagePath
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return Results{}, err
		}
	}

	return Results{}, nil
}

type config struct {
	PackageName string
	Size        int
}

func newConfig(packagePath string) config {
	return config{
		PackageName: path.Base(packagePath),
		Size:        0,
	}
}
