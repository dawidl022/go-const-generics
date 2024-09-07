package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_literal/basic/basic.go
var arrayLiteralBasicGo []byte

func TestTypeCheck_givenValidArrayLiteral_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayLiteralBasicGo)
}

//go:embed testdata/array_literal/element_subtypes/element_subtypes.go
var arrayLiteralSubtypesGo []byte

func TestTypeCheck_givenArrayLiteralSubtypeElements_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayLiteralSubtypesGo)
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

//go:embed testdata/array_literal/missing_elements/missing_elements.go
var arrayLiteralMissingElementsGo []byte

// this is allowed in Go, but for simplicity we disallow it (no zero values)
func TestTypeCheck_givenArrayLiteralWithMissingElements_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralMissingElementsGo,
		`ill-typed main expression: `+
			`expected 2 values in array literal of type "Arr" but got 1`)
}

//go:embed testdata/array_literal/extraneous_elements/extraneous_elements.go
var arrayLiteralExtraneousElementsGo []byte

func TestTypeCheck_givenArrayLiteralWithExtraneousElements_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralExtraneousElementsGo,
		`ill-typed main expression: `+
			`expected 2 values in array literal of type "Arr" but got 3`)
}
