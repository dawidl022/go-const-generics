package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
)

func (t typeVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Type, error) {
	if t.isStructTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckStructLiteral(v)
	}
	if t.isArrayTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckArrayLiteral(v)
	}
	return nil, fmt.Errorf("undeclared value literal type name: %q", v.TypeName)
}
