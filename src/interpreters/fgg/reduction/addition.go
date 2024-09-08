package reduction

import "github.com/dawidl022/go-const-generics/interpreters/fgg/ast"

func (r ReducingVisitor) VisitAdd(a ast.Add) (ast.Expression, error) {
	left, isLeftIntLiteral := a.Left.(ast.IntegerLiteral)
	if !isLeftIntLiteral {
		left, err := r.Reduce(a.Left)
		return ast.Add{left, a.Right}, err
	}
	right, isRightIntLiteral := a.Right.(ast.IntegerLiteral)
	if !isRightIntLiteral {
		right, err := r.Reduce(a.Right)
		return ast.Add{a.Left, right}, err
	}
	return ast.IntegerLiteral{left.IntValue + right.IntValue}, nil
}
