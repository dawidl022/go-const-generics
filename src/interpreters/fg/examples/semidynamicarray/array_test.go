package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-const-generics/interpreters/fg/entrypoint"
	fggEntrypoint "github.com/dawidl022/go-const-generics/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-const-generics/interpreters/shared/loop"
)

//go:embed array.go
var arrayLibGo string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayLen_givenEmptyArray_returnsZero(t *testing.T) {
	assert.Equal(t, 0, EmptyArrayFunc{}.call().Len().val())
}

func TestArrayLenFG_givenEmptyArray_returnsZero(t *testing.T) {
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Len().val()")
}

func TestArrayGet_givenEmptyArray_returnsZero(t *testing.T) {
	assert.Equal(t, 0, EmptyArrayFunc{}.call().Get(Zero{}))
}

func TestArrayGetFG_givenEmptyArray_returnsZero(t *testing.T) {
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Get(Zero{})")
}

func TestArrayGet_givenEmptyArrayAfterPushAndPop_returnsZero(t *testing.T) {
	assert.Equal(t, 0, EmptyArrayFunc{}.call().Push(42).Pop().Get(Zero{}))
}

func TestArrayGetFG_givenEmptyArrayAfterPushAndPop_returnsZero(t *testing.T) {
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Push(42).Pop().Get(Zero{})")
}

func TestArrayPop_givenEmptyArray_returnsEmptyArray(t *testing.T) {
	assert.Equal(t, EmptyArrayFunc{}.call(), EmptyArrayFunc{}.call().Pop())
}

func TestArrayPopFG_givenEmptyArray_returnsEmptyArray(t *testing.T) {
	assertReducesTo(t, "Array{Arr{0, 0, 0, 0, 0}, Zero{}}", "EmptyArrayFunc{}.call().Pop()")
}

func TestArray_givenElementIsPushed_appearsAtEndOfResultingArray(t *testing.T) {
	arr := EmptyArrayFunc{}.call().Push(21)
	assert.Equal(t, 21, arr.Get(arr.Len().pred()))
}

func TestArrayFG_givenElementIsPushed_appearsAtEndOfResultingArray(t *testing.T) {
	assertReducesTo(t, "21",
		`EmptyArrayFunc{}.call().Push(21)
		.Get(EmptyArrayFunc{}.call().Push(21).Len().pred())`)
}

func TestArray_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	assert.Equal(t, 1, EmptyArrayFunc{}.call().Push(21).Len().val())
}

func TestArrayFG_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	assertReducesTo(t, "1", "EmptyArrayFunc{}.call().Push(21).Len().val()")
}

func TestArray_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	assert.Equal(t, 0, EmptyArrayFunc{}.call().Push(21).Pop().Len().val())
}

func TestArrayFG_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	assertReducesTo(t, "0", "EmptyArrayFunc{}.call().Push(21).Pop().Len().val()")
}

func TestArray_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(1)
	arr = arr.Push(2)

	assert.Equal(t, 2, arr.Get(arr.Len().pred()))
	arr = arr.Pop()

	assert.Equal(t, 1, arr.Get(arr.Len().pred()))
}

func TestArrayFG_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	assertReducesTo(t, "2",
		`EmptyArrayFunc{}.call().Push(1).Push(2)
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Len().pred())`)
	assertReducesTo(t, "1",
		`EmptyArrayFunc{}.call().Push(1).Push(2).Pop()
		.Get(EmptyArrayFunc{}.call().Push(1).Push(2).Pop().Len().pred())`)
}

func TestArrayGet_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(1)
	arr = arr.Push(2)

	assert.Equal(t, 1, arr.Get(Zero{}))
	assert.Equal(t, 2, arr.Get(Succ{Zero{}}))
}

func TestArrayGetFG_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	assertReducesTo(t, "1",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Zero{})")
	assertReducesTo(t, "2",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ{Zero{}})")
}

func TestArrayGet_givenAccessOutOfBounds_returnsZero(t *testing.T) {
	assert.Equal(t, 0,
		EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ{Succ{Zero{}}}))
}

func TestArrayGetFG_givenAccessOutOfBounds_returnsZero(t *testing.T) {
	assertReducesTo(t, "0",
		"EmptyArrayFunc{}.call().Push(1).Push(2).Get(Succ{Succ{Zero{}}})")
}

func TestArray_givenMoreElementsPushedThanCapacity_doesNotChangeArray(t *testing.T) {
	arr := EmptyArrayFunc{}.call()

	for i := 0; i < 5; i++ {
		arr = arr.Push(i)
	}
	assert.Equal(t, arr, arr.Push(42))
}

func TestArray_givenMoreElementsPushedThanCapacityFG_doesNotChangeArray(t *testing.T) {
	assertReducesTo(t,
		"Array{Arr{0, 1, 2, 3, 4}, Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}}",
		"EmptyArrayFunc{}.call().Push(0).Push(1).Push(2).Push(3).Push(4).Push(42)")
}

// tests both under FG and FGG interpreters
func assertReducesTo(t *testing.T, expected string, main string) {
	program := arrayLibGo + fmt.Sprintf(mainTemplate, main)

	res, err := entrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)

	res, err = fggEntrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
