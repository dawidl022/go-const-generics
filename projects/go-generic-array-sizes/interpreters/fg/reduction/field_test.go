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
	p, err := parseAndReduceOneStep(fieldBasicGo)

	require.NoError(t, err)
	require.Equal(t, "1", p.Expression.Value().String())
	require.Equal(t, ast.IntegerLiteral{IntValue: 1}, p.Expression.Value())
}

//go:embed testdata/field/multiple_fields/multiple_fields.go
var fieldMultipleFieldsGo []byte

func TestReduceField_givenValidExpressionWithMultipleFields_returnsValueOfCorrectField(t *testing.T) {
	p, err := parseAndReduceOneStep(fieldMultipleFieldsGo)

	require.NoError(t, err)
	// TODO compare against actual ast struct, rather than the string
	require.Equal(t, "2", p.Expression.Value().String())
}

//go:embed testdata/field/array_value/array_value.go
var fieldArrayValueGo []byte

func TestReduceField_givenValidExpressionWithArrayField_returnsArrayValue(t *testing.T) {
	p, err := parseAndReduceOneStep(fieldArrayValueGo)

	require.NoError(t, err)
	require.Equal(t, "Arr{1, 2}", p.Expression.Value().String())
}

//go:embed testdata/field/struct_value/struct_value.go
var fieldStructValueGo []byte

func TestReduceField_givenValidExpressionWithStructField_returnsStructValue(t *testing.T) {
	p, err := parseAndReduceOneStep(fieldStructValueGo)

	require.NoError(t, err)
	require.Equal(t, "Structure{1, 2}", p.Expression.Value().String())
}

//go:embed testdata/field/incomplete_literal/incomplete_literal.go
var fieldZeroValuesGo []byte

func TestReduceField_givenStructLiteralWithLessFieldsThanDeclared_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(fieldZeroValuesGo)

	require.Error(t, err)
	require.Equal(t, "struct literal missing value at index 1", err.Error())
}

//go:embed testdata/field/invalid_field/invalid_field.go
var fieldInvalidFieldGo []byte

func TestReduceField_givenUndeclaredField_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(fieldInvalidFieldGo)

	require.Error(t, err)
	require.Equal(t, `no field named "y" found on struct of type "Foo"`, err.Error())
}

//go:embed testdata/field/non_struct/non_struct.go
var fieldNonStructGo []byte

func TestReduceField_givenFieldAccessOnArray_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(fieldNonStructGo)

	require.Error(t, err)
	require.Equal(t, `no struct type named "Foo" found in declarations`, err.Error())
}
