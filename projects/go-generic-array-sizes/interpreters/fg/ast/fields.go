package ast

import (
	"fmt"
)

func (s Select) Reduce(declarations []Declaration) (Expression, error) {
	structure, isStructValue := s.Expression.Value().(ValueLiteral)
	if !isStructValue {
		reducedStruct, err := s.Expression.Reduce(declarations)
		return Select{FieldName: s.FieldName, Expression: reducedStruct}, err
	}

	structTypeName := structure.TypeName
	structFields, err := fields(declarations, structTypeName)
	if err != nil {
		return nil, err
	}
	for i, field := range structFields {
		if field.Name == s.FieldName {
			values := s.Expression.Value().(ValueLiteral).Values
			if len(values) <= i {
				return nil, fmt.Errorf("struct literal missing value at index %d", i)
			}
			return values[i], nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, structTypeName)
}

func (s Select) Value() Value {
	return nil
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
