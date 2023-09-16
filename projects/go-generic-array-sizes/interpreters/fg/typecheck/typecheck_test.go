package typecheck

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
)

//go:embed testdata/acceptance/program.go
var acceptanceProgramGo []byte

// TODO
//func TestTypeCheck_givenWellTypedProgram_returnsNoError(t *testing.T) {
//	err := parseAndTypeCheck(acceptanceProgramGo)
//	require.NoError(t, err)
//}

// TODO eliminate duplication of reduction_test.go
func parseFGProgram(code []byte) ast.Program {
	input := antlr.NewIoStream(bytes.NewBuffer(code))
	lexer := parser.NewFGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGParser(stream)
	p.AddErrorListener(failingErrorListener{})
	p.BuildParseTrees = true

	tree := p.Program()
	astBuilder := parsetree.NewAntlrASTBuilder(tree)
	return astBuilder.BuildAST()
}

type failingErrorListener struct {
	*antlr.DefaultErrorListener
}

func (f failingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
	panic(msg)
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
