package ast

type Visitable interface {
	Accept(visitor Visitor) error
}

type Visitor interface {
	VisitProgram(p Program) error
	VisitTypeDeclaration(d TypeDeclaration) error
	VisitArrayTypeLiteral(a ArrayTypeLiteral) error
	VisitStructTypeLiteral(s StructTypeLiteral) error
	VisitTypeName(t TypeName) error
	VisitMethodSpecification(m MethodSpecification) error
	VisitInterfaceLiteral(i InterfaceTypeLiteral) error
	VisitMethodDeclaration(m MethodDeclaration) error
}

type TypeVisitable interface {
	Accept(visitor TypeVisitor) (Type, error)
}

type TypeVisitor interface {
	VisitVariable(v Variable) (Type, error)
	VisitValueLiteral(v ValueLiteral) (Type, error)
	VisitField(s Select) (Type, error)
	VisitIntLiteral(i IntegerLiteral) (Type, error)
}

type MethodVisitable interface {
	AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification
}

type MethodVisitor interface {
	VisitTypeName(t TypeName) []MethodSpecification
	VisitIntegerLiteral(i IntegerLiteral) []MethodSpecification
}
