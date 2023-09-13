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

//go:embed testdata/declaration/duplicate_field_names/duplicate_field_names.go
var declarationDuplicateStructFieldNamesGo []byte

func TestTypeCheck_givenDeclarationWithDuplicateFieldNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateStructFieldNamesGo,
		`ill-typed declaration: type "Foo": field name redeclared "x"`)
}

//go:embed testdata/declaration/invalid_field_type/invalid_field_type.go
var declarationInvalidFieldTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidStructFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidFieldTypeGo,
		`ill-typed declaration: type "Foo": field "x" type name not declared: "Bar"`)
}

//go:embed testdata/declaration/duplicate_interface_method_names/duplicate_interface_method_names.go
var declarationDuplicateInterfaceMethodNamesGo []byte

func TestTypeCheck_givenInterfaceDeclarationWithDuplicateMethodNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateInterfaceMethodNamesGo,
		`ill-typed declaration: type "xGetter": method name redeclared "getX"`)
}
