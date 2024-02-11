package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/self_ref/self_ref_generic/self_ref_generic.go
var selfRefGenericFgg []byte

func TestTypeCheck_givenGenericTypeIsSelfReferential_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericFgg,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "f" of type "Foo"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg/self_ref_in_type_arg.go
var selfRefInTypeArgFgg []byte

func TestTypeCheck_givenSelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`field "f" of type "Foo", which has: `+
			`field "b" of type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_instantiation/self_ref_instantiation.go
var selfRefInstantiationFgg []byte

// type instantiations are guaranteed to not be self-referential if the type declaration is not
func TestTypeCheck_givenTypeIsInstantiatedWithItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationFgg)
}

//go:embed testdata/self_ref/self_ref_nested/self_ref_nested.go
var selfRefNestedFgg []byte

func TestTypeCheck_givenNestedSelfReferentialType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`field "foo" of type "Foo", which has: `+
			`field "bar" of type "Bar", which has: `+
			`field "baz" of type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_indirect/self_ref_indirect.go
var selfRefIndirectFgg []byte

// this test program crashes the official Go compiler as of version 1.21.3
// TODO report compiler bug (check if still occurs in pre-release version)
func TestTypeCheck_givenIndirectSelfReferentialType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`field "bar" of type "Bar", which has: `+
			`field "foo" of type "Foo", which has: `+
			`field "baz" of type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_generic_array/self_ref_generic_array.go
var selfRefGenericArrayFgg []byte

func TestTypeCheck_givenSelfReferentialGenericArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericArrayFgg,
		`ill-typed declaration: type "Arr": circular reference: `+
			`array element type "Arr"`)
}

//go:embed testdata/self_ref/self_ref_generic_const_array/self_ref_generic_const_array.go
var selfRefGenericConstArrayFgg []byte

func TestTypeCheck_givenSelfReferentialConstGenericArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericConstArrayFgg,
		`ill-typed declaration: type "Arr": circular reference: `+
			`array element type "Arr"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg_array/self_ref_in_type_arg_array.go
var selfRefInTypeArgArrayFgg []byte

func TestTypeCheck_givenGenericArraySelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgArrayFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg_const_array/self_ref_in_type_arg_const_array.go
var selfRefInTypeArgConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArraySelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgArrayFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_indirect_array/self_ref_indirect_array.go
var selfRefIndirectArrayFgg []byte

// this program also crashes the Go compiler
func TestTypeCheck_givenIndirectSelfReferentialGenericArray_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Bar", which has: `+
			`array element type "Foo", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_indirect_const_array/self_ref_indirect_const_array.go
var selfRefIndirectConstArrayFgg []byte

func TestTypeCheck_givenIndirectSelfReferentialConstGenericArray(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Bar", which has: `+
			`array element type "Foo", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_instantiation_array/self_ref_instantiation_array.go
var selfRefInstantiationArrayFgg []byte

func TestTypeCheck_givenGenericArrayInstantiatedWithSameTypeAsItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationArrayFgg)
}

//go:embed testdata/self_ref/self_ref_instantiation_const_array/self_ref_instantiation_const_array.go
var selfRefInstantiationConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArrayInstantiatedWithSameTypeAsItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationConstArrayFgg)
}

//go:embed testdata/self_ref/self_ref_nested_array/self_ref_nested_array.go
var selfRefNestedArrayFgg []byte

func TestTypeCheck_givenGenericArrayWithNestedCircularReference_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_nested_const_array/self_ref_nested_const_array.go
var selfRefNestedConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArrayWithNestedCircularReference_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedConstArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_interface_struct/self_ref_interface_struct.go
var selfRefInterfaceStructFgg []byte

func TestTypeCheck_givenGenericInterfaceReferencingStruct_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInterfaceStructFgg)
}

//go:embed testdata/self_ref/self_ref_interface_interface/self_ref_interface_struct.go
var selfRefInterfaceInterfaceFgg []byte

func TestTypeCheck_givenGenericInterfaceReferencingItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInterfaceInterfaceFgg)
}
