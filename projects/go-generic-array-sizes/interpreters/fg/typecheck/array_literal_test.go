package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/array_literal/basic/basic.go
var arrayLiteralBasicGo []byte

func TestTypeCheck_givenValidArrayLiteral_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(arrayLiteralBasicGo)
	require.NoError(t, err)
}

//go:embed testdata/array_literal/element_subtypes/element_subtypes.go
var arrayLiteralSubtypesGo []byte

func TestTypeCheck_givenArrayLiteralSubtypeElements_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(arrayLiteralSubtypesGo)
	require.NoError(t, err)
}

//go:embed testdata/array_literal/invalid_element_type/invalid_element_type.go
var arrayLiteralInvalidElementTypeGo []byte

func TestTypeCheck_givenArrayLiteralWithInvalidElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralInvalidElementTypeGo,
		`ill-typed main expression: `+
			`cannot use "Arr{1, 2}" as element of array of type "Arr": `+
			`type "Arr" is not a subtype of "int"`)
}

//go:embed testdata/array_literal/undeclared_element_type/undeclared_element_type.go
var arrayLiteralUndeclaredElementTypeGo []byte

func TestTypeCheck_givenArrayLiteralWithUndeclaredElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralUndeclaredElementTypeGo,
		`ill-typed main expression: undeclared value literal type name: "Bar"`)
}
