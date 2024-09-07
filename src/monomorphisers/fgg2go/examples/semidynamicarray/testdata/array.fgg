package main

type any interface {}

type Array[N const, T any] struct {
	arr Arr[N, T]
	len Nat[N, T]
	cap Nat[N, T]
	zero T
}

func (a Array[N, T]) Push(el T) Array[N, T] {
	return a.Cap().ifLessEqA(a.len,
		ArrayFunc[N, T]{a},
		PushFunc[N, T]{a, el})
}
func (f PushFunc[N, T]) call() Array[N, T] {
	return Array[N, T]{
		f.a.arr.set(
			f.a.len.val(), f.el),
		Succ[N, T]{f.a.len},
		f.a.Cap(),
		f.a.zero}
}

func (a Array[N, T]) Pop() Array[N, T] {
	return Array[N, T]{a.arr, a.len.pred(), a.Cap(), a.zero}
}

func (a Array[N, T]) Get(i Nat[N, T]) T {
	return a.len.ifLessEq(i,
		ValueFunc[T]{a.zero},
		ArrGetFunc[N, T]{a.arr, i.val()})
}
func (f ArrGetFunc[N, T]) call() T {
	return f.arr[f.i]
}

func (a Array[N, T]) Len() Nat[N, T] {
	return a.len
}

func (a Array[N, T]) Cap() Nat[N, T] {
	return a.cap
}

type EmptyArrayFunc struct {
}

func (e EmptyArrayFunc) call() Array[5, int] {
	return Array[5, int]{
		Arr[5, int]{0, 0, 0, 0, 0}, Zero[5, int]{},
		Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Succ[5, int]{Zero[5, int]{}}}}}},
		0}
}

type Arr[N const, T any] [N]T

func (a Arr[N, T]) set(i int, val T) Arr[N, T] {
	a[i] = val;
	return a
}

type Nat[N const, T any] interface {
	val() int
	pred() Nat[N, T]
	ifLessEq(other Nat[N, T], ifTrue Func[T], ifFalse Func[T]) T
	ifLessEqA(other Nat[N, T], ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T]
	isZero() Bool[T]
	isZeroA() BoolA[N, T]
}

type Bool[T any] interface {
	eval(ifTrue Func[T], ifFalse Func[T]) T
}

type BoolA[N const, T any] interface {
	eval(ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T]
}

type Func[T any] interface {
	call() T
}

type FuncA[N const, T any] interface {
	call() Array[N, T]
}

type True[T any] struct {
}

type TrueA[N const, T any] struct {
}

func (t True[T]) eval(ifTrue Func[T], ifFalse Func[T]) T {
	return ifTrue.call()
}

func (t TrueA[N, T]) eval(ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T] {
	return ifTrue.call()
}

type False[T any] struct {
}

type FalseA[N const, T any] struct {
}

func (f False[T]) eval(ifTrue Func[T], ifFalse Func[T]) T {
	return ifFalse.call()
}

func (f FalseA[N, T]) eval(ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T] {
	return ifFalse.call()
}

type Zero[N const, T any] struct{}

func (z Zero[N, T]) val() int {
	return 0
}

func (z Zero[N, T]) pred() Nat[N, T] {
	return z
}

func (z Zero[N, T]) isZero() Bool[T] {
	return True[T]{}
}

func (z Zero[N, T]) isZeroA() BoolA[N, T] {
	return TrueA[N, T]{}
}

func (z Zero[N, T]) ifLessEq(other Nat[N, T], ifTrue Func[T], ifFalse Func[T]) T {
	return ifTrue.call()
}

func (z Zero[N, T]) ifLessEqA(other Nat[N, T], ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T] {
	return ifTrue.call()
}

type Succ[N const, T any] struct {
	predF Nat[N, T]
}

func (s Succ[N, T]) val() int {
	return s.predF.val() + 1
}

func (s Succ[N, T]) pred() Nat[N, T] {
	return s.predF
}

func (s Succ[N, T]) isZero() Bool[T] {
	return False[T]{}
}

func (s Succ[N, T]) isZeroA() BoolA[N, T] {
	return FalseA[N, T]{}
}

func (s Succ[N, T]) ifLessEq(other Nat[N, T], ifTrue Func[T], ifFalse Func[T]) T {
	return other.isZero().eval(
		ifFalse, IfLessEq[N, T]{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ[N, T]) ifLessEqA(other Nat[N, T], ifTrue FuncA[N, T], ifFalse FuncA[N, T]) Array[N, T] {
	return other.isZeroA().eval(
		ifFalse, IfLessEqR[N, T]{s.pred(), other.pred(), ifTrue, ifFalse})
}

type IfLessEq[N const, T any] struct {
	a       Nat[N, T]
	b       Nat[N, T]
	ifTrue  Func[T]
	ifFalse Func[T]
}

type IfLessEqR[N const, T any] struct {
	a       Nat[N, T]
	b       Nat[N, T]
	ifTrue  FuncA[N, T]
	ifFalse FuncA[N, T]
}

func (i IfLessEq[N, T]) call() T {
	return i.a.ifLessEq(i.b, i.ifTrue, i.ifFalse)
}

func (i IfLessEqR[N, T]) call() Array[N, T] {
	return i.a.ifLessEqA(i.b, i.ifTrue, i.ifFalse)
}

type ArrayFunc[N const, T any] struct {
	a Array[N, T]
}

func (r ArrayFunc[N, T]) call() Array[N, T] {
	return r.a
}

type PushFunc[N const, T any] struct {
	a  Array[N, T]
	el T
}

type ArrGetFunc[N const, T any] struct {
	arr Arr[N, T]
	i   int
}

type ValueFunc[T any] struct {
	val T
}

func (i ValueFunc[T]) call() T {
	return i.val
}
