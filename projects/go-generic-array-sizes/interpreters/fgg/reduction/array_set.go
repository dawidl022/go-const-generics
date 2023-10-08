package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (r ReducingVisitor) isArraySetMethod(receiverType ast.NamedType, methodName string) bool {
	for _, decl := range r.declarations {
		methodDecl, isArraySetMethodDecl := decl.(ast.ArraySetMethodDeclaration)
		if isArraySetMethodDecl && matchesMethod(methodDecl, receiverType, methodName) {
			return true
		}
	}
	return false
}

func (r ReducingVisitor) reduceArraySetMethodCall(m ast.MethodCall, receiver ast.ValueLiteral, receiverType ast.NamedType) (ast.Expression, error) {
	index, err := r.getArraySetIndex(m, receiverType)
	if err != nil {
		return nil, err
	}
	return reduceArrayValues(m, receiver, receiverType, index)
}

func reduceArrayValues(m ast.MethodCall, receiver ast.ValueLiteral, receiverType ast.NamedType, index ast.IntegerLiteral) (ast.Expression, error) {
	reducedArrayValues := make([]ast.Expression, len(receiver.Values))
	copy(reducedArrayValues, receiver.Values)

	reducedArrayValues[index.IntValue] = m.Arguments[1]

	return ast.ValueLiteral{
		Type:   receiverType,
		Values: reducedArrayValues,
	}, nil
}

func (r ReducingVisitor) getArraySetIndex(m ast.MethodCall, receiverType ast.NamedType) (ast.IntegerLiteral, error) {
	if len(m.Arguments) != 2 {
		return ast.IntegerLiteral{}, fmt.Errorf(
			`expected 2 arguments in call to "%s.%s", but got %d`,
			receiverType, m.MethodName, len(m.Arguments),
		)
	}

	index, isIntIndex := m.Arguments[0].(ast.IntegerLiteral)
	if !isIntIndex {
		return ast.IntegerLiteral{}, fmt.Errorf(
			`non-integer index %q in array set method call: %s.%s`,
			m.Arguments[0], receiverType, m.MethodName,
		)
	}
	withinBounds, err := inIndexBounds(r.declarations, receiverType, index.IntValue)
	if err != nil {
		return ast.IntegerLiteral{}, err
	}
	if !withinBounds {
		return ast.IntegerLiteral{}, fmt.Errorf(
			"array set index %d out of bounds for array of type %q",
			index.IntValue, receiverType,
		)
	}
	return index, nil
}
