package typecheck

import (
	_ "embed"
	"testing"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/fg/parser"
	"github.com/dawidl022/go-const-generics/interpreters/fg/parsetree"
	fggAst "github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	fggParser "github.com/dawidl022/go-const-generics/interpreters/fgg/parser"
	fggParsetree "github.com/dawidl022/go-const-generics/interpreters/fgg/parsetree"
	fggTypecheck "github.com/dawidl022/go-const-generics/interpreters/fgg/typecheck"
	"github.com/dawidl022/go-const-generics/interpreters/shared/testrunners"
)

//go:embed testdata/acceptance/program.go
var acceptanceProgramGo []byte

func TestTypeCheck_givenWellTypedProgram_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, acceptanceProgramGo)
}

func parseFGProgram(code []byte) ast.Program {
	return testrunners.ParseProgram[ast.Program, *parser.FGParser](code, parsetree.ParseFGActions{})
}

func parseFGGProgram(code []byte) fggAst.Program {
	return testrunners.ParseProgram[fggAst.Program, *fggParser.FGGParser](code, fggParsetree.ParseFGGActions{})
}

func parseAndTypeCheckFG(program []byte) error {
	p := parseFGProgram(program)
	return TypeCheck(p)
}

func parseAndTypeCheckFGG(program []byte) error {
	p := parseFGGProgram(program)
	return fggTypecheck.TypeCheck(p)
}

func assertFailsTypeCheckWithError(t *testing.T, program []byte, errMsg string) {
	testrunners.AssertFailsTypeCheckWithError(t, program, errMsg, testCases())
}

func assertPassesTypeCheck(t *testing.T, program []byte) {
	testrunners.AssertPassesTypeCheck(t, program, testCases())
}
