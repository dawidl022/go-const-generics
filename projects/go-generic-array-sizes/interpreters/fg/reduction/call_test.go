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

//go:embed testdata/call/undeclared_method/undeclared_method.go
var callUndeclaredMethodGo []byte

func TestReduceCall_givenUndeclaredMethodCall_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(callUndeclaredMethodGo)

	require.Error(t, err)
	require.Equal(t, `undeclared method "magic" on type "Foo"`, err.Error())
}

//go:embed testdata/call/integer_method/integer_method.go
var callIntegerMethodGo []byte

func TestReduceCall_givenIntegerMethodCall_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(callIntegerMethodGo)

	require.Error(t, err)
	require.Equal(t, `cannot call method "magic" on primitive value 1`, err.Error())
}

//go:embed testdata/call/identity/identity.go
var callIdentityGo []byte

func TestReduceCall_givenIdentityMethod_reducesToArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, callIdentityGo, "1")
}

//go:embed testdata/call/unbound_variable/unbound_variable.go
var callUnboundVariableGo []byte

func TestReduceCall_givenUnboundVariable_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(callUnboundVariableGo)

	require.Error(t, err)
	require.Equal(t, `cannot call method "Foo.unbound": unbound variable "x"`, err.Error())
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
