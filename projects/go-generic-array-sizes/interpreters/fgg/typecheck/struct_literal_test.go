package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/struct_literal/generic/generic.go
var structLiteralGenericFgg []byte

func TestTypeCheck_givenStructLiteralWithCorrectlyInstantiatedTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, structLiteralGenericFgg)
}
