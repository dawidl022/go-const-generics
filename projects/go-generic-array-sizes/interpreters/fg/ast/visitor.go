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
	Accept(visitor TypeVisitor) (TypeName, error)
}

type TypeVisitor interface {
	VisitVariable(v Variable) (TypeName, error)
}
