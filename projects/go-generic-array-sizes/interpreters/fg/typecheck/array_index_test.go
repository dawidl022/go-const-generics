package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_index/basic/basic.go
var arrayIndexBasicGo []byte

func TestTypeCheck_givenBasicArrayIndexExpression_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayIndexBasicGo)
}

//go:embed testdata/array_index/int_literal_receiver/int_literal_receiver.go
var arrayIndexIntLiteralReceiverGo []byte

func TestTypeCheck_givenArrayIndexOnIntLiteralReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexIntLiteralReceiverGo,
		`ill-typed declaration: method "Foo.nth": `+
			`cannot perform array index on value of primitive type "1"`)
}

//go:embed testdata/array_index/int_receiver/int_receiver.go
var arrayIndexIntReceiverGo []byte

func TestTypeCheck_givenArrayIndexOnIntReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexIntReceiverGo,
		`ill-typed declaration: method "Foo.nth": `+
			`cannot perform array index on value of primitive type "int"`)
}

//go:embed testdata/array_index/struct_receiver/struct_receiver.go
var arrayIndexStructReceiverGo []byte

func TestTypeCheck_givenArrayIndexOnStructReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexStructReceiverGo,
		`ill-typed declaration: method "Foo.nth": `+
			`cannot perform array index on value of non-array type "Foo"`)
}

//go:embed testdata/array_index/interface_receiver/interface_receiver.go
var arrayIndexInterfaceReceiverGo []byte

func TestTypeCheck_givenArrayIndexOnInterfaceReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexInterfaceReceiverGo,
		`ill-typed declaration: method "Foo.nth": `+
			`cannot perform array index on value of non-array type "any"`)
}

//go:embed testdata/array_index/undeclared_receiver/undeclared_receiver.go
var arrayIndexUndeclaredReceiverGo []byte

func TestTypeCheck_givenArrayIndexOnUndeclaredReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexUndeclaredReceiverGo,
		`ill-typed declaration: method "Foo.nth": undeclared value literal type name: "Bar"`)
}

//go:embed testdata/array_index/struct_index_argument/struct_index_argument.go
var arrayIndexStructIndexArgumentGo []byte

func TestTypeCheck_givenArrayIndexWithStructIndexArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexStructIndexArgumentGo,
		`ill-typed declaration: method "Arr.nth": `+
			`cannot use value "Foo{}" as array index argument: `+
			`type "Foo" is not a subtype of "int"`)
}

//go:embed testdata/array_index/int_literal_index_argument/int_literal_index_argument.go
var arrayIndexIntLiteralIndexArgumentGo []byte

func TestTypeCheck_givenArrayIndexWithIntLiteralIndexArgumentWithinBounds_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayIndexIntLiteralIndexArgumentGo)
}

//go:embed testdata/array_index/out_of_bounds/out_of_bounds.go
var arrayIndexOutOfBoundsGo []byte

func TestTypeCheck_givenArrayIndexWithIntLiteralIndexArgumentOutOfBounds_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexOutOfBoundsGo,
		`ill-typed declaration: method "Arr.outOfBounds": `+
			`cannot access index 2 on array of type "Arr" of size 2`)
}

//go:embed testdata/array_index/undeclared_argument_type/undeclared_argument_type.go
var arrayIndexUndeclaredArgumentType []byte

func TestTypeCheck_givenArrayIndexWithUndeclaredIndexArgumentType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexUndeclaredArgumentType,
		`ill-typed declaration: method "Arr.nth": `+
			`undeclared value literal type name: "Bar"`)
}
