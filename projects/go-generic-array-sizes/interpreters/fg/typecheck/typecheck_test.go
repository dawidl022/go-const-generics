package typecheck

import (
	_ "embed"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
)

//go:embed testdata/acceptance/program.go
var acceptanceProgramGo []byte

func TestTypeCheck_givenWellTypedProgram_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(acceptanceProgramGo)
	require.NoError(t, err)
}

func parseFGProgram(code []byte) ast.Program {
	return testrunners.ParseProgram[ast.Program, *parser.FGParser](code, parsetree.ParseFGActions{})
}

func parseAndTypeCheck(program []byte) error {
	p := parseFGProgram(program)
	return TypeCheck(p)
}

func assertFailsTypeCheckWithError(t *testing.T, program []byte, errMsg string) {
	err := parseAndTypeCheck(program)

	require.Error(t, err)
	require.Equal(t, errMsg, err.Error())
}
