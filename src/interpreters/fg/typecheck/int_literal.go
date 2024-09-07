package typecheck

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"

func (t typeVisitor) VisitIntLiteral(i ast.IntegerLiteral) (ast.Type, error) {
	return i, nil
}
