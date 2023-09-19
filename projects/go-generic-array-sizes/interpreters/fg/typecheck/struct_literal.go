package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	fields, err := ast.Fields(t.declarations, v.TypeName)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	if len(v.Values) != len(fields) {
		return fmt.Errorf("struct literal of type %q requires %d values, but got %d",
			v.TypeName, len(fields), len(v.Values))
	}
	for i, f := range fields {
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			return err
		}
		err = t.checkIsSubtypeOf(fieldType, f.TypeName)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w", v.Values[i], f.Name, v.TypeName, err)
		}
	}
	return nil
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	// TODO
	return nil
}
