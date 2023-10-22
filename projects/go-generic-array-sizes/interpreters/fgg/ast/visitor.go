package ast

type Visitable interface {
	Accept(visitor Visitor) error
}

type Visitor interface {
	VisitProgram(p Program) error
	VisitTypeDeclaration(t TypeDeclaration) error
	VisitMethodDeclaration(m MethodDeclaration) error
}

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

type TypeVisitable interface {
	AcceptTypeVisitor(visitor TypeVisitor) (Type, error)
}

type TypeVisitor interface {
	VisitVariable(v Variable) (Type, error)
	VisitValueLiteral(v ValueLiteral) (Type, error)
	VisitIntegerLiteral(i IntegerLiteral) (Type, error)
	VisitSelect(s Select) (Type, error)
	VisitArrayIndex(a ArrayIndex) (Type, error)
}

type EnvVisitable interface {
	AcceptEnvVisitor(visitor EnvVisitor) error
}

type EnvVisitor interface {
	AcceptArrayTypeLiteral(a ArrayTypeLiteral) error
	VisitNamedType(n NamedType) error
	VisitInterfaceTypeLiteral(i InterfaceTypeLiteral) error
	VisitMethodSpecification(m MethodSpecification) error
	VisitStructTypeLiteral(s StructTypeLiteral) error
}

type MethodVisitable interface {
	AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification
}

type MethodVisitor interface {
	VisitIntegerLiteral(i IntegerLiteral) []MethodSpecification
	VisitNamedType(n NamedType) []MethodSpecification
	TypeParameter(t TypeParameter) []MethodSpecification
}
