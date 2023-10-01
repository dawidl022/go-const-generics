package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (r ReducingVisitor) VisitSelect(s ast.Select) (ast.Expression, error) {
	if !s.Receiver.IsValue() {
		return r.selectWithReducedReceiver(s)
	}
	receiver, isReceiverValue := s.Receiver.(ast.ValueLiteral)
	if !isReceiverValue {
		return nil, fmt.Errorf("cannot access field %q on primitive value %s", s.FieldName, s.Receiver)
	}
	namedReceiverType, isNamedReceiverType := receiver.Type.(ast.NamedType)
	if !isNamedReceiverType {
		return nil, fmt.Errorf("type %q is not a valid value literal type", receiver.Type)
	}
	structFields, err := Fields(r.declarations, namedReceiverType)
	if err != nil {
		return nil, err
	}

	return r.reduceSelectToField(s, structFields, receiver)
}

func (r ReducingVisitor) selectWithReducedReceiver(s ast.Select) (ast.Expression, error) {
	reducedReceiver, err := r.Reduce(s.Receiver)
	return ast.Select{FieldName: s.FieldName, Receiver: reducedReceiver}, err
}

func Fields(declarations []ast.Declaration, typ ast.NamedType) ([]ast.Field, error) {
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)

		if isTypeDecl {
			structTypeLit, isStructLit := typeDecl.TypeLiteral.(ast.StructTypeLiteral)
			if isStructLit && typeDecl.TypeName == typ.TypeName {
				return structTypeLit.Fields, nil
			}
		}
	}
	return nil, fmt.Errorf("no struct type named %q found in declarations", typ.TypeName)
}

func (r ReducingVisitor) reduceSelectToField(s ast.Select, fields []ast.Field, receiver ast.ValueLiteral) (ast.Expression, error) {
	for i, field := range fields {
		if field.Name == s.FieldName {
			values := receiver.Values
			if len(values) <= i {
				return nil, fmt.Errorf("struct literal missing value at index %d", i)
			}
			return values[i], nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, receiver.Type)
}
