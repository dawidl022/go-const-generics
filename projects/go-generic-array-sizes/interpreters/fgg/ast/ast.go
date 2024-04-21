package ast

import (
	"fmt"
)

type Program struct {
	Declarations []Declaration
	Expression   Expression
}

type Declaration interface {
	Visitable
	MappingVisitable
	MapVisitable
	declarationNode()
	fmt.Stringer
}

type TypeDeclaration struct {
	TypeName       TypeName
	TypeParameters []TypeParameterConstraint
	TypeLiteral    TypeLiteral
}

type Type interface {
	EnvVisitable
	MethodVisitable
	MappingVisitable
	MapVisitable
	TypeRefVisitable
	Equal(other Type) bool
	typeNode()
	fmt.Stringer
}

type TypeParameterConstraint struct {
	TypeParameter TypeParameter
	Bound         Bound
}

type TypeParameter string

type Bound = Type

type ConstType struct{}

type TypeLiteral interface {
	EnvVisitable
	MappingVisitable
	MapVisitable
	RefVisitable
	typeLiteralNode()
	fmt.Stringer
}

type StructTypeLiteral struct {
	Fields []Field
}

type Field struct {
	Name string
	Type Type
}

type InterfaceTypeLiteral struct {
	MethodSpecifications []MethodSpecification
}

type MethodSpecification struct {
	MethodName      string
	MethodSignature MethodSignature
}

type MethodSignature struct {
	MethodParameters []MethodParameter
	ReturnType       Type
}

type MethodParameter struct {
	ParameterName string
	Type          Type
}

type ArrayTypeLiteral struct {
	Length      Type
	ElementType Type
}

type MethodDeclaration struct {
	MethodReceiver      MethodReceiver
	MethodSpecification MethodSpecification
	ReturnExpression    Expression
}

type MethodReceiver struct {
	ParameterName  string
	TypeName       TypeName
	TypeParameters []TypeParameter
}

type TypeName string

type ArraySetMethodDeclaration struct {
	MethodReceiver        MethodReceiver
	MethodName            string
	IndexParameter        MethodParameter
	ValueParameter        MethodParameter
	ReturnType            Type
	IndexReceiverVariable string
	IndexVariable         string
	NewValueVariable      string
	ReturnVariable        string
}

func (a ArraySetMethodDeclaration) MethodSpecification() MethodSpecification {
	return MethodSpecification{
		MethodName: a.MethodName,
		MethodSignature: MethodSignature{
			MethodParameters: []MethodParameter{a.IndexParameter, a.ValueParameter},
			ReturnType:       a.ReturnType,
		},
	}
}

const IntTypeName TypeName = "int"

type NamedType struct {
	TypeName      TypeName
	TypeArguments []Type
}

type Expression interface {
	IsValue() bool
	expressionNode()
	fmt.Stringer
	ExpressionVisitable
	TypeVisitable
	MappingVisitable
	MapVisitable
}

type IntegerLiteral struct {
	IntValue int
}

type Variable struct {
	Id string
}

type MethodCall struct {
	Receiver   Expression
	MethodName string
	Arguments  []Expression
}

type ValueLiteral struct {
	Type   Type
	Values []Expression
}

type Select struct {
	Receiver  Expression
	FieldName string
}

type ArrayIndex struct {
	Receiver Expression
	Index    Expression
}

type Add struct {
	Left  Expression
	Right Expression
}

func (a Add) IsValue() bool {
	return false
}

func (a Add) expressionNode() {
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.Left, a.Right)
}

func (a Add) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitAdd(a)
}

func (a Add) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitAdd(a)
}

func (a Add) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	//TODO implement me
	panic("implement me")
}

func (a Add) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	//TODO implement me
	panic("implement me")
}
