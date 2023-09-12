package ast

import (
	"fmt"
)

func (s Select) Reduce(declarations []Declaration) (Expression, error) {
	if s.Receiver.Value() == nil {
		return s.withReducedReceiver(declarations)
	}
	receiver, isReceiverValue := s.Receiver.(ValueLiteral)
	if !isReceiverValue {
		return nil, fmt.Errorf("cannot access field %q on primitive value %s", s.FieldName, s.Receiver)
	}

	structFields, err := fields(declarations, receiver.TypeName)
	if err != nil {
		return nil, err
	}
	return s.reduceToField(structFields, receiver)
}

func (s Select) withReducedReceiver(declarations []Declaration) (Expression, error) {
	reducedReceiver, err := s.Receiver.Reduce(declarations)
	return Select{FieldName: s.FieldName, Receiver: reducedReceiver}, err
}

func fields(declarations []Declaration, structTypeName string) ([]Field, error) {
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(TypeDeclaration)

		if isTypeDecl {
			structTypeLit, isStructTypeLit := typeDecl.TypeLiteral.(StructTypeLiteral)
			if isStructTypeLit && typeDecl.TypeName == structTypeName {
				return structTypeLit.Fields, nil
			}
		}
	}
	return nil, fmt.Errorf("no struct type named %q found in declarations", structTypeName)
}

func (s Select) reduceToField(structFields []Field, structure ValueLiteral) (Expression, error) {
	for i, field := range structFields {
		if field.Name == s.FieldName {
			values := s.Receiver.Value().(ValueLiteral).Values
			if len(values) <= i {
				return nil, fmt.Errorf("struct literal missing value at index %d", i)
			}
			return values[i], nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, structure.TypeName)
}

func (s Select) Value() Value {
	return nil
}
