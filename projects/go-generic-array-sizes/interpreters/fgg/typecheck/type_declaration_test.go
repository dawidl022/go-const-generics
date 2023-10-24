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

//go:embed testdata/type_declaration/array_missing_type_arguments/array_missing_type_arguments.go
var typeDeclArrayMissingTypeArgumentsFgg []byte

func TestTypeCheck_givenNestedArrayTypeDeclarationWithMissingTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclArrayMissingTypeArgumentsFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: expected 2 type arguments but got 1`)
}

//go:embed testdata/type_declaration/array_extraneous_type_arguments/array_extraneous_type_arguments.go
var typeDeclExtraneousTypeArgumentsFgg []byte

func TestTypeCheck_givenNestedArrayTypeWithExtraneousTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclExtraneousTypeArgumentsFgg,
		`ill-typed declaration: type "Matrix": `+
			`element type "Arr" badly instantiated: expected 2 type arguments but got 3`)
}

//go:embed testdata/type_declaration/type_param_undeclared_bound/type_param_undeclared_bound.go
var typeDeclTypeParamUndeclaredBoundFgg []byte

func TestTypeCheck_givenTypeParamWithUndeclaredBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamUndeclaredBoundFgg,
		`ill-typed declaration: type "Arr": `+
			`illegal bound of type parameter "T": type name not declared: "Foo"`)
}

//go:embed testdata/type_declaration/type_param_struct_bound/type_param_struct_bound.go
var typeDeclTypeParamStructBoundFgg []byte

// this is allowed in Go, but for simplicity, we only permit interface types and "const" as bounds
func TestTypeCheck_givenTypeParamWithStructBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamStructBoundFgg,
		`ill-typed declaration: type "Arr": `+
			`cannot use type "Foo" as bound: bound must be interface type or the keyword "const"`)
}

//go:embed testdata/type_declaration/type_param_int_literal_bound/type_param_int_literal_bound.go
var typeDeclTypeParamIntLiteralBoundFgg []byte

func TestTypeCheck_givenTypeParameterWithIntLiteralBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamIntLiteralBoundFgg,
		`ill-typed declaration: type "Arr": `+
			`cannot use type "2" as bound: bound must be interface type or the keyword "const"`)
}

//go:embed testdata/type_declaration/type_param_type_param_bound/type_param_type_param_bound.go
var typeDeclTypeParamTypeParamBoundFgg []byte

func TestTypeCheck_givenTypeParamTypeParamBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamTypeParamBoundFgg,
		`ill-typed declaration: type "Arr": `+
			`cannot use type "T" as bound: bound must be interface type or the keyword "const"`)
}

//go:embed testdata/type_declaration/type_param_recursive/type_param_recursive.go
var typeDeclTypeParamRecursiveFgg []byte

func TestTypeCheck_givenTypeParamDefinedInTermsOfItself_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamRecursiveFgg,
		`ill-typed declaration: type "Arr": `+
			`cannot use type "T" as bound: bound must be interface type or the keyword "const"`)
}

//go:embed testdata/type_declaration/type_param_scope/type_param_scope.go
var typeDeclTypeParamScopeFgg []byte

func TestTypeCheck_givenTypeParameterUsedInDefinitionOfOtherTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclTypeParamScopeFgg)
}

//go:embed testdata/type_declaration/type_param_scope_recursive/type_param_scope_recursive.go
var typeDeclTypeParamFgg []byte

func TestTypeCheck_givenTypeParameterUsedInDefinitionOfItself_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclTypeParamFgg)
}

//go:embed testdata/type_declaration/type_param_non_distinct/type_param_non_distinct.go
var typeDeclTypeParamNonDistinctFgg []byte

func TestTypeCheck_givenTypeParametersNonDistinct_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclTypeParamNonDistinctFgg,
		`ill-typed declaration: type "Foo": type parameter redeclared "T"`)
}
