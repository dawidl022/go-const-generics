package reduction

import (
	"fmt"
	"testing"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/parser"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/preprocessor"
	"github.com/dawidl022/go-const-generics/interpreters/shared/testrunners"
)

func parseFGGProgram(code []byte) ast.Program {
	p := testrunners.ParseProgram[ast.Program, *parser.FGGParser](code, parsetree.ParseFGGActions{})
	p, err := preprocessor.IdentifyTypeParams(p)
	if err != nil {
		panic(err)
	}
	return p
}

func parseFGGAndReduceOneStep(program []byte) (ast.Program, error) {
	p := parseFGGProgram(program)
	return NewProgramReducer().Reduce(p)
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	testrunners.AssertEqualAfterSingleReduction(t, program, expected, []testrunners.ReductionTestCase{fggTestCase{}})
}

func assertErrorAfterSingleReduction(t *testing.T, program []byte, expectedErrMsg string) {
	testrunners.AssertErrorAfterSingleReduction(t, program, expectedErrMsg, []testrunners.ReductionTestCase{fggTestCase{}})
}

func assertEqualValueAndFailsToReduce(t *testing.T, program []byte, expectedValue string) {
	testrunners.AssertEqualValueAndFailsToReduce(t, program, expectedValue, []testrunners.ValueTestCase{fggTestCase{}})
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
