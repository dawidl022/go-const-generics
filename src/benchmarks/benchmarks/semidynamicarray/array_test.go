package semidynamicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayLen_givenEmptyArray_returnsZero(t *testing.T) {
	arr := Array{}
	assert.Equal(t, 0, arr.Len())
}

func TestArrayGet_givenEmptyArray_panics(t *testing.T) {
	arr := Array{}
	assert.PanicsWithValue(t, "index out of bounds", func() { arr.Get(0) })
}

func TestArrayPop_givenEmptyArray_panics(t *testing.T) {
	arr := Array{}
	assert.PanicsWithValue(t, "array is empty", func() { arr.Pop() })
}

func TestArray_givenElementIsPushed_sameElementIsPopped(t *testing.T) {
	arr := Array{}
	arr.Push(21)
	assert.Equal(t, 21, arr.Pop())
}

func TestArray_givenElementIsPushed_lenIncreasesByOne(t *testing.T) {
	arr := Array{}
	arr.Push(21)
	assert.Equal(t, 1, arr.Len())
}

func TestArray_givenElementIsPushedThenPopped_lenDecreasesBackToZero(t *testing.T) {
	arr := Array{}
	arr.Push(21)
	arr.Pop()
	assert.Equal(t, 0, arr.Len())
}

func TestArray_givenTwoElementsArePushed_theyArePoppedInReverseOrder(t *testing.T) {
	arr := Array{}
	arr.Push(1)
	arr.Push(2)

	assert.Equal(t, 2, arr.Pop())
	assert.Equal(t, 1, arr.Pop())
}

func TestArrayGet_givenElementIsPushed_getReturnsElementAtIndex(t *testing.T) {
	arr := Array{}
	arr.Push(1)
	arr.Push(2)

	assert.Equal(t, 1, arr.Get(0))
	assert.Equal(t, 2, arr.Get(1))
	assert.PanicsWithValue(t, "index out of bounds", func() { arr.Get(2) })
}

func TestArray_givenMoreElementsPushedThanCapacity_panics(t *testing.T) {
	arr := Array{}

	for i := 0; i < N; i++ {
		arr.Push(i)
	}
	assert.PanicsWithValue(t, "array is full", func() { arr.Push(42) })
}

func TestArrayGet_givenNegativeIndex_panics(t *testing.T) {
	arr := Array{}
	arr.Push(1)

	assert.Panics(t, func() { arr.Get(-1) })
}

var resultArr Array

var resultSlice []int

func BenchmarkArrayAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultArr = Array{}
	}
}

func BenchmarkSliceAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultSlice = make([]int, 0, N)
	}
}

func BenchmarkArrayPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultArr = Array{}

		for j := 0; j < N; j++ {
			resultArr.Push(j)
		}
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultSlice = make([]int, 0, N)

		for j := 0; j < N; j++ {
			resultSlice = append(resultSlice, j)
		}
	}
}

func BenchmarkArrayPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultArr = Array{}

		for j := 0; j < N; j++ {
			resultArr.Push(j)
		}

		for j := 0; j < N; j++ {
			resultArr.Pop()
		}
	}
}

func BenchmarkSlicePop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultSlice = make([]int, 0, N)

		for j := 0; j < N; j++ {
			resultSlice = append(resultSlice, j)
		}

		for j := N - 1; j >= 0; j-- {
			resultSlice = resultSlice[:j]
		}
	}
}
