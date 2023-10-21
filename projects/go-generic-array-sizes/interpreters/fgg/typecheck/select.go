package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/auxiliary"
)

func (t typeVisitor) VisitSelect(s ast.Select) (ast.Type, error) {
	receiver, err := t.typeOf(s.Receiver)
	if err != nil {
		return nil, err
	}
	receiverNamedType, isReceiverNamedType := receiver.(ast.NamedType)
	if !isReceiverNamedType {
		// TODO error message may need adjusting when trying to access fields on type parameters
		return nil, fmt.Errorf("cannot access field %q on primitive value of type %q", s.FieldName, receiver)
	}
	fields, err := auxiliary.Fields(t.declarations, receiverNamedType)
	if err != nil {
		return nil, fmt.Errorf("cannot access field %q on type %q: %w", s.FieldName, receiver, err)
	}
	for _, f := range fields {
		if f.Name == s.FieldName {
			return f.Type, nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, receiver)
}
