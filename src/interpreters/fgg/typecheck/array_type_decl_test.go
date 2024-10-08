package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/array_length_type_undeclared/array_length_type_undeclared.fgg
var typeDeclArrayLengthTypeUndeclaredFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithUndeclaredLengthType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayLengthTypeUndeclaredFgg,
		`ill-typed declaration: type "Arr": length type name not declared: "Foo"`)
}

//go:embed testdata/type_declaration/array_non_const_length_type/array_non_const_length_type.fgg
var typeDeclArrayNonConstLengthTypeFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithNonConstantLengthType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayNonConstLengthTypeFgg,
		`ill-typed declaration: type "Arr": non-const type "Foo" used as length`)
}

//go:embed testdata/type_declaration/array_non_const_length_param/array_non_const_length_param.fgg
var typeDeclArrayNonConstLengthParamFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithNonConstantTypeParameterUsedAsLength_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayNonConstLengthParamFgg,
		`ill-typed declaration: type "Arr": non-const type "T" used as length`)
}

//go:embed testdata/type_declaration/array_int_literal_element_type/array_int_literal_element_type.fgg
var typeDeclArrayIntLiteralElementTypeFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithIntLiteralElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayIntLiteralElementTypeFgg,
		`ill-typed declaration: type "Arr": const type "2" used as element type`)
}

//go:embed testdata/type_declaration/array_const_element_param/array_const_element_param.fgg
var typeDeclArrayConstElementParamFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithContTypeParameterUsedAsElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayConstElementParamFgg,
		`ill-typed declaration: type "Arr": const type "N" used as element type`)
}

//go:embed testdata/type_declaration/array_invalid_element_type_arg/array_invalid_element_type_arg.fgg
var typeDeclArrayInvalidElementTypeArgFgg []byte

func TestTypeCheck_givenArrayDeclarationWithInvalidTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayInvalidElementTypeArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: `+
			`type "T" is not a subtype of "fooer": missing methods: "foo() int"`)
}

//go:embed testdata/type_declaration/array_const_type_param/array_const_type_param.fgg
var typeDeclArrayConstTypeParamFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithConstantLengthTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclArrayConstTypeParamFgg)
}

//go:embed testdata/type_declaration/array_type_params/array_type_params.fgg
var typeDeclArrayTypeParamsFgg []byte

func TestTypeCheck_givenArrayTypeDeclarationWithBothLengthAndElementTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclArrayTypeParamsFgg)
}

//go:embed testdata/type_declaration/array_type_nested/array_type_nested.fgg
var typeDeclArrayTypeNestedFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclArrayTypeNestedFgg)
}
