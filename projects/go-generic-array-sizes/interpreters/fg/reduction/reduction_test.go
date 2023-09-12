package reduction

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

func TestReduceToValue_givenValidProgram_completelyReducesProgram(t *testing.T) {
	p := parseFGProgram(acceptanceProgramGo)

	val, err := NewProgramReducer([]Observer{}).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, "6", val.String())
}

func TestReduceToValue_givenValidProgram_notifiesObserversOfAllReductions(t *testing.T) {
	p := parseFGProgram(acceptanceProgramGo)

	observer := &stringObserver{}
	_, err := NewProgramReducer([]Observer{observer}).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, []string{
		"Arr{4, 6}[Foo{3, Arr{1, 2}}.getY().first()]",
		"Arr{4, 6}[Foo{3, Arr{1, 2}}.y.first()]",
		"Arr{4, 6}[Arr{1, 2}.first()]",
		"Arr{4, 6}[Arr{1, 2}[0]]",
		"Arr{4, 6}[1]",
		"6",
	}, observer.steps)
}

type stringObserver struct {
	steps []string
}

func (s *stringObserver) Notify(expression ast.Expression) {
	s.steps = append(s.steps, expression.String())
}

func parseFGProgram(code []byte) ast.Program {
	// TODO handle parse errors
	input := antlr.NewIoStream(bytes.NewBuffer(code))
	lexer := parser.NewFGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGParser(stream)
	p.BuildParseTrees = true

	tree := p.Program()
	astBuilder := parsetree.NewAntlrASTBuilder(tree)
	return astBuilder.BuildAST()
}

func parseAndReduceOneStep(program []byte) (ast.Program, error) {
	p := parseFGProgram(program)
	return p.Reduce()
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	p, err := parseAndReduceOneStep(program)

	require.NoError(t, err)
	require.Equal(t, expected, p.Expression.String())
}
