package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedArrayGeneric_createsCopyOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	arr := newIncreasingArray()
	reversedArr := reversedArrayGeneric(arr)

	assert.NotEqual(t, arr, reversedArr)
}

func TestReversedArrayGeneric_createsReverseOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	reversedArr := reversedArrayGeneric(newIncreasingArray())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedArr[i], reversedArr[i+1])
	}
}

var resultArrayGeneric [N]int

func BenchmarkReversedArrayGeneric(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingArray()

	for i := 0; i < b.N; i++ {
		resultArray = reversedArrayGeneric(arr)
	}
}
