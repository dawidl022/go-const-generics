package semidynamicarray

import "testing"

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
			resultArr.push(j)
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
			resultArr.push(j)
		}

		for j := 0; j < N; j++ {
			resultArr.pop()
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
