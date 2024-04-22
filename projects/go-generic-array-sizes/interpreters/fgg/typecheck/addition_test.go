package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/addition/generic/generic.go
var additionGenericFgg []byte

func TestTypeCheck_givenAdditionWithTypeParametersInSubExpressions(t *testing.T) {
	assertPassesTypeCheck(t, additionGenericFgg)
}

//go:embed testdata/addition/untyped_left/untyped_left.go
var additionUntypedLeft []byte

func TestTypeCheck_givenAdditionWithUndeclaredTypeParameterInLeftSubexpression_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, additionUntypedLeft,
		`type parameter "T" does not accept any type arguments`)
}

//go:embed testdata/addition/untyped_right/untyped_right.go
var additionUntypedRight []byte

func TestTypeCheck_givenAdditionWithUndeclaredTypeParameterInRightSubexpression_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, additionUntypedRight,
		`type parameter "T" does not accept any type arguments`)
}
