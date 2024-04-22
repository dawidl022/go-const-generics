package typecheck

import (
	"testing"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/preprocessor"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
)

func parseFGGProgram(code []byte) (ast.Program, error) {
	p := testrunners.ParseProgram[ast.Program, *parser.FGGParser](code, parsetree.ParseFGGActions{})
	return preprocessor.IdentifyTypeParams(p)
}

func parseAndTypeCheckFGG(program []byte) error {
	p, err := parseFGGProgram(program)
	if err != nil {
		return err
	}
	return TypeCheck(p)
}

func assertFailsTypeCheckWithError(t *testing.T, program []byte, errMsg string) {
	testrunners.AssertFailsTypeCheckWithError(t, program, errMsg, testCases())
}

func assertPassesTypeCheck(t *testing.T, program []byte) {
	testrunners.AssertPassesTypeCheck(t, program, testCases())
}

func testCases() []testrunners.TypeTestCase {
	return []testrunners.TypeTestCase{fggTestCase{}}
}

type fggTestCase struct {
}

func (f fggTestCase) Name() string {
	return "FGG"
}

func (f fggTestCase) ParseAndTypeCheck(program []byte) error {
	return parseAndTypeCheckFGG(program)
}
