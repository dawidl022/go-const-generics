package ast

type Visitable interface {
	Accept(visitor Visitor) error
}

type Visitor interface {
	VisitProgram(p Program) error
	VisitTypeDeclaration(t TypeDeclaration) error
	VisitMethodDeclaration(m MethodDeclaration) error
	VisitArraySetMethodDeclaration(a ArraySetMethodDeclaration) error
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
	VisitMethodCall(m MethodCall) (Type, error)
}

type EnvVisitable interface {
	AcceptEnvVisitor(visitor EnvVisitor) error
	AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable
}

type EnvVisitor interface {
	VisitArrayTypeLiteral(a ArrayTypeLiteral) error
	VisitNamedType(n NamedType) error
	VisitInterfaceTypeLiteral(i InterfaceTypeLiteral) error
	VisitMethodSpecification(m MethodSpecification) error
	VisitStructTypeLiteral(s StructTypeLiteral) error
	VisitIntegerLiteral(i IntegerLiteral) error
	VisitTypeParameter(t TypeParameter) error
	VisitConstType(c ConstType) error
}

type EnvMapperVisitor interface {
	VisitMapNamedType(n NamedType) EnvVisitable
	VisitMapConstType(c ConstType) EnvVisitable
	VisitMapArrayTypeLiteral(a ArrayTypeLiteral) EnvVisitable
	VisitMapInterfaceTypeLiteral(i InterfaceTypeLiteral) EnvVisitable
	VisitMapMethodSpecification(m MethodSpecification) EnvVisitable
	VisitMapTypeParameter(t TypeParameter) EnvVisitable
}

type MethodVisitable interface {
	AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification
}

type MethodVisitor interface {
	VisitIntegerLiteral(i IntegerLiteral) []MethodSpecification
	VisitNamedType(n NamedType) []MethodSpecification
	VisitTypeParameter(t TypeParameter) []MethodSpecification
	VisitConstType(c ConstType) []MethodSpecification
}

type MappingVisitable interface {
	AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error)
}

type MappingVisitor interface {
	VisitMapProgram(p Program) (MappingVisitable, error)
	VisitMapTypeDeclaration(t TypeDeclaration) (MappingVisitable, error)
	VisitMapMethodDeclaration(m MethodDeclaration) (MappingVisitable, error)
	VisitMapArraySetMethodDeclaration(a ArraySetMethodDeclaration) (MappingVisitable, error)
	VisitMapTypeParameterConstraint(t TypeParameterConstraint) (MappingVisitable, error)
	VisitMapStructTypeLiteral(s StructTypeLiteral) (MappingVisitable, error)
	VisitMapInterfaceTypeLiteral(i InterfaceTypeLiteral) (MappingVisitable, error)
	VisitMapArrayTypeLiteral(a ArrayTypeLiteral) (MappingVisitable, error)
	VisitMapMethodSpecification(m MethodSpecification) (MappingVisitable, error)
	VisitMapIntegerLiteral(i IntegerLiteral) (MappingVisitable, error)
	VisitMapVariable(v Variable) (MappingVisitable, error)
	VisitMapMethodCall(m MethodCall) (MappingVisitable, error)
	VisitMapValueLiteral(v ValueLiteral) (MappingVisitable, error)
	VisitMapSelect(s Select) (MappingVisitable, error)
	VisitMapArrayIndex(a ArrayIndex) (MappingVisitable, error)
	VisitMapMethodParameter(p MethodParameter) (MappingVisitable, error)
	VisitMapConstType(c ConstType) (MappingVisitable, error)
	VisitMapNamedType(n NamedType) (MappingVisitable, error)
	VisitMapTypeParameter(t TypeParameter) (MappingVisitable, error)
	VisitMapField(f Field) (MappingVisitable, error)
	VisitMapMethodSignature(m MethodSignature) (MappingVisitable, error)
}
