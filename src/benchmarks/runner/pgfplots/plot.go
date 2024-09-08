package pgfplots

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/dawidl022/go-const-generics/benchmarks/runner/runner"
)

//go:embed templates/plot.tex.tmpl
var plotTmpl []byte

func NewPgfPlot(results runner.Results, outputPath string, comparison bool) (*bytes.Buffer, error) {
	tmpl, err := template.New("pfgPlot").Parse(string(plotTmpl))
	if err != nil {
		panic(err)
	}

	latexFile := &bytes.Buffer{}
	err = tmpl.Execute(latexFile, newConfig(results, outputPath, comparison))
	if err != nil {
		return nil, err
	}

	return latexFile, nil
}

type config struct {
	FuncNames   []string
	PackageName string
	OutputPath  string
	Comparison  bool
	Goos        string
	Goarch      string
}

func newConfig(results runner.Results, outputPath string, comparison bool) config {
	var funcNames []string
	for _, res := range results.Results {
		funcNames = append(funcNames, res.FuncName)
	}

	return config{
		FuncNames:   funcNames,
		PackageName: results.PackageName,
		OutputPath:  outputPath,
		Comparison:  comparison,
		Goos:        results.Goos,
		Goarch:      results.Goarch,
	}
}
