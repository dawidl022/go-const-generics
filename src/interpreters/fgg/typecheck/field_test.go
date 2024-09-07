package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/field/type_param_receiver/type_param_receiver.fgg
var fieldTypeParamReceiverFgg []byte

func TestTypeCheck_givenFieldAccessOnTypeParameter_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, fieldTypeParamReceiverFgg,
		`ill-typed declaration: method "Foo.invalidField": cannot access field "y" on value of type parameter "T"`)
}

//go:embed testdata/field/generic/generic.fgg
var fieldTypeGenericFgg []byte

func TestTypeCheck_givenFieldAccessOnCorrectlyInstantiatedGenericStruct_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, fieldTypeGenericFgg)
}
