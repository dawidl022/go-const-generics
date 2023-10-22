package typecheck

import (
	"errors"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) (ast.Type, error) {
	return i, nil
}

func (t typeEnvTypeCheckingVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) error {
	if i.IntValue < 0 {
		return errors.New("negative integer cannot be used as type")
	}
	return nil
}
