package reduction

import (
	"flag"
	"fmt"
	"slices"
	"strings"
)

var langsFlag = flag.String("langs", "fg", "comma-separated list of languages to run tests for")

type testConf struct {
	langFlags []string
}

func parseTestConf() *testConf {
	flag.Parse()
	return &testConf{langFlags: strings.Split(*langsFlag, ",")}
}

func (c testConf) enabledFG() bool {
	return slices.Contains(c.langFlags, "fg")
}

func (c testConf) enabledFGG() bool {
	return slices.Contains(c.langFlags, "fgg")
}

type reductionTestCase struct {
	name           string
	parseAndReduce func(program []byte) (fmt.Stringer, error)
	isValue        func(program []byte) bool
}

func reductionTests() []reductionTestCase {
	conf := parseTestConf()
	tests := []reductionTestCase{}

	if conf.enabledFG() {
		tests = append(tests, reductionTestCase{
			name: "FG",
			parseAndReduce: func(program []byte) (fmt.Stringer, error) {
				p, err := parseFGAndReduceOneStep(program)
				return p.Expression, err
			},
			isValue: func(program []byte) bool {
				return parseFGProgram(program).Expression.Value() != nil
			},
		})
	}
	if conf.enabledFGG() {
		tests = append(tests, reductionTestCase{
			name: "FGG",
			parseAndReduce: func(program []byte) (fmt.Stringer, error) {
				p, err := parseFGGAndReduceOneStep(program)
				return p.Expression, err
			},
			isValue: func(program []byte) bool {
				return parseFGGProgram(program).Expression.IsValue()
			},
		})
	}
	return tests
}

type valueTestCase struct {
	name           string
	parseAndReduce func(program []byte)
	parseAndValue  func(program []byte) fmt.Stringer
	isValue        func(program []byte) bool
}

func valueTests() []valueTestCase {
	conf := parseTestConf()
	tests := []valueTestCase{}

	if conf.enabledFG() {
		tests = append(tests, valueTestCase{
			name: "FG",
			parseAndReduce: func(program []byte) {
				parseFGAndReduceOneStep(program)
			},
			parseAndValue: func(program []byte) fmt.Stringer {
				return parseFGProgram(program).Expression.Value()
			},
			isValue: func(program []byte) bool {
				return parseFGProgram(program).Expression.Value() != nil
			},
		})
	}
	if conf.enabledFGG() {
		tests = append(tests, valueTestCase{
			name: "FGG",
			parseAndReduce: func(program []byte) {
				parseFGGAndReduceOneStep(program)
			},
			parseAndValue: func(program []byte) fmt.Stringer {
				return parseFGGProgram(program).Expression
			},
			isValue: func(program []byte) bool {
				return parseFGGProgram(program).Expression.IsValue()
			},
		})
	}
	return tests
}
