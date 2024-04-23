package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
)

//go:embed testdata/array.fgg
var arrayLibGo string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayLenFG_givenEmptyArray_returnsZero(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Len().val()")
}

func TestArrayGetFG_givenEmptyArray_returnsZero(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Get(Zero[5, int]{})")
}

func TestArrayGetFG_givenEmptyArrayAfterPushAndPop_returnsZero(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Push(42).Pop().Get(Zero[5, int]{})")
}

func TestArrayPopFG_givenEmptyArray_returnsEmptyArray(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Array[5, int]{Arr[5, int]{0, 0, 0, 0, 0}, Zero[5, int]{}, "+
			"Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Zero[5, int]{}}}}}}, 0}",
		"EmptyArrayFunc{}.call().Pop()")
}

func TestArrayFG_givenElementIsPushed_appearsAtEndOfResultingArray(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "21",
		`EmptyArrayFunc{}.call().Push(21)
		.Get(EmptyArrayFunc{}.call().Push(21).Len().pred())`)
}

func TestArrayFG_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyArrayFunc{}.call().Push(21).Len().val()")
}

func TestArrayFG_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Push(21).Pop().Len().val()")
}

func TestArrayFG_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "2",
		`EmptyArrayFunc{}.call().Push(1).Push(2)
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Len().pred())`)
	assertReducesTo(t, "1",
		`EmptyArrayFunc{}.call().Push(1).Push(2).Pop()
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Pop().Len().pred())`)
}

func TestArrayGetFG_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Zero[5, int]{})")
	assertReducesTo(t, "2",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ[5, int]{Zero[5, int]{}})")
}

func TestArrayGetFG_givenAccessOutOfBounds_returnsZero(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ[5, int]{Succ[5, int]{Zero[5, int]{}}})")
}

func TestArray_givenMoreElementsPushedThanCapacityFG_doesNotChangeArray(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Array[5, int]{Arr[5, int]{0, 1, 2, 3, 4}, "+
			"Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Zero[5, int]{}}}}}}, "+
			"Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Zero[5, int]{}}}}}}, 0}",
		"EmptyArrayFunc{}.call().Push(0).Push(1).Push(2).Push(3).Push(4).Push(42)")
}

// TODO make sure this works with both FGG interpreter and monomo
func assertReducesTo(t *testing.T, expected string, main string) {
	program := arrayLibGo + fmt.Sprintf(mainTemplate, main)
	res, err := entrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
