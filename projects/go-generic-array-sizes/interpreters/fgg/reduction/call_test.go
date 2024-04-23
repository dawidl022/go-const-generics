package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/call/generic_receiver/generic_receiver.fgg
var callGenericReceiverFgg []byte

func TestReduce_givenMethodCallOnGenericReceiver_reducesToBoundMethodBody(t *testing.T) {
	assertEqualAfterSingleReduction(t, callGenericReceiverFgg,
		`Foo[2, int]{Arr[2, int]{1, 2}, 3}.arr`)
}

//go:embed testdata/call/type_param_concrete/type_param_concrete.fgg
var callTypeParamConcreteFgg []byte

func TestReduce_givenMethodCallOnGenericReceiverWithConcreteTypeArgExpression(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamConcreteFgg,
		`Foo[10, any]{}`)
}

//go:embed testdata/call/type_param_substitution/type_param_substitution.fgg
var callTypeParamSubstitutionFgg []byte

func TestReduce_givenMethodCallWithRequiredTypeParamSubstitution_reducesToBodyWithConreteTypes(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamSubstitutionFgg,
		`Foo[2, int]{}`)
}

//go:embed testdata/call/type_param_substitution_nested/type_param_substitution_nested.fgg
var callTypeParamSubstitutionNestedFgg []byte

func TestReduce_givenMethodCallWithRequiredTypeParamSubstitutionInValueTypeValue_reducesToBodyWithConreteTypes(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamSubstitutionNestedFgg,
		`Foo[2, int, any]{Bar[2, int, any]{}, Bar[2, any, int]{}}`)
}

//go:embed testdata/call/type_param_method_receiver/type_param_method_receiver.fgg
var callTypeParamMethodReceiverFgg []byte

func TestReduce_givenMethodCallWithRequiredTypeParameterSubstitutionInMethodCall(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamMethodReceiverFgg,
		`Foo[2, int]{}.new()`)
}

//go:embed testdata/call/type_param_method_params/type_param_method_params.fgg
var callTypeParamMethodParamsFgg []byte

func TestReduce_givenMethodCallWithRequiredTypeParameterSubstitutionInMethodCallParam(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamMethodParamsFgg,
		`Foo[2, int]{}.new(Foo[2, int]{})`)
}

//go:embed testdata/call/type_param_select/type_param_select.fgg
var callTypeParamSelectFgg []byte

func TestReduce_givenMethodCallWithRequiredSubstitutionInFieldSelect(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamSelectFgg,
		`Foo[2, int]{}.a`)
}

//go:embed testdata/call/type_param_index/type_param_index.fgg
var callTypeParamIndexFgg []byte

func TestReduce_givenMethodCallWithRequiredSubstitutionInIndexExpression(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamIndexFgg,
		`Arr[int]{1, 2}[Foo[2, int]{}.a]`)
}

//go:embed testdata/call/type_param_add/type_param_add.go
var callTypeParamAddFgg []byte

func TestReduce_givenMethodCallWithRequiredSubstitutionInAdditionExpression(t *testing.T) {
	assertEqualAfterSingleReduction(t, callTypeParamAddFgg,
		`Foo[2, int]{}.a + Foo[2, int]{}.a`)
}
