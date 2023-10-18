package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Type, error) {
	namedType, hasNamedType := v.Type.(ast.NamedType)
	if !hasNamedType {
		panic("untested branch")
	}
	if t.isStructTypeName(namedType.TypeName) {
		return v.Type, t.typeCheckStructLiteral(v)
	}
	if t.isArrayTypeName(namedType.TypeName) {
		return v.Type, t.typeCheckArrayLiteral(v)
	}
	return nil, fmt.Errorf("undeclared value literal type name: %q", v.Type)
}

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	//	TODO
	return nil
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	// TODO
	return nil
}