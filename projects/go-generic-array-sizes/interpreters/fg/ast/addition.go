package ast

func (a Add) Reduce(declarations []Declaration) (Expression, error) {
	left, isLeftIntLiteral := a.Left.(IntegerLiteral)
	if !isLeftIntLiteral {
		left, err := a.Left.Reduce(declarations)
		return Add{left, a.Right}, err
	}
	right, isRightIntLiteral := a.Right.(IntegerLiteral)
	if !isRightIntLiteral {
		right, err := a.Right.Reduce(declarations)
		return Add{a.Left, right}, err
	}
	return IntegerLiteral{left.IntValue + right.IntValue}, nil
}

func (a Add) Value() Value {
	return nil
}
