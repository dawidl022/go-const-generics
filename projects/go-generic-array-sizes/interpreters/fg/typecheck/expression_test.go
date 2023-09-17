package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

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
		`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar"`)
}

//go:embed testdata/expression/struct_nonempty_interface_field_int_literal/struct_nonempty_interface_field_int_literal.go
var expressionStructNonEmptyInterfaceFieldIntLiteralGo []byte

func TestTypeCheck_givenNonEmptyInterfaceStructField_cannotAssignIntLiteralToTheField(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionStructNonEmptyInterfaceFieldIntLiteralGo,
		`cannot use "1" as field "x" of struct "Foo": `+
			`type "1" is not a subtype of "Bar": missing methods: "something() int"`)
}
