package main

type Array struct {
	arr Arr
	len Nat
}

func (a Array) Push(el int) Array {
	return a.Cap().ifLessEqA(a.len,
		ArrayFunc{a},
		PushFunc{a, el})
}
func (f PushFunc) call() Array {
	return Array{
		f.a.arr.set(
			f.a.len.val(), f.el),
		Succ{f.a.len}}
}

func (a Array) Pop() Array {
	return Array{a.arr, a.len.pred()}
}

func (a Array) Get(i Nat) int {
	return a.len.ifLessEq(i,
		IntFunc{0},
		ArrGetFunc{a.arr, i.val()})
}
func (f ArrGetFunc) call() int {
	return f.arr[f.i]
}

func (a Array) Len() Nat {
	return a.len
}

func (a Array) Cap() Nat {
	return Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}
}

type EmptyArrayFunc struct {
}

func (e EmptyArrayFunc) call() Array {
	return Array{Arr{0, 0, 0, 0, 0}, Zero{}}
}

type Arr [5]int

func (a Arr) set(i int, val int) Arr {
	a[i] = val
	return a
}

type Nat interface {
	val() int
	pred() Nat
	ifLessEq(other Nat, ifTrue Func, ifFalse Func) int
	ifLessEqA(other Nat, ifTrue FuncA, ifFalse FuncA) Array
	isZero() Bool
	isZeroA() BoolA
}

type Bool interface {
	eval(ifTrue Func, ifFalse Func) int
}

type BoolA interface {
	eval(ifTrue FuncA, ifFalse FuncA) Array
}

type Func interface {
	call() int
}

type FuncA interface {
	call() Array
}

type True struct {
}

type TrueA struct {
}

func (t True) eval(ifTrue Func, ifFalse Func) int {
	return ifTrue.call()
}

func (t TrueA) eval(ifTrue FuncA, ifFalse FuncA) Array {
	return ifTrue.call()
}

type False struct {
}

type FalseA struct {
}

func (f False) eval(ifTrue Func, ifFalse Func) int {
	return ifFalse.call()
}

func (f FalseA) eval(ifTrue FuncA, ifFalse FuncA) Array {
	return ifFalse.call()
}

type Zero struct{}

func (z Zero) val() int {
	return 0
}

func (z Zero) pred() Nat {
	return z
}

func (z Zero) isZero() Bool {
	return True{}
}

func (z Zero) isZeroA() BoolA {
	return TrueA{}
}

func (z Zero) ifLessEq(other Nat, ifTrue Func, ifFalse Func) int {
	return ifTrue.call()
}

func (z Zero) ifLessEqA(other Nat, ifTrue FuncA, ifFalse FuncA) Array {
	return ifTrue.call()
}

type Succ struct {
	predF Nat
}

func (s Succ) val() int {
	return s.predF.val() + 1
}

func (s Succ) pred() Nat {
	return s.predF
}

func (s Succ) isZero() Bool {
	return False{}
}

func (s Succ) isZeroA() BoolA {
	return FalseA{}
}

func (s Succ) ifLessEq(other Nat, ifTrue Func, ifFalse Func) int {
	return other.isZero().eval(
		ifFalse, IfLessEq{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ) ifLessEqA(other Nat, ifTrue FuncA, ifFalse FuncA) Array {
	return other.isZeroA().eval(
		ifFalse, IfLessEqR{s.pred(), other.pred(), ifTrue, ifFalse})
}

type IfLessEq struct {
	a       Nat
	b       Nat
	ifTrue  Func
	ifFalse Func
}

type IfLessEqR struct {
	a       Nat
	b       Nat
	ifTrue  FuncA
	ifFalse FuncA
}

func (i IfLessEq) call() int {
	return i.a.ifLessEq(i.b, i.ifTrue, i.ifFalse)
}

func (i IfLessEqR) call() Array {
	return i.a.ifLessEqA(i.b, i.ifTrue, i.ifFalse)
}

type ArrayFunc struct {
	a Array
}

func (r ArrayFunc) call() Array {
	return r.a
}

type PushFunc struct {
	a  Array
	el int
}

type ArrGetFunc struct {
	arr Arr
	i   int
}

type IntFunc struct {
	i int
}

func (i IntFunc) call() int {
	return i.i
}
