package typecheck

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

func (t typeVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) (ast.Type, error) {
	return i, nil
}
