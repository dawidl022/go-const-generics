package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_set_method_declaration/generic/generic.go
var arraySetMethodDeclGenericFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationUsingTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arraySetMethodDeclGenericFgg)
}

//go:embed testdata/array_set_method_declaration/subtype_value_type_param/subtype_value_type_param.go
var arraySetMethodDeclSubtypeValueTypeParamFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithValueTypeParameterSubtypeOfElementType_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arraySetMethodDeclSubtypeValueTypeParamFgg)
}

//go:embed testdata/array_set_method_declaration/invalid_subtype_value_type_param/invalid_subtype_value_type_param.go
var arraySetMethodDeclInvalidSubtypeValueTypeParamFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithValueTypeParameterNotSubtypeOfElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclInvalidSubtypeValueTypeParamFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`second parameter "v" cannot be used as element of array type "Arr": `+
			`type "T" is not a subtype of "int"`)
}

//go:embed testdata/array_set_method_declaration/renamed_type_params/renamed_type_params.go
var arraySetMethodDeclRenamedTypeParamsFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithRenamedReceiverTypeParameters_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclRenamedTypeParamsFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`receiver type parameter name "X" does not match type declaration parameter name "N"`)
}

//go:embed testdata/array_set_method_declaration/missing_type_params/missing_type_params.go
var arraySetMethodDeclMissingTypeParamsFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithMissingTypeParametersInReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclMissingTypeParamsFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`expected 2 type parameters on receiver but got 0`)
}

//go:embed testdata/array_set_method_declaration/extraneous_type_params/extraneous_type_params.go
var arraySetMethodDeclExtraneousTypeParamsFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithExtraneousTypeParametersInReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclExtraneousTypeParamsFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`expected 2 type parameters on receiver but got 3`)
}

//go:embed testdata/array_set_method_declaration/invalid_return_type/invalid_return_type.go
var arraySetMethodDeclInvalidReturnTypeFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithInvalidReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclInvalidReturnTypeFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`return type must be same as receiver type "Arr[N, T]"`)
}

//go:embed testdata/array_set_method_declaration/invalid_value_type_parameter/invalid_value_type_param.go
var arraySetMethodDeclInvalidValueTypeParamFgg []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithValueTypeNotSubtypeOfElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetMethodDeclInvalidValueTypeParamFgg,
		`ill-typed declaration: array-set method "Arr.set": `+
			`second parameter "v" cannot be used as element of array type "Arr": `+
			`type "int" is not a subtype of "T"`)
}
