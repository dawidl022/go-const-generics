package ast

type ExpressionVisitable interface {
	Accept(visitor ExpressionVisitor) (Expression, error)
}

type ExpressionVisitor interface {
	VisitIntegerLiteral(i IntegerLiteral) (Expression, error)
	VisitVariable(v Variable) (Expression, error)
	VisitMethodCall(m MethodCall) (Expression, error)
	VisitValueLiteral(v ValueLiteral) (Expression, error)
	VisitSelect(s Select) (Expression, error)
	VisitArrayIndex(a ArrayIndex) (Expression, error)
}

func (i IntegerLiteral) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitIntegerLiteral(i)
}

func (v Variable) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitVariable(v)
}

func (m MethodCall) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitMethodCall(m)
}

func (v ValueLiteral) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitValueLiteral(v)
}

func (s Select) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitSelect(s)
}

func (a ArrayIndex) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitArrayIndex(a)
}
