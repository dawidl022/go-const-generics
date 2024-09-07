package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedSliceGeneric_createsCopyOfSlice(t *testing.T) {
	if N < 2 {
		return
	}

	s := newIncreasingSlice()
	reversedS := reversedSliceGeneric(s)

	assert.NotEqual(t, s, reversedS)
}

func TestReversedSliceGeneric_createsReverseOfSlice(t *testing.T) {
	if N < 2 {
		return
	}

	reversedS := reversedSliceGeneric(newIncreasingSlice())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedS[i], reversedS[i+1])
	}
}

var resultSliceGeneric []int

func BenchmarkReversedSliceGeneric(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingSlice()
	for i := 0; i < b.N; i++ {
		resultSlice = reversedSliceGeneric(arr)
	}
}
