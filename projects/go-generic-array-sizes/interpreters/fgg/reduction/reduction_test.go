package reduction

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
)

func parseFGGProgram(code []byte) ast.Program {
	return testrunners.ParseProgram[ast.Program](code, parseFGGActions{})
}

type parseFGGActions struct {
}

func (parseFGGActions) NewLexer(input antlr.CharStream) antlr.Lexer {
	return parser.NewFGGLexer(input)
}

func (parseFGGActions) NewParser(input antlr.TokenStream) antlr.Parser {
	return parser.NewFGGParser(input)
}

func (parseFGGActions) Program(p antlr.Parser) antlr.ParseTree {
	return p.(*parser.FGGParser).Program()
}

func (parseFGGActions) NewAstBuilder(tree antlr.ParseTree) testrunners.ASTBuilder[ast.Program] {
	return parsetree.NewAntlrASTBuilder(tree)
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
