package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-generic-array-sizes/monomorphisers/fgg2go/monomo"
)

//go:embed testdata/array.fgg
var arrayLibFgg string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayLenFG_givenEmptyArray_returnsZero(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Len().val()")
}

func TestArrayGetFG_givenEmptyArray_returnsZero(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Get(Zero[5, int]{})")
}

func TestArrayGetFG_givenEmptyArrayAfterPushAndPop_returnsZero(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(42).Pop().Get(Zero[5, int]{})")
}

func TestArrayPopFG_givenEmptyArray_returnsEmptyArray(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Pop()")
}

func TestArrayFG_givenElementIsPushed_appearsAtEndOfResultingArray(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, `EmptyArrayFunc{}.call().Push(21)
		.Get(EmptyArrayFunc{}.call().Push(21).Len().pred())`)
}

func TestArrayFG_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(21).Len().val()")
}

func TestArrayFG_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(21).Pop().Len().val()")
}

func TestArrayFG_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, `EmptyArrayFunc{}.call().Push(1).Push(2)
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Len().pred())`)
	assertBisimulates(t, `EmptyArrayFunc{}.call().Push(1).Push(2).Pop()
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Pop().Len().pred())`)
}

func TestArrayGetFG_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(1).Push(2).Get(Zero[5, int]{})")
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ[5, int]{Zero[5, int]{}})")
}

func TestArrayGetFG_givenAccessOutOfBounds_returnsZero(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ[5, int]{Succ[5, int]{Zero[5, int]{}}})")
}

func TestArray_givenMoreElementsPushedThanCapacityFG_doesNotChangeArray(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyArrayFunc{}.call().Push(0).Push(1).Push(2).Push(3).Push(4).Push(42)")
}

func assertBisimulates(t *testing.T, main string) {
	program := arrayLibFgg + fmt.Sprintf(mainTemplate, main)
	interpreter := entrypoint.Interpreter{}

	fggProg, err := interpreter.ParseProgram(strings.NewReader(program))
	require.NoError(t, err)

	fggMonoProg := entrypoint.FggProgram{monomo.Monomorphise(fggProg.Program)}

	for !fggProg.Expression().IsValue() {
		fggProg, err = interpreter.Reduce(fggProg)
		require.NoError(t, err)

		fggMonoProg, err = interpreter.Reduce(fggMonoProg)
		require.NoError(t, err)

		expectedMonoProg := monomo.Monomorphise(fggProg.Program)

		// monomorphise is called on monomorphised-line only for dead-type elimination!
		fggMonoProg = entrypoint.FggProgram{monomo.Monomorphise(fggMonoProg.Program)}

		require.Equal(t, expectedMonoProg, fggMonoProg.Program)
	}
}
