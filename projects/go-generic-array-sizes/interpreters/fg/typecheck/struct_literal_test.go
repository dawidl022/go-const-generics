package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/struct_literal/struct_int_field_int_literal/struct_int_field_int_literal.go
var expressionStructIntFieldIntLiteralGo []byte

func TestTypeCheck_givenStructIntFieldInstantiatedWithIntLiteral_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(expressionStructIntFieldIntLiteralGo)
	require.NoError(t, err)
}

//go:embed testdata/struct_literal/struct_any_field_int_literal/struct_any_field_int_literal.go
var expressionStructAnyFieldIntLiteralGo []byte

func TestTypeCheck_givenStructAnyFieldInstantiatedWithIntLiteral_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(expressionStructAnyFieldIntLiteralGo)
	require.NoError(t, err)
}

//go:embed testdata/struct_literal/struct_invalid_field_int_literal/struct_invalid_field_int_literal.go
var expressionStructInvalidFieldIntLiteralGo []byte

func TestTypeCheck_givenIncompatibleStructFieldWithIntLiteral_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructInvalidFieldIntLiteralGo,
		`ill-typed main expression: `+
			`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar"`)
}

//go:embed testdata/struct_literal/struct_nonempty_interface_field_int_literal/struct_nonempty_interface_field_int_literal.go
var expressionStructNonEmptyInterfaceFieldIntLiteralGo []byte

func TestTypeCheck_givenNonEmptyInterfaceStructField_cannotAssignIntLiteralToTheField(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructNonEmptyInterfaceFieldIntLiteralGo,
		`ill-typed main expression: `+
			`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar": missing methods: "something() int"`)
}

//go:embed testdata/struct_literal/struct_field_invalid_value_type/struct_field_invalid_value_type.go
var expressionStructFieldInvalidValueTypeGo []byte

func TestTypeCheck_givenStructWithInvalidValueLiteralAsField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructFieldInvalidValueTypeGo,
		`ill-typed main expression: undeclared value literal type name: "Bar"`)
}

//go:embed testdata/struct_literal/struct_int_field_struct_literal/struct_int_field_struct_literal.go
var expressionStructIntFieldStructLiteralGo []byte

func TestTypeCheck_givenStructLiteralWithStructUsedInPlaceOfIntField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructIntFieldStructLiteralGo,
		`ill-typed main expression: `+
			`cannot use "Bar{}" as field "x" of struct "Foo": `+
			`type "Bar" is not a subtype of "int"`)
}

//go:embed testdata/struct_literal/struct_insufficient_values/struct_insufficient_values.go
var expressionStructInsufficientValuesGo []byte

func TestTypeCheck_givenStructLiteralInstantiatedWithLessValuesThanRequired_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructInsufficientValuesGo,
		`ill-typed main expression: struct literal of type "Foo" requires 2 values, but got 1`)
}

//go:embed testdata/struct_literal/struct_extraneous_values/struct_extraneous_values.go
var expressionStructExtraneousValuesGo []byte

func TestTypeCheck_givenStructLiteralIntantiatedWithMoreValuesThanRequired_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructExtraneousValuesGo,
		`ill-typed main expression: struct literal of type "Foo" requires 2 values, but got 3`)
}
