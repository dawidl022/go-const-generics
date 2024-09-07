package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/int_literal/basic/basic.go
var intLiteralBasicGo []byte

func TestTypeCheck_givenBasicIntLiteralExpression_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, intLiteralBasicGo)
}
