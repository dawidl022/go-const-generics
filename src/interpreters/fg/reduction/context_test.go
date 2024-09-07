package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/context/index_of_field/index_of_field.go
var contextIndexOfFieldGo []byte

func TestReduce_givenIndexOfField_reducesField(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextIndexOfFieldGo, "Arr{1, 2}[1]")
}

//go:embed testdata/context/field_of_index/field_of_index.go
var contextFieldOfIndexGo []byte

func TestReduce_givenFieldOfIndex_reducesIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextFieldOfIndexGo, "Structure{1, 2}.y")
}

//go:embed testdata/context/index_of_index/index_of_index.go
var contextIndexOfIndexGo []byte

func TestReduce_givenIndexOfIndex_reducesInnerIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextIndexOfIndexGo, "Arr{3, 4}[0]")
}

//go:embed testdata/context/field_of_field/field_of_field.go
var contextFieldOfFieldGo []byte

func TestReduce_givenFieldOfField_reducesInnerField(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextFieldOfFieldGo, "Foo{1}.y")
}

//go:embed testdata/context/non_value_field/non_value_field.go
var contextNonValueFieldGo []byte

func TestReduce_givenFieldOfNonValueField_reducesToFieldOfStructLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextNonValueFieldGo, "Foo{1}.x")
}

//go:embed testdata/context/non_value_index/non_value_index.go
var contextNonValueIndexGo []byte

func TestReduce_givenIndexOfNonValueIndex_reducesToIndexOfArrayLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextNonValueIndexGo, "Arr{1}[0]")
}

//go:embed testdata/context/index_expression/index_expression.go
var contextIndexExpressionGo []byte

func TestReduce_givenNonValueIndexArgument_reducesIndexArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextIndexExpressionGo, "Arr{1, 2}[1]")
}

//go:embed testdata/context/method_of_expression/method_of_expression.go
var contextMethodOfExpressionGo []byte

func TestReduce_givenNonValueMethodReceiver_reducesMethodReceiver(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextMethodOfExpressionGo, "Foo{2}.getX()")
}

//go:embed testdata/context/method_argument_expressions/method_argument_expressions.go
var contextMethodArgumentExpressionGo []byte

func TestReduce_givenNonValueMethodArguments_reducesFirstNonValueExpression(t *testing.T) {
	assertEqualAfterSingleReduction(t, contextMethodArgumentExpressionGo, "Foo{1, 2, 3}.update(4, 5, Arr{6}[0])")
}
