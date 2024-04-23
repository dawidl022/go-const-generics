package main

type Array[N const] struct {
	arr Arr[N]
	len Nat[N]
	cap Nat[N]
}

func (a Array[N]) Push(el int) Array[N] {
	return a.Cap().ifLessEqA(a.len,
		ArrayFunc[N]{a},
		PushFunc[N]{a, el})
}
func (f PushFunc[N]) call() Array[N] {
	return Array[N]{
		f.a.arr.set(
			f.a.len.val(), f.el),
		Succ[N]{f.a.len},
		f.a.Cap()}
}

func (a Array[N]) Pop() Array[N] {
	return Array[N]{a.arr, a.len.pred(), a.Cap()}
}

func (a Array[N]) Get(i Nat[N]) int {
	return a.len.ifLessEq(i,
		IntFunc{0},
		ArrGetFunc[N]{a.arr, i.val()})
}
func (f ArrGetFunc[N]) call() int {
	return f.arr[f.i]
}

func (a Array[N]) Len() Nat[N] {
	return a.len
}

func (a Array[N]) Cap() Nat[N] {
	return a.cap
}

type EmptyArrayFunc struct {
}

func (e EmptyArrayFunc) call() Array[5] {
	return Array[5]{
		Arr[5]{0, 0, 0, 0, 0}, Zero[5]{}, Succ[5]{Succ[5]{Succ[5]{Succ[5]{Succ[5]{Zero[5]{}}}}}}}
}

type Arr[N const] [N]int

func (a Arr[N]) set(i int, val int) Arr[N] {
	a[i] = val;
	return a
}

type Nat[N const] interface {
	val() int
	pred() Nat[N]
	ifLessEq(other Nat[N], ifTrue Func, ifFalse Func) int
	ifLessEqA(other Nat[N], ifTrue FuncA[N], ifFalse FuncA[N]) Array[N]
	isZero() Bool
	isZeroA() BoolA[N]
}

type Bool interface {
	eval(ifTrue Func, ifFalse Func) int
}

type BoolA[N const] interface {
	eval(ifTrue FuncA[N], ifFalse FuncA[N]) Array[N]
}

type Func interface {
	call() int
}

type FuncA[N const] interface {
	call() Array[N]
}

type True struct {
}

type TrueA[N const] struct {
}

func (t True) eval(ifTrue Func, ifFalse Func) int {
	return ifTrue.call()
}

func (t TrueA[N]) eval(ifTrue FuncA[N], ifFalse FuncA[N]) Array[N] {
	return ifTrue.call()
}

type False struct {
}

type FalseA[N const] struct {
}

func (f False) eval(ifTrue Func, ifFalse Func) int {
	return ifFalse.call()
}

func (f FalseA[N]) eval(ifTrue FuncA[N], ifFalse FuncA[N]) Array[N] {
	return ifFalse.call()
}

type Zero[N const] struct{}

func (z Zero[N]) val() int {
	return 0
}

func (z Zero[N]) pred() Nat[N] {
	return z
}

func (z Zero[N]) isZero() Bool {
	return True{}
}

func (z Zero[N]) isZeroA() BoolA[N] {
	return TrueA[N]{}
}

func (z Zero[N]) ifLessEq(other Nat[N], ifTrue Func, ifFalse Func) int {
	return ifTrue.call()
}

func (z Zero[N]) ifLessEqA(other Nat[N], ifTrue FuncA[N], ifFalse FuncA[N]) Array[N] {
	return ifTrue.call()
}

type Succ[N const] struct {
	predF Nat[N]
}

func (s Succ[N]) val() int {
	return s.predF.val() + 1
}

func (s Succ[N]) pred() Nat[N] {
	return s.predF
}

func (s Succ[N]) isZero() Bool {
	return False{}
}

func (s Succ[N]) isZeroA() BoolA[N] {
	return FalseA[N]{}
}

func (s Succ[N]) ifLessEq(other Nat[N], ifTrue Func, ifFalse Func) int {
	return other.isZero().eval(
		ifFalse, IfLessEq[N]{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ[N]) ifLessEqA(other Nat[N], ifTrue FuncA[N], ifFalse FuncA[N]) Array[N] {
	return other.isZeroA().eval(
		ifFalse, IfLessEqR[N]{s.pred(), other.pred(), ifTrue, ifFalse})
}

type IfLessEq[N const] struct {
	a       Nat[N]
	b       Nat[N]
	ifTrue  Func
	ifFalse Func
}

type IfLessEqR[N const] struct {
	a       Nat[N]
	b       Nat[N]
	ifTrue  FuncA[N]
	ifFalse FuncA[N]
}

func (i IfLessEq[N]) call() int {
	return i.a.ifLessEq(i.b, i.ifTrue, i.ifFalse)
}

func (i IfLessEqR[N]) call() Array[N] {
	return i.a.ifLessEqA(i.b, i.ifTrue, i.ifFalse)
}

type ArrayFunc[N const] struct {
	a Array[N]
}

func (r ArrayFunc[N]) call() Array[N] {
	return r.a
}

type PushFunc[N const] struct {
	a  Array[N]
	el int
}

type ArrGetFunc[N const] struct {
	arr Arr[N]
	i   int
}

type IntFunc struct {
	i int
}

func (i IntFunc) call() int {
	return i.i
}
