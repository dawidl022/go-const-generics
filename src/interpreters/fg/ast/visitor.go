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
	VisitArraySetMethodDeclaration(a ArraySetMethodDeclaration) error
}

type TypeVisitable interface {
	Accept(visitor TypeVisitor) (Type, error)
}

type TypeVisitor interface {
	VisitVariable(v Variable) (Type, error)
	VisitValueLiteral(v ValueLiteral) (Type, error)
	VisitField(s Select) (Type, error)
	VisitIntLiteral(i IntegerLiteral) (Type, error)
	VisitMethodCall(m MethodCall) (Type, error)
	VisitArrayIndex(a ArrayIndex) (Type, error)
	VisitAdd(a Add) (Type, error)
}

type MethodVisitable interface {
	AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification
}

type MethodVisitor interface {
	VisitTypeName(t TypeName) []MethodSpecification
	VisitIntegerLiteral(i IntegerLiteral) []MethodSpecification
}

type RefVisitable interface {
	AcceptRef(visitor RefVisitor) error
}

type RefVisitor interface {
	VisitStructTypeLiteral(s StructTypeLiteral) error
	VisitArrayTypeLiteral(a ArrayTypeLiteral) error
	VisitInterfaceTypeLiteral(i InterfaceTypeLiteral) error
}
