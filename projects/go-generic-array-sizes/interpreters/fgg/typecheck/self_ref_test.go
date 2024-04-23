package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/self_ref/self_ref_generic/self_ref_generic.fgg
var selfRefGenericFgg []byte

func TestTypeCheck_givenGenericTypeIsSelfReferential_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericFgg,
		`ill-typed declaration: type "Foo": circular reference: `+
			`field "f" of type "Foo"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg/self_ref_in_type_arg.fgg
var selfRefInTypeArgFgg []byte

func TestTypeCheck_givenSelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`field "f" of type "Foo", which has: `+
			`field "b" of type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_instantiation/self_ref_instantiation.fgg
var selfRefInstantiationFgg []byte

// type instantiations are guaranteed to not be self-referential if the type declaration is not
func TestTypeCheck_givenTypeIsInstantiatedWithItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationFgg)
}

//go:embed testdata/self_ref/self_ref_nested/self_ref_nested.fgg
var selfRefNestedFgg []byte

func TestTypeCheck_givenNestedSelfReferentialType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`field "foo" of type "Foo", which has: `+
			`field "bar" of type "Bar", which has: `+
			`field "baz" of type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_indirect/self_ref_indirect.fgg
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

//go:embed testdata/self_ref/self_ref_generic_array/self_ref_generic_array.fgg
var selfRefGenericArrayFgg []byte

func TestTypeCheck_givenSelfReferentialGenericArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericArrayFgg,
		`ill-typed declaration: type "Arr": circular reference: `+
			`array element type "Arr"`)
}

//go:embed testdata/self_ref/self_ref_generic_const_array/self_ref_generic_const_array.fgg
var selfRefGenericConstArrayFgg []byte

func TestTypeCheck_givenSelfReferentialConstGenericArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefGenericConstArrayFgg,
		`ill-typed declaration: type "Arr": circular reference: `+
			`array element type "Arr"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg_array/self_ref_in_type_arg_array.fgg
var selfRefInTypeArgArrayFgg []byte

func TestTypeCheck_givenGenericArraySelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgArrayFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_in_type_arg_const_array/self_ref_in_type_arg_const_array.fgg
var selfRefInTypeArgConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArraySelfReferenceViaTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefInTypeArgConstArrayFgg,
		`ill-typed declaration: type "Bar": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar"`)
}

//go:embed testdata/self_ref/self_ref_indirect_array/self_ref_indirect_array.fgg
var selfRefIndirectArrayFgg []byte

