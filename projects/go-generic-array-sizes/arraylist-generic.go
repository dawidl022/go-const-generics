package main

import "fmt"

type Function[T, R any] interface {
	Apply(x T) R
}

type negate struct{}

func (negate) Apply(x int) int {
	return -x
}

type List[T any] interface {
	Map[R any](f Function[T, R]) List[R]
}

type Array[const N, T any] [N]T

func (this Array[N, T]) Map[R any](f Function[T, R]) List[R] {
	return this.MapRec(Array[N, R]{}, 0, f)
}

func (this Array[N, T]) MapRec[R any](target Array[N, R], i int, f Function[T, R]) Array[N, R] {
	if i == N {
		return target
	}
	return this.MapRec(target.Set(i, f.Apply(this[i])), i+1, f)
}

func (this Array[N, T]) Set(i int, v T) Array[N, T] {
	this[i] = v
	return this
}

func main() {
	var l List = Array{1, 2}
	fmt.Printf("%#v\n", l.Map(negate{}))
}
