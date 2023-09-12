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
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), observer.steps)
}

func expectedAcceptanceProgramReductionSteps() []string {
	return []string{
		"Arr{4, 6}[Foo{3, Arr{1, 2}}.getY().first()]",
		"Arr{4, 6}[Foo{3, Arr{1, 2}}.y.first()]",
		"Arr{4, 6}[Arr{1, 2}.first()]",
		"Arr{4, 6}[Arr{1, 2}[0]]",
		"Arr{4, 6}[1]",
		"6",
	}
}

func TestReduceToValue_givenValidProgram_doesNotMutateExpressionsGivenToObservers(t *testing.T) {
	p := parseFGProgram(acceptanceProgramGo)

	observer1 := &savingObserver{}
	observer2 := &savingObserver{}
	_, err := NewProgramReducer([]Observer{observer1, observer2}).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), observer1.stringifySteps())
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), observer2.stringifySteps())
}

func TestReduceToValue_givenInvalidProgram_returnsError(t *testing.T) {
	p := parseFGProgram(fieldInvalidFieldGo)

	_, err := NewProgramReducer([]Observer{}).ReduceToValue(p)

	require.Error(t, err)
	require.Equal(t, `program stuck: no field named "y" found on struct of type "Foo"`, err.Error())
}

func TestReduceToValue_givenInfiniteLoop_terminatesReductionWithError(t *testing.T) {
	p := parseFGProgram(callRecursiveGo)

	observer := &stringObserver{}
	_, err := NewProgramReducer([]Observer{observer}).ReduceToValue(p)

	require.Error(t, err)
	require.Equal(t, `infinite loop detected at term: "Foo{}.recurse()"`, err.Error())
	require.Equal(t, []string{"Foo{}.recurse()"}, observer.steps)
}

type stringObserver struct {
	steps []string
}

func (s *stringObserver) Notify(expression ast.Expression) {
	s.steps = append(s.steps, expression.String())
}

type savingObserver struct {
	steps []ast.Expression
}

func (s *savingObserver) Notify(expression ast.Expression) {
	s.steps = append(s.steps, expression)
}

func (s *savingObserver) stringifySteps() interface{} {
	res := make([]string, 0, len(s.steps))
	for _, step := range s.steps {
		res = append(res, step.String())
	}
	return res
}

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

func parseAndReduceOneStep(program []byte) (ast.Program, error) {
	p := parseFGProgram(program)
	return p.Reduce()
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	p, err := parseAndReduceOneStep(program)

	require.NoError(t, err)
	require.Equal(t, expected, p.Expression.String())
}
