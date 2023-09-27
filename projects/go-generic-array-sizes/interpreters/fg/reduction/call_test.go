package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/call/basic/basic.go
var callBasicGo []byte

func TestReduceCall_givenBasicMethodCall_reducesToReturnValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, callBasicGo, "42")
}

func TestMethodCall_isNotAValue(t *testing.T) {
	p := parseFGProgram(callBasicGo)
	require.Nil(t, p.Expression.Value())
}

//go:embed testdata/call/undeclared_method/undeclared_method.go
var callUndeclaredMethodGo []byte

func TestReduceCall_givenUndeclaredMethodCall_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUndeclaredMethodGo,
		`undeclared method "magic" on type "Foo"`)
}

//go:embed testdata/call/integer_method/integer_method.go
var callIntegerMethodGo []byte

func TestReduceCall_givenIntegerMethodCall_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callIntegerMethodGo,
		`cannot call method "magic" on primitive value 1`)
}

//go:embed testdata/call/identity/identity.go
var callIdentityGo []byte

func TestReduceCall_givenIdentityMethod_reducesToArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, callIdentityGo, "1")
}

//go:embed testdata/call/unbound_variable/unbound_variable.go
var callUnboundVariableGo []byte

func TestReduceCall_givenUnboundVariable_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundVariableGo,
		`cannot call method "Foo.unbound": unbound variable "x"`)
}

//go:embed testdata/call/receiver_identity/receiver_identity.go
var callReceiverIdentityGo []byte

func TestReduceCall_givenReceiverIdentityMethod_reducesToReceiver(t *testing.T) {
	assertEqualAfterSingleReduction(t, callReceiverIdentityGo, "Foo{1}")
}

//go:embed testdata/call/expression_body/expression_body.go
var callExpressionBodyGo []byte

func TestReduceCall_givenNonValueExpressionInBody_reducesToExpression(t *testing.T) {
	assertEqualAfterSingleReduction(t, callExpressionBodyGo, "Foo{2}.x")
}

//go:embed testdata/call/receiver_expression/receiver_expression.go
var callReceiverExpressionGo []byte

func TestReduceCall_givenNonValueExpressionOnReceiver_reducesToExpressionWithReceiverBound(t *testing.T) {
	assertEqualAfterSingleReduction(t, callReceiverExpressionGo, "Foo{1}.x")
}

//go:embed testdata/call/recursive/recursive.go
var callRecursiveGo []byte

func TestReduceCall_givenRecursiveMethodCall_reducesSingleStepOfRecursion(t *testing.T) {
	assertEqualAfterSingleReduction(t, callRecursiveGo, "Foo{}.recurse()")
}

//go:embed testdata/call/recursive_arguments/recursive_arguments.go
var callRecursiveArgumentsGo []byte

func TestReduceCall_givenRecursiveCallWithArguments_reducesSingleStepOfRecursion(t *testing.T) {
	assertEqualAfterSingleReduction(t, callRecursiveArgumentsGo, "Foo{}.recurse(1, 2)")
}

//go:embed testdata/call/unbound_method_receiver/unbound_method_receiver.go
var callUnboundMethodReceiverGo []byte

func TestReduceCall_givenUnboundMethodReceiver_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundMethodReceiverGo,
		`cannot call method "Foo.unbound": unbound variable "x"`)
}

//go:embed testdata/call/unbound_method_argument/unbound_method_argument.go
var callUnboundMethodArgumentGo []byte

func TestReduceCall_givenUnboundMethodArguments_returnsErrorWithFirstEncounteredUnboundVariable(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundMethodArgumentGo,
		`cannot call method "Foo.unbound": unbound variable "y"`)
}

//go:embed testdata/call/insufficient_arguments/insufficient_arguments.go
var callInsufficientArgumentsGo []byte

func TestReduceCall_givenLessArgumentsThanRequired_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callInsufficientArgumentsGo,
		`expected 2 argument(s) in call to "Foo.firstArg", but got 1`)
}

//go:embed testdata/call/extraneous_arguments/extraneous_arguments.go
var callExtraneousArgumentsGo []byte

func TestReduceCall_givenMoreArgumentsThanRequired_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callExtraneousArgumentsGo,
		`expected 0 argument(s) in call to "Foo.answer", but got 2`)
}

//go:embed testdata/call/struct_literal_variables/struct_literal_variables.go
var callStructLiteralVariablesGo []byte

func TestReduceCall_givenMethodWithParameterisedStructLiteral_reducesToLiteralWithBoundedValues(t *testing.T) {
	assertEqualAfterSingleReduction(t, callStructLiteralVariablesGo, "Foo{3, 4}")
}

//go:embed testdata/call/array_literal_variables/array_literal_variables.go
var callArrayLiteralVariablesGo []byte

func TestReduceCall_givenMethodWithParameterisedArrayLiteral_reducesToLiteralWithBoundedValues(t *testing.T) {
	assertEqualAfterSingleReduction(t, callArrayLiteralVariablesGo, "Arr{3, 4}")
}

//go:embed testdata/call/unbound_value_literal_variable/unbound_value_literal.go
var callUnboundValueLiteralVariable []byte

func TestReduceCall_givenMethodWithUnboundValueLiteralVariables_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundValueLiteralVariable,
		`cannot call method "Foo.unbound": unbound variable "x"`)
}

//go:embed testdata/call/index_receiver_variable/index_receiver_variable.go
var callIndexReceiverVariableGo []byte

func TestReduceCall_givenMethodWithArrayIndexOnMethodReceiver_reducesToArrayIndexOnLiteralReceiver(t *testing.T) {
	assertEqualAfterSingleReduction(t, callIndexReceiverVariableGo, "Arr{1, 2}[0]")
}

//go:embed testdata/call/index_index_variable/index_index_variable.go
var callIndexIndexVariableGo []byte

func TestReduceCall_givenMethodWithArrayIndexWithIndexVariable_reducesToArrayIndexWithIntegerLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, callIndexIndexVariableGo, "Arr{0, 1, 1, 2, 3}[3]")
}

//go:embed testdata/call/index_variables/index_variables.go
var callIndexIndexGo []byte

func TestReduceCall_givenMethodWithVariableArrayIndexReceiverAndIndex_reducesToArrayIndexWithArrayLiteralAndIntegerLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, callIndexIndexGo, "Arr{1, 2}[1]")
}

//go:embed testdata/call/unbound_index_receiver/unbound_index_receiver_variable.go
var callUnboundIndexReceiverGo []byte

func TestReduceCall_givenMethodWithUnboundArrayIndexReceiver_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundIndexReceiverGo,
		`cannot call method "Arr.unboundIndex": unbound variable "b"`)
}

//go:embed testdata/call/unbound_index_index/unbound_index_index.go
var callUnboundIndexIndexGo []byte

func TestReduceCall_givenMethodWithUnboundArrayIndexIndex_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callUnboundIndexIndexGo,
		`cannot call method "Arr.unboundIndex": unbound variable "i"`)
}
