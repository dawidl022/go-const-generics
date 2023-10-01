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

//go:embed testdata/call/generic_receiver/generic_receiver.fgg
var callGenericReceiverFgg []byte

func TestReduce_givenMethodCallOnGenericReceiver_reducesToBoundMethodBody(t *testing.T) {
	assertEqualAfterSingleReduction(t, callGenericReceiverFgg,
		`Foo[2, int]{Arr[2, int]{1, 2}, 3}.arr`)
}
