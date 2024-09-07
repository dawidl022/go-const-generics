package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/addition/simple-literal/simple-literal.go
var additionSimpleLiteralGo []byte

func TestReduceAddition_givenSimpleLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionSimpleLiteralGo,
		"2")
}

//go:embed testdata/addition/chained-literal/chained-literal.go
var additionChainedLiteralGo []byte

func TestReduceAddition_givenChainedLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionChainedLiteralGo,
		"3 + 3")
}

//go:embed testdata/addition/multi-chained-literal/multi-chained-literal.go
var additionMultiChainedLiteralGo []byte

func TestReduceAddition_givenMultiChainedLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionMultiChainedLiteralGo,
		"1 + 2 + 3")
}

//go:embed testdata/addition/variable-left/variable-left.go
var additionVariableLeftGo []byte

func TestReduceAddition_givenVariableOnLeftHandSide_substitutesVariable(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionVariableLeftGo,
		"1 + 1 + 1")
}

//go:embed testdata/addition/variable-right/variable-right.go
var additionVariableRightGo []byte

func TestReduceAddition_givenVariableOnRightHandSide_substitutesVariable(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionVariableRightGo,
		"1 + 1 + 1")
}

//go:embed testdata/addition/variable-both/variable-both.go
var additionVariableBothGo []byte

func TestReduceAddition_givenVariableOnBothSides_reducesLeftHandSideFirst(t *testing.T) {
	assertEqualAfterSingleReduction(t, additionVariableBothGo,
		"1 + 1 + Foo{}.foo(2)")
}

//go:embed testdata/addition/untyped-left/untyped-left.go
var additionUntypedLeftGo []byte

func TestReduceAddition_givenUndeclaredVariableOnLeftHandSide_propagatesError(t *testing.T) {
	assertErrorAfterSingleReduction(t, additionUntypedLeftGo,
		`unbound variable "x"`)
}

//go:embed testdata/addition/untyped-method-left/untyped-method-left.go
var additionUntypedAdditionLeftGo []byte

func TestReduceAddition_givenUndeclaredVariableOnLeftHandSideInMethod_propagatesError(t *testing.T) {
	assertErrorAfterSingleReduction(t, additionUntypedAdditionLeftGo,
		`cannot call method "Foo.foo": unbound variable "x"`)
}

//go:embed testdata/addition/untyped-right/untyped-right.go
var additionUntypedRightGo []byte

func TestReduceAddition_givenUndeclaredVariableOnRightHandSide_propagatesError(t *testing.T) {
	assertErrorAfterSingleReduction(t, additionUntypedRightGo,
		`unbound variable "x"`)
}

//go:embed testdata/addition/untyped-method-right/untyped-method-right.go
var additionUntypedAdditionRightGo []byte

func TestReduceAddition_givenUndeclaredVariableOnRightHandSideInMethod_propagatesError(t *testing.T) {
	assertErrorAfterSingleReduction(t, additionUntypedAdditionRightGo,
		`cannot call method "Foo.foo": unbound variable "x"`)
}
