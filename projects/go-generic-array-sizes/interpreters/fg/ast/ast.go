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
	declarationNode()
}

type TypeDeclaration struct {
	TypeName    TypeName
	TypeLiteral TypeLiteral
}

type TypeLiteral interface {
	Visitable
	typeLiteralNode()
}

type StructTypeLiteral struct {
	Fields []Field
}

type Field struct {
	Name     string
	TypeName TypeName
}

type Type interface {
	MethodVisitable
	typeNode()
	fmt.Stringer
}

type TypeName string

type InterfaceTypeLiteral struct {
	MethodSpecifications []MethodSpecification
}

type MethodSpecification struct {
	MethodName      string
	MethodSignature MethodSignature
}

type MethodSignature struct {
	MethodParameters []MethodParameter
	ReturnTypeName   TypeName
}

type MethodParameter struct {
	ParameterName string
	TypeName      TypeName
}

type ArrayTypeLiteral struct {
	Length          int
	ElementTypeName TypeName
}

type MethodDeclaration struct {
	MethodReceiver      MethodParameter
	MethodSpecification MethodSpecification
	ReturnExpression    Expression
}

type ArraySetMethodDeclaration struct {
	MethodReceiver        MethodParameter
	MethodName            string
	IndexParameter        MethodParameter
	ValueParameter        MethodParameter
	ReturnType            TypeName
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
			ReturnTypeName:   a.ReturnType,
		},
	}
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
	TypeName TypeName
	Values   []Expression
}

type Select struct {
	Receiver  Expression
	FieldName string
}

type ArrayIndex struct {
	Receiver Expression
	Index    Expression
}
