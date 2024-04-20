package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayLen_givenEmptyArray_returnsZero(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	assert.Equal(t, 0, arr.Len().val())
}

func TestArrayGet_givenEmptyArray_returnsZero(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	assert.Equal(t, 0, arr.Get(Zero{}))
}

func TestArrayGet_givenEmptyArrayAfterPushAndPop_returnsZero(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(42)
	arr = arr.Pop()

	assert.Equal(t, 0, arr.Get(Zero{}))
}

func TestArrayPop_givenEmptyArray_returnsEmptyArray(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	assert.Equal(t, EmptyArrayFunc{}.call(), arr.Pop())
}

func TestArray_givenElementIsPushed_appearsAtEndOfResultingArray(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(21)

	assert.Equal(t, 21, arr.Get(arr.Len().pred()))
}

func TestArray_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(21)
	assert.Equal(t, 1, arr.Len().val())
}

func TestArray_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(21)
	arr = arr.Pop()
	assert.Equal(t, 0, arr.Len().val())
}

func TestArray_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(1)
	arr = arr.Push(2)

	assert.Equal(t, 2, arr.Get(arr.Len().pred()))
	arr = arr.Pop()

	assert.Equal(t, 1, arr.Get(arr.Len().pred()))
}

func TestArrayGet_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(1)
	arr = arr.Push(2)

	assert.Equal(t, 1, arr.Get(Zero{}))
	assert.Equal(t, 2, arr.Get(Succ{Zero{}}))
}

func TestArrayGet_givenAccessOutOfBounds_returnsZero(t *testing.T) {
	arr := EmptyArrayFunc{}.call()
	arr = arr.Push(1)
	arr = arr.Push(2)

	assert.Equal(t, 0, arr.Get(Succ{Succ{Zero{}}}))
}

func TestArray_givenMoreElementsPushedThanCapacity_doesNotChangeArray(t *testing.T) {
	arr := EmptyArrayFunc{}.call()

	for i := 0; i < 5; i++ {
		arr = arr.Push(i)
	}
	assert.Equal(t, arr, arr.Push(42))
}
