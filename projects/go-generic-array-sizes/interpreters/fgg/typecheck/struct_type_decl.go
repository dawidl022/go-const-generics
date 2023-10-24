package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
)

func (t typeEnvTypeCheckingVisitor) VisitStructTypeLiteral(s ast.StructTypeLiteral) error {
	if err := checkDistinctFieldNames(s); err != nil {
		return err
	}
	for _, field := range s.Fields {
		if err := t.typeCheck(field.Type); err != nil {
			return fmt.Errorf("field %q %w", field.Name, err)
		}
		if t.isConst(field.Type) {
			return fmt.Errorf("cannot use const type %q as field type", field.Type)
		}
	}
	return nil
}

func checkDistinctFieldNames(s ast.StructTypeLiteral) error {
	fieldNames := []name{}
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, name(field.Name))
	}
	if err := auxiliary.Distinct(fieldNames); err != nil {
		return fmt.Errorf("field name %w", err)
	}
	return nil
}
