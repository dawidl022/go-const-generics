package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/self_ref_field/self_ref_field.go
var declarationSelfRefFieldGo []byte

func TestTypeCheck_givenStructDeclarationWithSelfReferentialFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationSelfRefFieldGo,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "foo" of type "Foo"`)
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

//go:embed testdata/type_declaration/self_ref_array/self_ref_array.go
var declarationSelfRefArrayGo []byte

func TestTypeCheck_givenArrayDeclarationWhoseElementTypeReferencesItself_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationSelfRefArrayGo,
		`ill-typed declaration: type "Arr": circular reference: `+
			`array element type "Arr"`)
}

//go:embed testdata/type_declaration/indirect_self_ref_array/indirect_self_ref_array.go
var declarationIndirectSelfRefArrayGo []byte

func TestTypeCheck_givenArrayDeclarationWithCircularElementType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationIndirectSelfRefArrayGo,
		`ill-typed declaration: type "ArrFirst": circular reference: `+
			`array element type "ArrSecond", which has: `+
			`array element type "ArrFirst"`)
}

//go:embed testdata/type_declaration/nested_self_ref_array/nested_self_ref_array.go
var declarationNestedSelfRefArrayGo []byte

func TestTypeCheck_givenArrayDeclarationReferencingCircularType(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationNestedSelfRefArrayGo,
		`ill-typed declaration: type "ArrFirst": circular reference: `+
			`array element type "ArrSecond", which has: `+
			`array element type "ArrSecond"`)
}

// TODO self ref interfaces (direct and indirect) - should be allowed
// TODO when type arg references tupe
