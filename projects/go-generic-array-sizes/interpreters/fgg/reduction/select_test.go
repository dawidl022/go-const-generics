package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/select/int_val_literal_type/int_val_literal_type.fgg
var selectIntValLiteralTypeFgg []byte

func TestReduce_givenSelectOnValueLiteralOfIntLiteralType_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, selectIntValLiteralTypeFgg,
		`type "1" is not a valid value literal type`)
}
