package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func (t typeEnvTypeCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if err := t.typeCheck(a.Length); err != nil {
		return fmt.Errorf("length %w", err)
	}
	if !t.isConst(a.Length) {
		return fmt.Errorf("non-const type %q used as length", a.Length)
	}
	if err := t.typeCheck(a.ElementType); err != nil {
		return fmt.Errorf("element %w", err)
	}
	if t.isConst(a.ElementType) {
		return fmt.Errorf("const type %q used as element type", a.ElementType)
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) isConst(typ ast.Type) bool {
	return t.CheckIsSubtypeOf(typ, ast.ConstType{}) == nil
}
