package ast

type Visitor interface {
	VisitProgram(p Program) error
	VisitTypeDeclaration(d TypeDeclaration) error
	VisitArrayTypeLiteral(a ArrayTypeLiteral) error
	VisitStructTypeLiteral(s StructTypeLiteral) error
	VisitTypeName(t TypeName) error
	VisitMethodSpecification(m MethodSpecification) error
	VisitInterfaceLiteral(i InterfaceTypeLiteral) error
}
