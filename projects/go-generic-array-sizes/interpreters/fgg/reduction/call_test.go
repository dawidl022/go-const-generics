package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/call/int_val_literal_type/int_val_literal_type.fgg
var callIntValLiteralTypeFgg []byte

func TestReduce_givenMethodCallOnValueLiteralOfIntLiteralType_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, callIntValLiteralTypeFgg,
		`type "1" is not a valid value literal type`)
}
