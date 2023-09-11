package ast

import "fmt"

type Program struct {
	Declarations []Declaration
	Expression   Expression
}

type Declaration interface {
	declarationNode()
}

type TypeDeclaration struct {
	TypeName    string
	TypeLiteral TypeLiteral
}

type TypeLiteral interface {
	typeLiteralNode()
}

type StructTypeLiteral struct {
	Fields []Field
}

type Field struct {
	Name     string
	TypeName string
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
	ReturnTypeName   string
}

type MethodParameter struct {
	ParameterName string
	TypeName      string
}

type ArrayTypeLiteral struct {
	Length          int
	ElementTypeName string
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
	ReturnType            string
	IndexReceiverVariable string
	IndexVariable         string
	NewValueVariable      string
	ReturnVariable        string
}

type IntegerLiteral struct {
	IntValue int
}

func (i IntegerLiteral) bind(variables map[string]Expression) (Expression, error) {
	return i, nil
}

type Variable struct {
	Id string
}

func (v Variable) bind(variables map[string]Expression) (Expression, error) {
	if val, isBound := variables[v.Id]; isBound {
		return val, nil
	}
	return nil, fmt.Errorf("unbound variable %q", v.Id)
}

type MethodCall struct {
	Receiver   Expression
	MethodName string
	Arguments  []Expression
}

func (m MethodCall) bind(variables map[string]Expression) (Expression, error) {
	boundReceiver, err := m.Receiver.bind(variables)
	if err != nil {
		return nil, err
	}
	boundArgs := []Expression{}
	for _, arg := range m.Arguments {
		boundArg, err := arg.bind(variables)
		if err != nil {
			return nil, err
		}
		boundArgs = append(boundArgs, boundArg)
	}
	return MethodCall{
		Receiver:   boundReceiver,
		MethodName: m.MethodName,
		Arguments:  boundArgs,
	}, nil
}

type ValueLiteral struct {
	TypeName string
	Values   []Expression
}

func (v ValueLiteral) bind(variables map[string]Expression) (Expression, error) {
	boundValues := []Expression{}
	for _, val := range v.Values {
		boundVal, err := val.bind(variables)
		if err != nil {
			return nil, err
		}
		boundValues = append(boundValues, boundVal)
	}
	return ValueLiteral{TypeName: v.TypeName, Values: boundValues}, nil
}

type Select struct {
	Receiver  Expression
	FieldName string
}

func (s Select) bind(variables map[string]Expression) (Expression, error) {
	boundExpression, err := s.Receiver.bind(variables)
	return Select{FieldName: s.FieldName, Receiver: boundExpression}, err
}

func (s Select) String() string {
	return fmt.Sprintf("%s.%s", s.Receiver, s.FieldName)
}

type ArrayIndex struct {
	Receiver Expression
	Index    Expression
}

func (a ArrayIndex) bind(variables map[string]Expression) (Expression, error) {
	boundReceiver, err := a.Receiver.bind(variables)
	if err != nil {
		return nil, err
	}
	boundIndex, err := a.Index.bind(variables)
	return ArrayIndex{Receiver: boundReceiver, Index: boundIndex}, err
}

func (a ArrayIndex) String() string {
	return fmt.Sprintf("%s[%s]", a.Receiver, a.Index)
}
