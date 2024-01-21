package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedArrayGenericUnion_createsCopyOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	arr := newIncreasingArray()
	reversedArr := reversedArrayGenericUnion[int, [N]int](arr)

	assert.NotEqual(t, arr, reversedArr)
}

func TestReversedArrayGenericUnion_createsReverseOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	reversedArr := reversedArrayGenericUnion[int, [N]int](newIncreasingArray())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedArr[i], reversedArr[i+1])
	}
}

var resultArrayGenericUnion [N]int

func BenchmarkReversedArrayGenericUnion(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingArray()

	for i := 0; i < b.N; i++ {
		resultArray = reversedArrayGenericUnion[int, [N]int](arr)
	}
}
