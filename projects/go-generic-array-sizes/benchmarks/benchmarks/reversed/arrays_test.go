package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedArray_createsCopyOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	arr := newIncreasingArray()
	reversedArr := reversedArray(arr)

	assert.NotEqual(t, arr, reversedArr)
}

func TestReversedArray_createsReverseOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	reversedArr := reversedArray(newIncreasingArray())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedArr[i], reversedArr[i+1])
	}
}

var resultArray [N]int

func BenchmarkReversedArray(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingArray()

	for i := 0; i < b.N; i++ {
		resultArray = reversedArray(arr)
	}
}

func newIncreasingArray() [N]int {
	var arr [N]int
	for i := 0; i < N; i++ {
		arr[i] = i
	}
	return arr
}
