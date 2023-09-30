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
	fggAst "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	fggParser "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	fggParsetree "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	fggReduction "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/reduction"
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

func (s *stringObserver) Notify(expression ast.Expression) error {
	s.steps = append(s.steps, expression.String())
	return nil
}

type savingObserver struct {
	steps []ast.Expression
}

func (s *savingObserver) Notify(expression ast.Expression) error {
	s.steps = append(s.steps, expression)
	return nil
}

func (s *savingObserver) stringifySteps() interface{} {
	res := make([]string, 0, len(s.steps))
	for _, step := range s.steps {
		res = append(res, step.String())
	}
	return res
}

type parseActionable[T any] interface {
	newLexer(input antlr.CharStream) antlr.Lexer
	newParser(input antlr.TokenStream) antlr.Parser
	program(parser antlr.Parser) antlr.ParseTree
	newAstBuilder(tree antlr.ParseTree) parsetree.ASTBuilder[T]
}

func parseProgram[T any](code []byte, actions parseActionable[T]) T {
	input := antlr.NewIoStream(bytes.NewBuffer(code))
	lexer := actions.newLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := actions.newParser(stream)
	p.AddErrorListener(failingErrorListener{})

	tree := actions.program(p)
	astBuilder := actions.newAstBuilder(tree)
	return astBuilder.BuildAST()
}

func parseFGProgram(code []byte) ast.Program {
	return parseProgram[ast.Program](code, parseFGActions{})
}

type parseFGActions struct{}

func (parseFGActions) newLexer(input antlr.CharStream) antlr.Lexer {
	return parser.NewFGLexer(input)
}

func (parseFGActions) newParser(input antlr.TokenStream) antlr.Parser {
	return parser.NewFGParser(input)
}

func (parseFGActions) program(p antlr.Parser) antlr.ParseTree {
	return p.(*parser.FGParser).Program()
}

func (parseFGActions) newAstBuilder(tree antlr.ParseTree) parsetree.ASTBuilder[ast.Program] {
	return parsetree.NewAntlrASTBuilder(tree)
}

func parseFGGProgram(code []byte) fggAst.Program {
	return parseProgram[fggAst.Program](code, parseFGGActions{})
}

type parseFGGActions struct {
}

func (parseFGGActions) newLexer(input antlr.CharStream) antlr.Lexer {
	return fggParser.NewFGGLexer(input)
}

func (parseFGGActions) newParser(input antlr.TokenStream) antlr.Parser {
	return fggParser.NewFGGParser(input)
}

func (parseFGGActions) program(p antlr.Parser) antlr.ParseTree {
	return p.(*fggParser.FGGParser).Program()
}

func (parseFGGActions) newAstBuilder(tree antlr.ParseTree) parsetree.ASTBuilder[fggAst.Program] {
	return fggParsetree.NewAntlrASTBuilder(tree)
}

type failingErrorListener struct {
	*antlr.DefaultErrorListener
}

func (f failingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
	panic(msg)
}

func parseFGAndReduceOneStep(program []byte) (ast.Program, error) {
	p := parseFGProgram(program)
	return p.Reduce()
}

func parseFGGAndReduceOneStep(program []byte) (fggAst.Program, error) {
	p := parseFGGProgram(program)
	return fggReduction.NewProgramReducer().Reduce(p)
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	for _, test := range reductionTests() {
		t.Run(test.name, func(t *testing.T) {
			expr, err := test.parseAndReduce(program)

			require.NoError(t, err)
			require.Equal(t, expected, expr.String())
			require.False(t, test.isValue(program))
		})
	}
}

func assertErrorAfterSingleReduction(t *testing.T, program []byte, expectedErrMsg string) {
	for _, test := range reductionTests() {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.parseAndReduce(program)

			require.EqualError(t, err, expectedErrMsg)
			require.False(t, test.isValue(program))
		})
	}
}

func assertEqualValueAndFailsToReduce(t *testing.T, program []byte, expectedValue string) {
	for _, test := range valueTests() {
		t.Run(test.name, func(t *testing.T) {
			require.Panics(t, func() { test.parseAndReduce(program) })

			require.True(t, test.isValue(program))
			require.Equal(t, expectedValue, test.parseAndValue(program).String())
		})
	}
}
