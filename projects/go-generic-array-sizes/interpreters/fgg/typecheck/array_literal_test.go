package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_literal/illegal_non_const_type_arg/illegal_non_const_type_arg.go
var arrayLiteralIllegalNonConstTypeArgFgg []byte

func TestTypeCheck_givenArrayLiteralWithNonConstTypeArgumentWhereConstIsExpected_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralIllegalNonConstTypeArgFgg,
		`ill-typed main expression: type "Arr" badly instantiated: `+
			`type "int" cannot be used as const type argument`)
}

//go:embed testdata/array_literal/illegal_const_type_arg/illegal_const_type_arg.go
var arrayLiteralIllegalConstTypeArgFgg []byte

func TestTypeCheck_givenArrayLiteralWithConstTypeArgumentWhereNonConstIsExpected_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralIllegalConstTypeArgFgg,
		`ill-typed main expression: type "Arr" badly instantiated: `+
			`type "3" cannot be used as non-const type argument`)
}

//go:embed testdata/array_literal/illegal_type_arg_subtype/illegal_type_arg_subtype.go
var arrayLiteralIllegalTypeArgSubtype []byte

func TestTypeCheck_givenArrayLiteralWithTypeArgumentThatIsNotSubtypeOfTypeParamBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralIllegalTypeArgSubtype,
		`ill-typed main expression: type "Arr" badly instantiated: `+
			`type "int" is not a subtype of "fooer": missing methods: "foo() int"`)
}

//go:embed testdata/array_literal/generic/generic.go
var arrayLiteralGenericFgg []byte

func TestTypeCheck_givenArrayLiteralInstantiatedCorrectlyWithTypeArguments_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayLiteralGenericFgg)
}

//go:embed testdata/array_literal/generic_invalid_values/generic_invalid_values.go
var arrayLiteralGenericInvalidValuesFgg []byte

func TestTypeCheck_givenArrayLiteralWithValueNotSubtypeOfElementTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralGenericInvalidValuesFgg,
		`ill-typed main expression: `+
			`cannot use "foo{}" as element of array of type "Arr[2, int]": `+
			`type "foo" is not a subtype of "int"`)
}

//go:embed testdata/array_literal/generic_literal/generic_literal.go
var arrayLiteralGenericLiteralFgg []byte

func TestTypeCheck_givenArrayLiteralWithTypeParamElementArgument_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayLiteralGenericLiteralFgg)
}

//go:embed testdata/array_literal/generic_non_empty_length_literal/generic_non_empty_length_literal.go
var arrayLiteralNonEmptyGenericLengthLiteral []byte

func TestTypeCheck_givenArrayLiteralWithLengthTypeParameterAndNonEmptyElementList_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayLiteralNonEmptyGenericLengthLiteral,
		`ill-typed declaration: method "Arr.new": `+
			`cannot create array literal of type "Arr[N, T]" with non-concrete length "N"`)
}
