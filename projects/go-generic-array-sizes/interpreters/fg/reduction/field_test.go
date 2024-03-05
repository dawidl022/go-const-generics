package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

//go:embed testdata/field/basic/basic.go
var fieldBasicGo []byte

func TestReduceField_givenBasicValidExpression_returnsValueOfField(t *testing.T) {
	skipIfNotFG(t)
	p, err := parseFGAndReduceOneStep(fieldBasicGo)

	require.NoError(t, err)
	require.Equal(t, "1", p.Expression.Value().String())
	require.Equal(t, ast.IntegerLiteral{IntValue: 1}, p.Expression.Value())
}

//go:embed testdata/field/multiple_fields/multiple_fields.go
var fieldMultipleFieldsGo []byte

func TestReduceField_givenValidExpressionWithMultipleFields_returnsValueOfCorrectField(t *testing.T) {
	assertEqualAfterSingleReduction(t, fieldMultipleFieldsGo, "2")
}

//go:embed testdata/field/array_value/array_value.go
var fieldArrayValueGo []byte

func TestReduceField_givenValidExpressionWithArrayField_returnsArrayValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, fieldArrayValueGo, "Arr{1, 2}")
}

//go:embed testdata/field/struct_value/struct_value.go
var fieldStructValueGo []byte

func TestReduceField_givenValidExpressionWithStructField_returnsStructValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, fieldStructValueGo, "Structure{1, 2}")
}

//go:embed testdata/field/incomplete_literal/incomplete_literal.go
var fieldZeroValuesGo []byte

func TestReduceField_givenStructLiteralWithLessFieldsThanDeclared_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, fieldZeroValuesGo,
		"struct literal missing value at index 1")
}

//go:embed testdata/field/invalid_field/invalid_field.go
var fieldInvalidFieldGo []byte

func TestReduceField_givenUndeclaredField_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, fieldInvalidFieldGo,
		`no field named "y" found on struct of type "Foo"`)
}

//go:embed testdata/field/array_receiver/array_receiver.go
var fieldNonStructGo []byte

func TestReduceField_givenFieldAccessOnArray_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, fieldNonStructGo,
		`no struct type named "Foo" found in declarations`)
}

//go:embed testdata/field/integer_receiver/integer_receiver.go
var fieldIntegerReceiverGo []byte

func TestReduceField_givenIntegerLiteralReceiver_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, fieldIntegerReceiverGo,
		`cannot access field "base" on primitive value 1`)
}
