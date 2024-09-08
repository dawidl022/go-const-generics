package typecheck

import "github.com/dawidl022/go-const-generics/interpreters/fg/ast"

func (t typeVisitor) VisitAdd(a ast.Add) (ast.Type, error) {
	leftTyp, err := t.typeOf(a.Left)
	if err != nil {
		return nil, err
	}
	rightTyp, err := t.typeOf(a.Right)
	if err != nil {
		return nil, err
	}
	leftInt, isLeftIntLiteral := leftTyp.(ast.IntegerLiteral)
	if !isLeftIntLiteral {
		return intTypeName, nil
	}
	rightInt, isRightIntLiteral := rightTyp.(ast.IntegerLiteral)
	if !isRightIntLiteral {
		return intTypeName, nil
	}
	return ast.IntegerLiteral{
		IntValue: leftInt.IntValue + rightInt.IntValue,
	}, nil
}
