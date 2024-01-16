package reversed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedSlice_createsCopyOfSlice(t *testing.T) {
	if N < 2 {
		return
	}

	s := newIncreasingSlice()
	reversedS := reversedSlice(s)

	assert.NotEqual(t, s, reversedS)
}

func TestReversedSlice_createsReverseOfSlice(t *testing.T) {
	if N < 2 {
		return
	}

	reversedS := reversedSlice(newIncreasingSlice())

	for i := 0; i < N-1; i++ {
		assert.Greater(t, reversedS[i], reversedS[i+1])
	}
}

func BenchmarkReversedSlice(b *testing.B) {
	b.ReportAllocs()

	arr := newIncreasingSlice()
	for i := 0; i < b.N; i++ {
		reversedSlice(arr)
	}
}

func newIncreasingSlice() []int {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i
	}
	return s
}
