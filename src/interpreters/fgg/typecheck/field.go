package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/auxiliary"
)

func (t typeVisitor) VisitSelect(s ast.Select) (ast.Type, error) {
	receiver, err := t.typeOf(s.Receiver)
	if err != nil {
		return nil, err
	}
	receiverNamedType, err := t.receiverNamedType(s, receiver)
	if err != nil {
		return nil, err
	}
	fields, err := auxiliary.Fields(t.declarations, receiverNamedType)
	if err != nil {
		return nil, fmt.Errorf("cannot access field %q on type %q: %w", s.FieldName, receiver, err)
	}
	typeDecl := t.typeDeclarationOf(receiverNamedType.TypeName)
	substituter, err := newTypeParamSubstituter(receiverNamedType.TypeArguments, typeDecl.TypeParameters)
	if err != nil {
		return nil, err
	}
	for _, f := range fields {
		if f.Name == s.FieldName {
			return substituter.substituteTypeParams(f.Type).(ast.Type), nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, receiver)
}

func (t typeVisitor) receiverNamedType(s ast.Select, receiver ast.Type) (ast.NamedType, error) {
	switch receiver.(type) {
	case ast.NamedType:
		return receiver.(ast.NamedType), nil
	case ast.IntegerLiteral:
		return ast.NamedType{}, fmt.Errorf("cannot access field %q on primitive value of type %q", s.FieldName, receiver)
	case ast.TypeParameter:
		return ast.NamedType{}, fmt.Errorf("cannot access field %q on value of type parameter %q", s.FieldName, receiver)
	default:
		panic("unhandled receiver type")
	}
}
