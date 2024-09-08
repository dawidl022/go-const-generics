package typecheck

import "github.com/dawidl022/go-const-generics/interpreters/fg/ast"

func (t typeVisitor) VisitIntLiteral(i ast.IntegerLiteral) (ast.Type, error) {
	return i, nil
}
