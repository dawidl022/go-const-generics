package main

type ResizableArray struct {
	arr Array
	len Nat
}

func (a ResizableArray) Push(element int) ResizableArray {
	return a.Capacity().ifLessEqR(a.len,
		ResizableArrayFunction{a},
		PushFunction{a, element})
}

func (g PushFunction) call() ResizableArray {
	return ResizableArray{
		g.a.arr.set(g.a.len.value(), g.element),
		Succ{g.a.len}}
}

func (a ResizableArray) Pop() ResizableArray {
	return ResizableArray{a.arr, a.len.pred()}
}

func (a ResizableArray) Get(i Nat) int {
	return a.len.ifLessEq(i,
		IntLiteralFunction{0},
		ArrGetFunction{a.arr, i.value()})
}

func (a ArrGetFunction) call() int {
	return a.arr[a.i]
}

func (a ResizableArray) Len() Nat {
	return a.len
}

func (a ResizableArray) Capacity() Nat {
	return Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}
}

type EmptyResizableArrFunction struct {
}

func (e EmptyResizableArrFunction) call() ResizableArray {
	return ResizableArray{Array{0, 0, 0, 0, 0}, Zero{}}
}

type Array [5]int

func (a Array) set(i int, val int) Array {
	a[i] = val
	return a
}

type Nat interface {
	value() int
	pred() Nat
	ifLessEq(other Nat, ifTrue Function, ifFalse Function) int
	ifLessEqR(other Nat, ifTrue FunctionR, ifFalse FunctionR) ResizableArray
	isZero() Boolean
	isZeroR() BooleanR
}

type Boolean interface {
	eval(ifTrue Function, ifFalse Function) int
}

type BooleanR interface {
	eval(ifTrue FunctionR, ifFalse FunctionR) ResizableArray
}

type Function interface {
	call() int
}

type FunctionR interface {
	call() ResizableArray
}

type True struct {
}

type TrueR struct {
}

func (t True) eval(ifTrue Function, ifFalse Function) int {
	return ifTrue.call()
}

func (t TrueR) eval(ifTrue FunctionR, ifFalse FunctionR) ResizableArray {
	return ifTrue.call()
}

type False struct {
}

type FalseR struct {
}

func (f False) eval(ifTrue Function, ifFalse Function) int {
	return ifFalse.call()
}

func (f FalseR) eval(ifTrue FunctionR, ifFalse FunctionR) ResizableArray {
	return ifFalse.call()
}

type Zero struct{}

func (z Zero) value() int {
	return 0
}

func (z Zero) pred() Nat {
	return z
}

func (z Zero) isZero() Boolean {
	return True{}
}

func (z Zero) isZeroR() BooleanR {
	return TrueR{}
}

func (z Zero) ifLessEq(other Nat, ifTrue Function, ifFalse Function) int {
	return ifTrue.call()
}

func (z Zero) ifLessEqR(other Nat, ifTrue FunctionR, ifFalse FunctionR) ResizableArray {
	return ifTrue.call()
}

type Succ struct {
	predF Nat
}

func (s Succ) value() int {
	return s.predF.value() + 1
}

func (s Succ) pred() Nat {
	return s.predF
}

func (s Succ) isZero() Boolean {
	return False{}
}

func (s Succ) isZeroR() BooleanR {
	return FalseR{}
}

func (s Succ) ifLessEq(other Nat, ifTrue Function, ifFalse Function) int {
	return other.isZero().eval(ifFalse, IfLessEq{s.pred(), other.pred(), ifTrue, ifFalse})
}

func (s Succ) ifLessEqR(other Nat, ifTrue FunctionR, ifFalse FunctionR) ResizableArray {
	return other.isZeroR().eval(ifFalse, IfLessEqR{s.pred(), other.pred(), ifTrue, ifFalse})
}

type IfLessEq struct {
	a       Nat
	b       Nat
	ifTrue  Function
	ifFalse Function
}

type IfLessEqR struct {
	a       Nat
	b       Nat
	ifTrue  FunctionR
	ifFalse FunctionR
}

func (i IfLessEq) call() int {
	return i.a.ifLessEq(i.b, i.ifTrue, i.ifFalse)
}

func (i IfLessEqR) call() ResizableArray {
	return i.a.ifLessEqR(i.b, i.ifTrue, i.ifFalse)
}

type ResizableArrayFunction struct {
	a ResizableArray
}

func (r ResizableArrayFunction) call() ResizableArray {
	return r.a
}

type PushFunction struct {
	a       ResizableArray
	element int
}

type ArrGetFunction struct {
	arr Array
	i   int
}

type IntLiteralFunction struct {
	i int
}

func (i IntLiteralFunction) call() int {
	return i.i
}
