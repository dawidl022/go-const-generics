package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/call/basic/basic.go
var callBasicGo []byte

func TestTypeCheck_givenBasicMethodCall_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(callBasicGo)
	require.NoError(t, err)
}

//go:embed testdata/call/with_arguments/with_arguments.go
var callWithArgumentsGo []byte

func TestTypeCheck_givenMethodCallWithValidArguments_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(callWithArgumentsGo)
	require.NoError(t, err)
}

//go:embed testdata/call/with_invalid_argument_types/with_invalid_argument_types.go
var callWithInvalidArgumentTypesGo []byte

func TestTypeCheck_givenMethodCallWithInvalidArgumentTypes_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callWithInvalidArgumentTypesGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as argument "y" in call to "Foo.answer": `+
			`type "Foo" is not a subtype of "int"`)
}

//go:embed testdata/call/undeclared_receiver/undeclared_receiver.go
var callUndeclaredReceiverGo []byte

func TestTypeCheck_givenMethodCallOnUndeclaredReceiverType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callUndeclaredReceiverGo,
		`ill-typed main expression: undeclared value literal type name: "Foo"`)
}

//go:embed testdata/call/no_method_on_receiver/no_method_on_receiver.go
var callNoMethodOnReceiverGo []byte

func TestTypeCheck_givenNonExistentMethodCall_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callNoMethodOnReceiverGo,
		`ill-typed main expression: no method named "answer" on receiver of type "Foo"`)
}

//go:embed testdata/call/int_literal_receiver/int_literal_receiver.go
var callIntLiteralReceiverGo []byte

func TestTypeCheck_givenMethodCallOnIntLiteralReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callIntLiteralReceiverGo,
		`ill-typed main expression: no method named "answer" on receiver of type "42"`)
}

//go:embed testdata/call/with_argument_subtypes/with_argument_subtypes.go
var callWithArgumentSubtypesGo []byte

func TestTypeCheck_givenMethodCallWithSubtypeArguments_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(callWithArgumentSubtypesGo)
	require.NoError(t, err)
}

//go:embed testdata/call/with_undeclared_argument_type/with_undeclared_argument_type.go
var callWithUndeclaredArgumentTypeGo []byte

func TestTypeCheck_givenMethodCallWithUndeclaredArgumentType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callWithUndeclaredArgumentTypeGo,
		`ill-typed main expression: `+
			`cannot use "Bar{}" as argument "y" in call to "Foo.answer": `+
			`undeclared value literal type name: "Bar"`)
}

//go:embed testdata/call/with_missing_arguments/with_missing_arguments.go
var callWithMissingArgumentsGo []byte

func TestTypeCheck_givenMethodCallWithMissingArguments_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callWithMissingArgumentsGo,
		`ill-typed main expression: expected 2 arguments in call to "Foo.answer", but got 1`)
}

//go:embed testdata/call/with_extraneous_arguments/with_extraneous_arguments.go
var callWithExtraneousArgumentsGo []byte

func TestTypeCheck_givenMethodCallWithExtraneousArguments_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callWithExtraneousArgumentsGo,
		`ill-typed main expression: expected 2 arguments in call to "Foo.answer", but got 3`)
}

//go:embed testdata/call/as_method_expression/as_method_expression.go
var callAsMethodExpressionGo []byte

func TestTypeCheck_givenMethodCallUsedAsMethodExpressionOfTheCorrectType_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(callAsMethodExpressionGo)
	require.NoError(t, err)
}

//go:embed testdata/call/invalid_method_expression_type/invalid_method_expression.go
var callInvalidMethodExpressionType []byte

func TestTypeCheck_givenMethodCallUsedAsMethodExpressionOfAnIncorrectType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, callInvalidMethodExpressionType,
		`ill-typed declaration: method "Foo.something": `+
			`return expression of type "any" is not a subtype of "int"`)
}
