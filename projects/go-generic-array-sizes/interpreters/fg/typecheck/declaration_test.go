package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/declaration/invalid_array_type/invalid_array_type.go
var declarationInvalidArrayTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidArrayTypeGo,
		`ill-typed declaration: type "Arr": element type name not declared: "Bar"`)
}

//go:embed testdata/declaration/zero_size_array/zero_size_array.go
var declarationZeroSizeArrayGo []byte

func TestTypeCheck_givenDeclarationWithZeroSizeArray_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(declarationZeroSizeArrayGo)
	require.NoError(t, err)
}

//go:embed testdata/declaration/invalid_specification_return_type/invalid_specification_return_type.go
var declarationInvalidSpecificationReturnTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidSpecificationReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidSpecificationReturnTypeGo,
		`ill-typed declaration: type "Foo": method specification "something": return type name not declared: "thing"`)
}

//go:embed testdata/declaration/invalid_specification_param_type/invalid_argument_param_type.go
var declarationInvalidSpecificationParamReturnTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidSpecificationParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidSpecificationParamReturnTypeGo,
		`ill-typed declaration: type "Foo": method specification "something": argument "y" type name not declared: "thing"`)
}

//go:embed testdata/declaration/duplicate_specification_param_names/duplicate_specification_param_names.go
var declarationDuplicateSpecificationParamNamesGo []byte

func TestTypeCheck_givenDeclarationWithDuplicateParameterNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateSpecificationParamNamesGo,
		`ill-typed declaration: type "Foo": method specification "something": argument name redeclared "x"`)
}
