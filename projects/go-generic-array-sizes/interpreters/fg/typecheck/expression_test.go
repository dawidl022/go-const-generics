package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/expression/unbound_variable/unbound_variable.go
var expressionUnboundVariableGo []byte

func TestTypeCheck_givenVariableInMainFunc_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionUnboundVariableGo,
		`ill-typed main expression: unbound variable: "x"`)
}

//go:embed testdata/expression/struct_int_field_int_literal/struct_int_field_int_literal.go
var expressionStructIntFieldIntLiteralGo []byte

func TestTypeCheck_givenStructIntFieldInstantiatedWithIntLiteral_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(expressionStructIntFieldIntLiteralGo)
	require.NoError(t, err)
}

//go:embed testdata/expression/struct_any_field_int_literal/struct_any_field_int_literal.go
var expressionStructAnyFieldIntLiteralGo []byte

func TestTypeCheck_givenStructAnyFieldInstantiatedWithIntLiteral_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(expressionStructAnyFieldIntLiteralGo)
	require.NoError(t, err)
}

//go:embed testdata/expression/struct_invalid_field_int_literal/struct_invalid_field_int_literal.go
var expressionStructInvalidFieldIntLiteralGo []byte

func TestTypeCheck_givenIncompatibleStructFieldWithIntLiteral_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructInvalidFieldIntLiteralGo,
		`ill-typed main expression: `+
			`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar"`)
}

//go:embed testdata/expression/struct_nonempty_interface_field_int_literal/struct_nonempty_interface_field_int_literal.go
var expressionStructNonEmptyInterfaceFieldIntLiteralGo []byte

func TestTypeCheck_givenNonEmptyInterfaceStructField_cannotAssignIntLiteralToTheField(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructNonEmptyInterfaceFieldIntLiteralGo,
		`ill-typed main expression: `+
			`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar": missing methods: "something() int"`)
}

//go:embed testdata/expression/undeclared_value_literal_type/undeclared_value_literal_type.go
var expressionUndeclaredValueLiteralTypeGo []byte

func TestTypeCheck_givenValueLiteralTypeUndeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionUndeclaredValueLiteralTypeGo,
		`ill-typed main expression: undeclared value literal type name: "Foo"`)
}

//go:embed testdata/expression/struct_field_invalid_value_type/struct_field_invalid_value_type.go
var expressionStructFieldInvalidValueTypeGo []byte

func TestTypeCheck_givenStructWithInvalidValueLiteralAsField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructFieldInvalidValueTypeGo,
		`ill-typed main expression: undeclared value literal type name: "Bar"`)
}
