package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/auxiliary"
)

func (r ReducingVisitor) VisitSelect(s ast.Select) (ast.Expression, error) {
	if !s.Receiver.IsValue() {
		return r.selectWithReducedReceiver(s)
	}
	receiver, isReceiverValue := s.Receiver.(ast.ValueLiteral)
	if !isReceiverValue {
		return nil, fmt.Errorf("cannot access field %q on primitive value %s", s.FieldName, s.Receiver)
	}
	structFields, err := auxiliary.Fields(r.declarations, receiver.Type)
	if err != nil {
		return nil, err
	}

	return r.reduceSelectToField(s, structFields, receiver)
}

func (r ReducingVisitor) selectWithReducedReceiver(s ast.Select) (ast.Expression, error) {
	reducedReceiver, err := r.Reduce(s.Receiver)
	return ast.Select{FieldName: s.FieldName, Receiver: reducedReceiver}, err
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
