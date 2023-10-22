package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/auxiliary"
)

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	namedValueType, isNamedValueType := v.Type.(ast.NamedType)
	if !isNamedValueType {
		panic("untested branch")
	}
	fields, err := auxiliary.Fields(t.declarations, namedValueType)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	if len(v.Values) != len(fields) {
		return fmt.Errorf("struct literal of type %q requires %d values, but got %d",
			namedValueType.TypeName, len(fields), len(v.Values))
	}
	for i, f := range fields {
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			return err
		}
		namedFieldType, isNamedFieldType := f.Type.(ast.NamedType)
		if !isNamedFieldType {
			panic("untested branch")
		}
		err = t.checkIsSubtypeOf(fieldType, namedFieldType)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w",
				v.Values[i], f.Name, namedValueType.TypeName, err)
		}
	}
	return nil
}
