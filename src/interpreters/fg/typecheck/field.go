package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeVisitor) VisitField(s ast.Select) (ast.Type, error) {
	receiver, err := t.typeOf(s.Receiver)
	if err != nil {
		return nil, err
	}
	valueTypeName, isValueTypeName := receiver.(ast.TypeName)
	if !isValueTypeName {
		return nil, fmt.Errorf("cannot access field %q on primitive value of type %q", s.FieldName, receiver)
	}

	fields, err := ast.Fields(t.declarations, valueTypeName)
	if err != nil {
		return nil, fmt.Errorf("cannot access field %q on type %q: %w", s.FieldName, receiver, err)
	}
	for _, f := range fields {
		if f.Name == s.FieldName {
			return f.TypeName, nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, receiver)
}
