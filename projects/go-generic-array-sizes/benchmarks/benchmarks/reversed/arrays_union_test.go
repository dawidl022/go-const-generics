package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedArrayUnion_createsCopyOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	arr := newIncreasingArray()
	reversedArr := reversedArrayUnion(arr)

	assert.NotEqual(t, arr, reversedArr)
}

func TestReversedArrayUnion_createsReverseOfArray(t *testing.T) {
	if N < 2 {
		return
	}

	reversedArr := reversedArrayUnion(newIncreasingArray())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedArr[i], reversedArr[i+1])
	}
}

var resultArrayUnion [N]int

func BenchmarkReversedArrayUnion(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingArray()

	for i := 0; i < b.N; i++ {
		resultArray = reversedArrayUnion(arr)
	}
}
