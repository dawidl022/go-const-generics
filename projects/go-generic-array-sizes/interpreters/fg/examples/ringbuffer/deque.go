package main

type Deque struct {
	arr   Arr
	front Nat
	back  Nat
}

func (d Deque) PushFront(el int) Deque {
	return d.succ(d.front).ifEqD(d.back, DequeFunc{d}, PushFrontFunc{d, el})
}
func (f PushFrontFunc) call() Deque {
	return Deque{
		f.d.arr.set(f.d.front.val(), f.el),
		f.d.succ(f.d.front), f.d.back}
}

func (d Deque) PopFront() Deque {
	return d.front.ifEqD(d.back, DequeFunc{d}, PopFrontFunc{d})
}
func (f PopFrontFunc) call() Deque {
	return Deque{f.d.arr, f.d.pred(f.d.front), f.d.back}
}

func (d Deque) GetFront() int {
	return d.front.ifEq(d.back,
		IntFunc{0}, ArrGetFunc{d.arr, d.pred(d.front).val()})
}

func (d Deque) PushBack(el int) Deque {
	return d.front.ifEqD(d.pred(d.back), DequeFunc{d}, PushBackFunc{d, el})
}
func (f PushBackFunc) call() Deque {
	return Deque{
		f.d.arr.set(f.d.pred(f.d.back).val(), f.el),
		f.d.front, f.d.pred(f.d.back)}
}

func (d Deque) PopBack() Deque {
	return d.front.ifEqD(d.back, DequeFunc{d}, PopBackFunc{d})
}
func (f PopBackFunc) call() Deque {
	return Deque{f.d.arr, f.d.front, f.d.succ(f.d.back)}
}

func (d Deque) GetBack() int {
	return d.front.ifEq(d.back, IntFunc{0}, ArrGetFunc{d.arr, d.back.val()})
}

func (d Deque) pred(n Nat) Nat {
	return n.ifEqN(Zero{}, CapFunc{d}, PredFunc{n})
}

func (d Deque) succ(n Nat) Nat {
	return n.ifEqN(d.Cap(), ZeroFunc{}, SuccFunc{n})
}

func (d Deque) Cap() Nat {
	return Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}
}

type EmptyDequeFunc struct {
}

func (e EmptyDequeFunc) call() Deque {
	return Deque{Arr{0, 0, 0, 0, 0}, Zero{}, Zero{}}
}

type CapFunc struct {
	d Deque
}

func (f CapFunc) call() Nat {
	return f.d.Cap()
}

type PredFunc struct {
	n Nat
}

func (f PredFunc) call() Nat {
	return f.n.pred()
}

type ZeroFunc struct {
}

func (z ZeroFunc) call() Nat {
	return Zero{}
}

type SuccFunc struct {
	n Nat
}

func (f SuccFunc) call() Nat {
	return Succ{f.n}
}

type Arr [5 + 1]int

func (a Arr) set(i int, val int) Arr {
	a[i] = val
	return a
}

type Nat interface {
	val() int
	pred() Nat
	ifEq(other Nat, ifTrue Func, ifFalse Func) int
	ifEqN(other Nat, ifTrue FuncN, ifFalse FuncN) Nat
	ifEqD(other Nat, ifTrue FuncD, ifFalse FuncD) Deque
	isZero() Bool
	isZeroN() BoolN
	isZeroD() BoolD
}

type Bool interface {
	eval(ifTrue Func, ifFalse Func) int
}

type BoolN interface {
	eval(ifTrue FuncN, ifFalse FuncN) Nat
}

type BoolD interface {
	eval(ifTrue FuncD, ifFalse FuncD) Deque
}

type Func interface {
	call() int
}

type FuncN interface {
	call() Nat
}

type FuncD interface {
	call() Deque
}

type True struct {
}

type TrueN struct {
}

type TrueD struct {
}

func (t True) eval(ifTrue Func, ifFalse Func) int {
	return ifTrue.call()
}

func (t TrueN) eval(ifTrue FuncN, ifFalse FuncN) Nat {
	return ifTrue.call()
}

func (t TrueD) eval(ifTrue FuncD, ifFalse FuncD) Deque {
	return ifTrue.call()
}

type False struct {
}

type FalseN struct{}

type FalseD struct {
}

func (f False) eval(ifTrue Func, ifFalse Func) int {
	return ifFalse.call()
}

func (f FalseN) eval(ifTrue FuncN, ifFalse FuncN) Nat {
	return ifFalse.call()
}

func (f FalseD) eval(ifTrue FuncD, ifFalse FuncD) Deque {
	return ifFalse.call()
}

type Zero struct{}

func (z Zero) val() int {
	return 0
}

func (z Zero) pred() Nat {
	return Zero{}
}

func (z Zero) ifEq(other Nat, ifTrue Func, ifFalse Func) int {
	return other.isZero().eval(ifTrue, ifFalse)
}

func (z Zero) ifEqN(other Nat, ifTrue FuncN, ifFalse FuncN) Nat {
	return other.isZeroN().eval(ifTrue, ifFalse)
}

func (z Zero) ifEqD(other Nat, ifTrue FuncD, ifFalse FuncD) Deque {
	return other.isZeroD().eval(ifTrue, ifFalse)
}

func (z Zero) isZero() Bool {
	return True{}
}

func (z Zero) isZeroN() BoolN {
	return TrueN{}
}

func (z Zero) isZeroD() BoolD {
	return TrueD{}
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

func (s Succ) ifEq(other Nat, ifTrue Func, ifFalse Func) int {
	return other.isZero().eval(ifFalse, IfEq{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ) ifEqN(other Nat, ifTrue FuncN, ifFalse FuncN) Nat {
	return other.isZeroN().eval(ifFalse, IfEqN{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ) ifEqD(other Nat, ifTrue FuncD, ifFalse FuncD) Deque {
	return other.isZeroD().eval(ifFalse, IfEqD{s.pred(), other.pred(), ifTrue, ifFalse})
}

type IfEq struct {
	a       Nat
	b       Nat
	ifTrue  Func
	ifFalse Func
}

type IfEqN struct {
	a       Nat
	b       Nat
	ifTrue  FuncN
	ifFalse FuncN
}

type IfEqD struct {
	a       Nat
	b       Nat
	ifTrue  FuncD
	ifFalse FuncD
}

func (i IfEq) call() int {
	return i.a.ifEq(i.b, i.ifTrue, i.ifFalse)
}

func (i IfEqN) call() Nat {
	return i.a.ifEqN(i.b, i.ifTrue, i.ifFalse)
}

func (i IfEqD) call() Deque {
	return i.a.ifEqD(i.b, i.ifTrue, i.ifFalse)
}

func (s Succ) isZero() Bool {
	return False{}
}

func (s Succ) isZeroN() BoolN {
	return FalseN{}
}

func (s Succ) isZeroD() BoolD {
	return FalseD{}
}

type DequeFunc struct {
	d Deque
}

func (f DequeFunc) call() Deque {
	return f.d
}

type PushFrontFunc struct {
	d  Deque
	el int
}

type PopFrontFunc struct {
	d Deque
}

type IntFunc struct {
	i int
}

func (f IntFunc) call() int {
	return f.i
}

type ArrGetFunc struct {
	arr Arr
	i   int
}

func (f ArrGetFunc) call() int {
	return f.arr[f.i]
}

type PushBackFunc struct {
	d  Deque
	el int
}

type PopBackFunc struct {
	d Deque
}
