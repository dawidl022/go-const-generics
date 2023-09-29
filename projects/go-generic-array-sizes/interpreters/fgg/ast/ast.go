package ast

import (
	"fmt"
)

type Program struct {
	Declarations []Declaration
	Expression   Expression
}

type Declaration interface {
	declarationNode()
}

type TypeDeclaration struct {
	TypeName       TypeName
	TypeParameters []TypeParameterConstraint
	TypeLiteral    TypeLiteral
}

type Type interface {
	Bound
	typeNode()
	fmt.Stringer
}

type TypeParameterConstraint struct {
	TypeParameter TypeParameter
	Bound         Bound
}

type TypeParameter string

type Bound interface {
	boundNode()
}

type ConstType struct{}

type TypeLiteral interface {
	typeLiteralNode()
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

type NamedType struct {
	TypeName      TypeName
	TypeArguments []Type
}

type Expression interface {
	IsValue() bool
	expressionNode()
	fmt.Stringer
	ExpressionVisitable
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
