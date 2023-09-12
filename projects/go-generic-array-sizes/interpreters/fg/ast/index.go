package ast

import (
	"fmt"
)

func (a ArrayIndex) Reduce(declarations []Declaration) (Expression, error) {
	if a.Receiver.Value() == nil {
		return a.withReducedReceiver(declarations)
	}
	receiver, isReceiverValue := a.Receiver.(ValueLiteral)
	if !isReceiverValue {
		return nil, fmt.Errorf("cannot access index on primitive value %s", a.Receiver)
	}

	if a.Index.Value() == nil {
		return a.withReducedIndex(declarations)
	}
	intArgument, isIntArgument := a.Index.(IntegerLiteral)
	if !isIntArgument {
		return nil, fmt.Errorf("non integer value %q used as index argument", a.Index)
	}

	return a.reduceToIndex(declarations, receiver, intArgument)
}

func (a ArrayIndex) withReducedReceiver(declarations []Declaration) (Expression, error) {
	reducedReceiver, err := a.Receiver.Reduce(declarations)
	return ArrayIndex{Index: a.Index, Receiver: reducedReceiver}, err
}

func (a ArrayIndex) withReducedIndex(declarations []Declaration) (Expression, error) {
	reducedIndex, err := a.Index.Reduce(declarations)
	return ArrayIndex{Index: reducedIndex, Receiver: a.Receiver}, err
}

func (a ArrayIndex) reduceToIndex(declarations []Declaration, receiver ValueLiteral, intArgument IntegerLiteral) (Expression, error) {
	index := intArgument.IntValue

	withinBounds, err := inIndexBounds(declarations, receiver.TypeName, index)
	if err != nil {
		return nil, err
	}
	if !withinBounds {
		return nil, fmt.Errorf("index %d out of bounds for array of type %q", index, receiver.TypeName)
	}
	if len(receiver.Values) <= index {
		return nil, fmt.Errorf("array literal missing value at index %d", index)
	}
	return receiver.Values[index], nil
}

func inIndexBounds(declarations []Declaration, arrayTypeName string, n int) (bool, error) {
	if n < 0 {
		return false, nil
	}
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(TypeDeclaration)

		if isTypeDecl {
			arrayTypeLit, isArrayTypeLit := typeDecl.TypeLiteral.(ArrayTypeLiteral)
			if isArrayTypeLit && typeDecl.TypeName == arrayTypeName {
				return n < arrayTypeLit.Length, nil
			}
		}
	}
	return false, fmt.Errorf("no array type named %q found in declarations", arrayTypeName)
}

func (a ArrayIndex) Value() Value {
	return nil
}
