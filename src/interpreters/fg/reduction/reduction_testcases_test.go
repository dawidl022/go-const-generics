package reduction

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fg/testconf"
	"github.com/dawidl022/go-const-generics/interpreters/shared/testrunners"
)

func reductionTests() []testrunners.ReductionTestCase {
	tests := []testrunners.ReductionTestCase{}
	for _, test := range valueTests() {
		tests = append(tests, test)
	}
	return tests
}

func valueTests() []testrunners.ValueTestCase {
	conf := testconf.ParseTestConf()
	tests := []testrunners.ValueTestCase{}

	if conf.EnabledFG() {
		tests = append(tests, fgTestCase{})
	}
	if conf.EnabledFGG() {
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
