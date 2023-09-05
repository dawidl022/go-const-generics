package ast

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

type Expression interface {
	expressionNode()
}

type IntegerLiteral struct {
	Value int
}

type Variable struct {
	Id string
}

type MethodCall struct {
	Expression Expression
	MethodName string
	Arguments  []Expression
}

type ValueLiteral struct {
	TypeName string
	Values   []Expression
}

type Select struct {
	Expression Expression
	FieldName  string
}

type ArrayIndex struct {
	Receiver Expression
	Index    Expression
}
