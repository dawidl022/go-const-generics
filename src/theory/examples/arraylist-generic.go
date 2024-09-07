package main

import "fmt"

type List[T any] interface {
	Reverse() List[T]
}

type Array[N const, T any] [N]T

func (this Array[N, T]) Reverse() List[T] {
	return this.ReverseRec(Array[N, T]{}, 0)
}

func (this Array[N, T]) ReverseRec(target Array[N, T], i int) Array[N, T] {
	if i == N {
		return target
	}
	return this.ReverseRec(target.Set(N-i-1, (this[i])), i+1)
}

func (this Array[N, T]) Set(i int, v T) Array[N, T] {
	this[i] = v
	return this
}

func main() {
	var l List[int] = Array[2, int]{1, 2}
	fmt.Printf("%#v\n", l.Reverse())
}
