package ast

import "fmt"

func (a ArrayIndex) Reduce(declarations []Declaration) (Expression, error) {
	receiver, isReceiverValue := a.Receiver.(ValueLiteral)
	if !isReceiverValue {
		reducedReceiver, err := a.Receiver.Reduce(declarations)
		return ArrayIndex{Index: a.Index, Receiver: reducedReceiver}, err
	}
	// TODO write test where this is necessary
	//rec := a.Receiver.Value()
	//if rec == nil {
	//	reducedReceiver, err := a.Receiver.Reduce(declarations)
	//	return ArrayIndex{Index: a.Index, Receiver: reducedReceiver}, err
	//}
	//receiver := rec.(ValueLiteral)
	arrTypeName := receiver.TypeName
	i := a.Index.Value().(IntegerLiteral).IntValue

	withinBounds, err := inIndexBounds(declarations, arrTypeName, i)
	if err != nil {
		return nil, err
	}
	if !withinBounds {
		return nil, fmt.Errorf("index %d out of bounds for array of type %q", i, arrTypeName)
	}
	if len(receiver.Values) <= i {
		return nil, fmt.Errorf("array literal missing value at index %d", i)
	}
	return receiver.Values[i], nil
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
