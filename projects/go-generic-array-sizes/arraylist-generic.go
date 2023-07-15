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

type Matrix[const N, M, T any] [N]Array[M, T]
type Matrix[const N, M, T any] Array[N, Array[M, T]] // alternative syntax, not permitted in FG
type Matrix[const N, M, T any] [N][M]T // more realistic usage, not permitted in FG

func (this Matrix[N, M, T]) Map[R any](f Function[T, R]) List[R] {
	return this.MapRec(Matrix[N, M, R]{}, 0, f)
}

func (this Matrix[N, M, T]) MapRec[R any](target Matrix[N, M, R], i int, f Function[T, R]) Matrix[N, M, R] {
	if i == N {
		return target
	}
	return this.MapRec(target.Set(i, this[i].MapRec(Array[N, R]{}, 0, f)), i+1, f)
}

func (this Matrix[N, M, T]) Set(i int, v [M]T) Matrix[N, M, T] {
	this[i] = v
	return this
}

func main() {
	var l List[int] = Array{1, 2}
	fmt.Printf("%#v\n", l.Map(negate{}))


	var m List[int] = AnyMatrix2By2{{1, 2}, {3, 4}}
	fmt.Printf("%#v\n", m.Map(negate{}))
}
