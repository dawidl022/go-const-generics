package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
	fggAst "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	fggParser "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	fggParsetree "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	fggReduction "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/reduction"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
)

//go:embed testdata/acceptance/program.go
var acceptanceProgramGo []byte

func TestReduceToValue_givenValidProgram_completelyReducesProgram(t *testing.T) {
	p := parseFGProgram(acceptanceProgramGo)

	val, err := newProgramReducer(nil).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, "6", val.String())
}

func TestReduceToValue_givenValidProgram_notifiesObserversOfAllReductions(t *testing.T) {
	p := parseFGProgram(acceptanceProgramGo)

	obs := &stringObserver{}
	_, err := newProgramReducer([]observer{obs}).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), obs.steps)
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
	_, err := newProgramReducer([]observer{observer1, observer2}).ReduceToValue(p)

	require.NoError(t, err)
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), observer1.stringifySteps())
	require.Equal(t, expectedAcceptanceProgramReductionSteps(), observer2.stringifySteps())
}

func TestReduceToValue_givenInvalidProgram_returnsError(t *testing.T) {
	p := parseFGProgram(fieldInvalidFieldGo)

	_, err := newProgramReducer(nil).ReduceToValue(p)

	require.Error(t, err)
	require.Equal(t, `program stuck: no field named "y" found on struct of type "Foo"`, err.Error())
}

func TestReduceToValue_givenInfiniteLoop_terminatesReductionWithError(t *testing.T) {
	p := parseFGProgram(callRecursiveGo)

	obs := &stringObserver{}
	_, err := newProgramReducer([]observer{obs}).ReduceToValue(p)

	require.Error(t, err)
	require.Equal(t, `infinite loop detected at term: "Foo{}.recurse()"`, err.Error())
	require.Equal(t, []string{"Foo{}.recurse()"}, obs.steps)
}

func newProgramReducer(observers []observer) programReducer {
	loopObservers := []loop.Observer[fgExpression]{}
	for _, obs := range observers {
		loopObservers = append(loopObservers, obs)
	}
	return programReducer{
		loop.NewProgramReducer[fgProgram, fgExpression](fgReducer{}, loopObservers),
	}
}

type observer interface {
	Notify(expression fgExpression) error
}

type programReducer struct {
	reducer *loop.ProgramReducer[fgProgram, fgExpression]
}

func (r programReducer) ReduceToValue(p ast.Program) (fgExpression, error) {
	return r.reducer.ReduceToValue(fgProgram{p})
}

type fgReducer struct {
}

func (f fgReducer) Reduce(program fgProgram) (fgProgram, error) {
	newProgram, err := program.program.Reduce()
	return fgProgram{newProgram}, err
}

type fgProgram struct {
	program ast.Program
}

func (f fgProgram) Expression() fgExpression {
	return fgExpression{f.program.Expression}
}

type fgExpression struct {
	ast.Expression
}

func (f fgExpression) IsValue() bool {
	return f.Value() != nil
}

type stringObserver struct {
	steps []string
}

func (s *stringObserver) Notify(expression fgExpression) error {
	s.steps = append(s.steps, expression.String())
	return nil
}

type savingObserver struct {
	steps []ast.Expression
}

func (s *savingObserver) Notify(expression fgExpression) error {
	s.steps = append(s.steps, expression.Expression)
	return nil
}

func (s *savingObserver) stringifySteps() interface{} {
	res := make([]string, 0, len(s.steps))
	for _, step := range s.steps {
		res = append(res, step.String())
	}
	return res
}

func parseFGProgram(code []byte) ast.Program {
	return testrunners.ParseProgram[ast.Program, *parser.FGParser](code, parsetree.ParseFGActions{})
}

func parseFGGProgram(code []byte) fggAst.Program {
	return testrunners.ParseProgram[fggAst.Program, *fggParser.FGGParser](code, fggParsetree.ParseFGGActions{})
}

func parseFGAndReduceOneStep(program []byte) (ast.Program, error) {
	return parseFGProgram(program).Reduce()
}

func parseFGGAndReduceOneStep(program []byte) (fggAst.Program, error) {
	p := parseFGGProgram(program)
	return fggReduction.NewProgramReducer().Reduce(p)
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	testrunners.AssertEqualAfterSingleReduction(t, program, expected, reductionTests())
}

func assertErrorAfterSingleReduction(t *testing.T, program []byte, expectedErrMsg string) {
	testrunners.AssertErrorAfterSingleReduction(t, program, expectedErrMsg, reductionTests())
}

func assertEqualValueAndFailsToReduce(t *testing.T, program []byte, expectedValue string) {
	testrunners.AssertEqualValueAndFailsToReduce(t, program, expectedValue, valueTests())
}
