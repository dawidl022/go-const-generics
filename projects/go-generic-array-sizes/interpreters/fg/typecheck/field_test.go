package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/field/basic/basic.go
var fieldBasicGo []byte

func TestTypeCheck_givenBasicFieldAccess_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(fieldBasicGo)
	require.NoError(t, err)
}

//go:embed testdata/field/invalid_field/invalid_field.go
var fieldInvalidFieldGo []byte

func TestTypeCheck_givenInvalidFieldAccess_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldInvalidFieldGo,
		`ill-typed main expression: no field named "x" found on struct of type "Foo"`)
}

//go:embed testdata/field/int_literal_receiver/int_literal_receiver.go
var fieldIntLiteralReceiverGo []byte

func TestTypeCheck_givenFieldAccessOnIntLiteral_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldIntLiteralReceiverGo,
		`ill-typed main expression: cannot access field "x" on primitive value of type "1"`)
}

//go:embed testdata/field/int_variable_receiver/int_variable_receiver.go
var fieldIntVariableReceiverGo []byte

func TestTypeCheck_givenFieldAccessOnIntVariable_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldIntVariableReceiverGo,
		`ill-typed declaration: method "Foo.getField": `+
			`cannot access field "field" on type "int": `+
			`no struct type named "int" found in declarations`)
}

//go:embed testdata/field/undeclared_receiver/undeclared_receiver.go
var fieldUndeclaredReceiverGo []byte

func TestTypeCheck_givenFieldAccessWithUndeclaredReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldUndeclaredReceiverGo,
		`ill-typed main expression: undeclared value literal type name: "Foo"`)
}

//go:embed testdata/field/used_as_type/used_as_type.go
var fieldUsedAsTypeGo []byte

func TestTypeCheck_givenFieldAccessExpressionUsedAsSameTypeAsField_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(fieldUsedAsTypeGo)
	require.NoError(t, err)
}

//go:embed testdata/field/used_as_supertype/used_as_supertype.go
var fieldUsedAsSupertypeGo []byte

func TestTypeCheck_givenFieldAccessExpressionUsedAsSupertype_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(fieldUsedAsSupertypeGo)
	require.NoError(t, err)
}

//go:embed testdata/field/used_as_subtype/used_as_subtype.go
var fieldUsedAsSubtypeGo []byte

func TestTypeCheck_givenFieldAccessExpressionUsedAsSubtype_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldUsedAsSubtypeGo,
		`ill-typed declaration: method "Foo.getX": `+
			`return expression of type "any" is not a subtype of "int"`)
}
