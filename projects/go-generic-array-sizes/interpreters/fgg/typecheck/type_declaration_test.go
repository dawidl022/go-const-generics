package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/array_undeclared_length_arg/array_undeclared_length_arg.go
var typeDeclArrayUndeclaredLengthArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithUndeclaredLengthTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayUndeclaredLengthArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: type name not declared: "Foo"`)
}

//go:embed testdata/type_declaration/array_undeclared_element_arg/array_undeclared_element_arg.go
var typeDeclArrayUndeclaredElementArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithUndeclaredElementTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayUndeclaredElementArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: type name not declared: "Foo"`)
}

//go:embed testdata/type_declaration/array_non_const_length_arg/array_non_const_length_arg.go
var typeDeclArrayNonConstLengthArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithNonConstLengthTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayNonConstLengthArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: type "T" cannot be used as const type argument`)
}

//go:embed testdata/type_declaration/array_const_element_arg/array_const_element_arg.go
var typeDeclArrayConstElementArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithConstElementTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayConstElementArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: type "N" cannot be used as non-const type argument`)
}

//go:embed testdata/type_declaration/array_int_literal_length_arg/array_int_literal_length_arg.go
var typeDeclArrayIntLiteralLengthArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithIntLiteralLengthTypeArgument_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclArrayIntLiteralLengthArgFgg)
}

//go:embed testdata/type_declaration/array_int_literal_element_arg/array_int_literal_element_arg.go
var typeDeclArrayIntLiteralElementArgFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithIntLiteralElementArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayIntLiteralElementArgFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: type "10" cannot be used as non-const type argument`)
}

//go:embed testdata/type_declaration/type_param_struct_bound/type_param_struct_bound.go
var typeDeclTypeParamStructBoundFgg []byte

// this is allowed in Go, but for simplicity, we only permit interface types and "const" as bounds
func TestTypeCheck_givenTypeParamWithStructBound_returnsError(t *testing.T) {
}
