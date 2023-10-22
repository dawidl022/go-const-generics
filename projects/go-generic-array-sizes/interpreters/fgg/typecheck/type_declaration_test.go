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
		`ill-typed declaration: type "Matrix": element type "T" is not a subtype of "const"`)
}

//go:embed testdata/type_declaration/array_const_element_arg/array_const_element_arg.go
var typeDeclArrayConstElementArgFgg []byte

// TODO fix error message: by formal rules, const is a subtype of any, but it cannot be used as a type argument
// since the parameter is non-const
func TestTypeCheck_givenNestedArrayTypeDeclarationWithConstElementTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayConstElementArgFgg,
		`ill-typed declaration: type "Matrix": element type "N" is not a subtype of "any"`)
}
