package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedSlice_createsCopyOfSlice(t *testing.T) {
	s := newIncreasingSlice()
	reversedS := reversedSlice(s)

	assert.NotEqual(t, s, reversedS)
}

func TestReversedSlice_createsReverseOfSlice(t *testing.T) {
	reversedS := reversedSlice(newIncreasingSlice())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedS[i], reversedS[i+1])
	}
}

var sliceResult []int

func BenchmarkReversedSlice(b *testing.B) {
	arr := newIncreasingSlice()
	for i := 0; i < b.N; i++ {
		sliceResult = reversedSlice(arr)
	}
}

func newIncreasingSlice() []int {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i
	}
	return s
}
