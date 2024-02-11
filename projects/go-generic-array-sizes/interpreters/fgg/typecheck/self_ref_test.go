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
func TestTypeCheck_givenIndirectSelfReferentialType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, selfRefIndirectFgg, ``)
}

// TODO try to cause infinite loop in type param substitutions while self-ref
// check is occurring
//
// e.g. when T is a type argument of a field type
