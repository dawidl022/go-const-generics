package reduction

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
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

func reductionTests() []testrunners.ReductionTestCase {
	tests := []testrunners.ReductionTestCase{}
	for _, test := range valueTests() {
		tests = append(tests, test)
	}
	return tests
}

func valueTests() []testrunners.ValueTestCase {
	conf := parseTestConf()
	tests := []testrunners.ValueTestCase{}

	if conf.enabledFG() {
		tests = append(tests, fgTestCase{})
	}
	if conf.enabledFGG() {
		tests = append(tests, fggTestCase{})
	}
	return tests
}

type fgTestCase struct {
}

func (f fgTestCase) Name() string {
	return "FG"
}

func (f fgTestCase) ParseAndReduce(program []byte) (fmt.Stringer, error) {
	p, err := parseFGAndReduceOneStep(program)
	return p.Expression, err
}

func (f fgTestCase) IsValue(program []byte) bool {
	return parseFGProgram(program).Expression.Value() != nil
}

func (f fgTestCase) ParseAndValue(program []byte) fmt.Stringer {
	return parseFGProgram(program).Expression.Value()
}

type fggTestCase struct {
}

func (f fggTestCase) Name() string {
	return "FGG"
}

func (f fggTestCase) ParseAndReduce(program []byte) (fmt.Stringer, error) {
	p, err := parseFGGAndReduceOneStep(program)
	return p.Expression, err
}

func (f fggTestCase) IsValue(program []byte) bool {
	return parseFGGProgram(program).Expression.IsValue()
}

func (f fggTestCase) ParseAndValue(program []byte) fmt.Stringer {
	return parseFGGProgram(program).Expression
}
