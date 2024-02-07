package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/invalid_array_type/invalid_array_type.go
var declarationInvalidArrayTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidArrayTypeGo,
		`ill-typed declaration: type "Arr": element type name not declared: "Bar"`)
}

//go:embed testdata/type_declaration/zero_size_array/zero_size_array.go
var declarationZeroSizeArrayGo []byte

func TestTypeCheck_givenDeclarationWithZeroSizeArray_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, declarationZeroSizeArrayGo)
}

//go:embed testdata/type_declaration/invalid_specification_return_type/invalid_specification_return_type.go
var declarationInvalidSpecificationReturnTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidSpecificationReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidSpecificationReturnTypeGo,
		`ill-typed declaration: type "Foo": method specification "something": return type name not declared: "thing"`)
}

//go:embed testdata/type_declaration/invalid_specification_param_type/invalid_argument_param_type.go
var declarationInvalidSpecificationParamReturnTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidSpecificationParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidSpecificationParamReturnTypeGo,
		`ill-typed declaration: type "Foo": method specification "something": parameter "y": type name not declared: "thing"`)
}

//go:embed testdata/type_declaration/duplicate_specification_param_names/duplicate_specification_param_names.go
var declarationDuplicateSpecificationParamNamesGo []byte

func TestTypeCheck_givenDeclarationWithDuplicateParameterNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateSpecificationParamNamesGo,
		`ill-typed declaration: type "Foo": method specification "something": argument name redeclared "x"`)
}

//go:embed testdata/type_declaration/duplicate_field_names/duplicate_field_names.go
var declarationDuplicateStructFieldNamesGo []byte

func TestTypeCheck_givenDeclarationWithDuplicateFieldNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateStructFieldNamesGo,
		`ill-typed declaration: type "Foo": field name redeclared "x"`)
}

//go:embed testdata/type_declaration/invalid_field_type/invalid_field_type.go
var declarationInvalidFieldTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidStructFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidFieldTypeGo,
		`ill-typed declaration: type "Foo": field "x": type name not declared: "Bar"`)
}

//go:embed testdata/type_declaration/duplicate_interface_method_names/duplicate_interface_method_names.go
var declarationDuplicateInterfaceMethodNamesGo []byte

func TestTypeCheck_givenInterfaceDeclarationWithDuplicateMethodNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateInterfaceMethodNamesGo,
		`ill-typed declaration: type "xGetter": method name redeclared "getX"`)
}

//go:embed testdata/type_declaration/duplicate_interface_methods/duplicate_interface_methods.go
var declarationDuplicateInterfaceMethods []byte

func TestTypeCheck_givenInterfaceDeclarationWithDuplicateMethods_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDuplicateInterfaceMethods,
		`ill-typed declaration: type "xGetter": method name redeclared "getX"`)
}

//go:embed testdata/type_declaration/self_ref_field/self_ref_field.go
var declarationSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWithSelfReferentialFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: field "foo" of type "Foo"`)
}

//go:embed testdata/type_declaration/indirect_self_ref_field/indirect_self_ref_field.go
var declarationIndirectSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWithCircularFieldTypeReferences_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationIndirectSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "bar" of type "Bar", which has: field "foo" of type "Foo"`)
}

//go:embed testdata/type_declaration/double_indirect_self_ref_field/double_indirect_self_ref_field.go
var declarationDoubleIndirectSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWithDoublyIndirectCircularFieldTypeReferences_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationDoubleIndirectSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "bar" of type "Bar", which has: `+
			`field "baz" of type "Baz", which has: `+
			`field "foo" of type "Foo"`)
}

//go:embed testdata/type_declaration/nested_self_ref_field/nested_self_ref_field.go
var declarationNestedSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWhoseFieldReferencesSelfReferentialStructType_returnsError(t *testing.T) {
	// naive approach could lead to infinite loop, hence the need for this test case
	assertFailsTypeCheckWithError(t, declarationNestedSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "bar" of type "Bar", which has: `+
			`field "bar" of type "Bar"`)
}

//go:embed testdata/type_declaration/nested_indirect_self_ref_field/nested_indirect_self_ref_field.go
var declarationNestedIndirectSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWhoseFieldReferencesStructWithCircularFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationNestedIndirectSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "bar" of type "Bar", which has: `+
			`field "baz" of type "Baz", which has: `+
			`field "bar" of type "Bar"`)
}

// TODO self ref (direct and indirect) array types
// TODO self ref interfaces (direct and indirect) - should be allowed