// this program also crashes the Go compiler
func TestTypeCheck_givenIndirectSelfReferentialGenericArray_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Bar", which has: `+
			`array element type "Foo", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_indirect_const_array/self_ref_indirect_const_array.fgg
var selfRefIndirectConstArrayFgg []byte

func TestTypeCheck_givenIndirectSelfReferentialConstGenericArray(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectConstArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Bar", which has: `+
			`array element type "Foo", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_instantiation_array/self_ref_instantiation_array.fgg
var selfRefInstantiationArrayFgg []byte

func TestTypeCheck_givenGenericArrayInstantiatedWithSameTypeAsItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationArrayFgg)
}

//go:embed testdata/self_ref/self_ref_instantiation_const_array/self_ref_instantiation_const_array.fgg
var selfRefInstantiationConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArrayInstantiatedWithSameTypeAsItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInstantiationConstArrayFgg)
}

//go:embed testdata/self_ref/self_ref_nested_array/self_ref_nested_array.fgg
var selfRefNestedArrayFgg []byte

func TestTypeCheck_givenGenericArrayWithNestedCircularReference_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_nested_const_array/self_ref_nested_const_array.fgg
var selfRefNestedConstArrayFgg []byte

func TestTypeCheck_givenConstGenericArrayWithNestedCircularReference_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedConstArrayFgg,
		`ill-typed declaration: type "Baz": circular reference: `+
			`array element type "Foo", which has: `+
			`array element type "Bar", which has: `+
			`array element type "Baz"`)
}

//go:embed testdata/self_ref/self_ref_interface_struct/self_ref_interface_struct.fgg
var selfRefInterfaceStructFgg []byte

func TestTypeCheck_givenGenericInterfaceReferencingStruct_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInterfaceStructFgg)
}

//go:embed testdata/self_ref/self_ref_interface_interface/self_ref_interface_struct.fgg
var selfRefInterfaceInterfaceFgg []byte

func TestTypeCheck_givenGenericInterfaceReferencingItself_passesTypeCheck(t *testing.T) {
	assertPassesTypeCheck(t, selfRefInterfaceInterfaceFgg)
}

//go:embed testdata/self_ref_type_param/recursive_bound_type/recursive_bound_type.fgg
var selfRefRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenBoundReferencingTypeBeingDeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefRecursiveBoundTypeFgg,
		`ill-typed declaration: type "Eq": circular reference via type parameter: `+
			`bound of "T" references "Eq"`)
}

//go:embed testdata/self_ref_type_param/nested_recursive_bound_type/nested_recursive_bound_type.fgg
var selfRefNestedRecursiveBoundTypeFGG []byte

func TestTypeCheck_givenNestedBoundReferencingTypeBeingDeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefNestedRecursiveBoundTypeFGG,
		`ill-typed declaration: type "Eq": circular reference via type parameter: `+
			`bound of "T" references "Eq"`)
}

//go:embed testdata/self_ref_type_param/indirect_recursive_bound_type/indirect_recursive_bound_type.fgg
var selfRefIndirectRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenTypeDeclarationWithCircularlyDefinedBounds_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectRecursiveBoundTypeFgg,
		`ill-typed declaration: type "Foo": circular reference via type parameter: `+
			`bound of "T" references "Bar", where: `+
			`bound of "T" references "Foo"`)
}

//go:embed testdata/self_ref_type_param/method_recursive_bound_type/method_recursive_bound_type.fgg
var selfRefMethodRecursiveBoundTypeFgg []byte

// TODO figure this out
//
// Interestingly, this program passes the Go type checker, whereas the next program
// (which differs only in type declaration order) does not. This inconsistency
// seems to me like a compiler bug. From, the spec, it is unclear to me, which is
// the desired behaviour.
//
// If we inline the Bar constraint, then the Go type checker DOES reject the program,
// which is shown in the Go spec.
func TestTypeCheck_givenTypeDeclarationWithCircularDefinedBoundsViaInterfaceMethod_returnsError(t *testing.T) {
	t.Skip()
	assertFailsTypeCheckWithError(t, selfRefMethodRecursiveBoundTypeFgg,
		``)
}

//go:embed testdata/self_ref_type_param/method_rejected_recursive_bound_type/method_rejected_recursive_bound_type.fgg
var selfRefMethodRejectedRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenTypeDeclarationWithCircularDefinedBoundsViaInterfaceMethodInAnotherOrder_returnsError(t *testing.T) {
	t.Skip()
	assertFailsTypeCheckWithError(t, selfRefMethodRejectedRecursiveBoundTypeFgg,
		``)
}

//go:embed testdata/self_ref_type_param/method_indirect_recursive_bound_type/method_indirect_recursive_bound_type.fgg
var selfRefMethodIndirectRecursiveBoundTypeFgg []byte

// situation is analogous: this program passes the Go type checker, whereas a different
// permutation of type declarations (e.g. the next test program) does not
func TestTypeCheck_givenTypeDeclarationWithIndirectlyCircularDefinedBoundsViaInterfaceMethod_returnsError(t *testing.T) {
	t.Skip()
	assertFailsTypeCheckWithError(t, selfRefMethodIndirectRecursiveBoundTypeFgg,
		``)
}

//go:embed testdata/self_ref_type_param/method_rejected_indirect_recursive_bound_type/method_rejected_indirect_recursive_bound_type.fgg
var selfRefMethodRejectedIndirectRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenTypeDeclarationWithIndirectlyCircularDefinedBoundsViaInterfaceMethodInAnotherOrder_returnsError(t *testing.T) {
	t.Skip()
	assertFailsTypeCheckWithError(t, selfRefMethodIndirectRecursiveBoundTypeFgg,
		``)
}
