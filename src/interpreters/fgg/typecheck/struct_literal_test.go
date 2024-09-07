package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/struct_literal/generic/generic.fgg
var structLiteralGenericFgg []byte

func TestTypeCheck_givenStructLiteralWithCorrectlyInstantiatedTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, structLiteralGenericFgg)
}

//go:embed testdata/struct_literal/illegal_type_arg_subtype/illegal_type_arg_subtype.fgg
var structLiteralIllegalTypeArgSubtypeFgg []byte

func TestTypeCheck_givenStructLiteralWithIncorrectlyInstantiatedTypeParameters_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, structLiteralIllegalTypeArgSubtypeFgg,
		`ill-typed main expression: type "Foo" badly instantiated: `+
			`type "int" is not a subtype of "fooer": missing methods: "foo() int"`)
}

//go:embed testdata/struct_literal/generic_invalid_values/generic_invalid_values.fgg
var structLiteralGenericInvalidValuesFgg []byte

func TestTypeCheck_givenStructLiteralWithValueNotSubtypeOfField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, structLiteralGenericInvalidValuesFgg,
		`ill-typed main expression: `+
			`cannot use "Arr[3, int]{2, 3, 4}" as field "arr" of struct "Foo": `+
			`type "Arr[3, int]" is not a subtype of "Arr[2, int]"`)
}
