package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func (t typeVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Type, error) {
	if t.isStructTypeName(v.Type.TypeName) {
		if err := t.typeCheck(v.Type); err != nil {
			return nil, err
		}
		return v.Type, t.typeCheckStructLiteral(v)
	}
	if t.isArrayTypeName(v.Type.TypeName) {
		if err := t.typeCheck(v.Type); err != nil {
			return nil, err
		}
		return v.Type, t.typeCheckArrayLiteral(v, v.Type)
	}
	return nil, fmt.Errorf("undeclared value literal type name: %q", v.Type)
}
