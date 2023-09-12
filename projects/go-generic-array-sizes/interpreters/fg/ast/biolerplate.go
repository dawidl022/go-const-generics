package ast

func (t TypeDeclaration) declarationNode() {
}

func (s StructTypeLiteral) typeLiteralNode() {
}

func (i InterfaceTypeLiteral) typeLiteralNode() {
}

func (a ArrayTypeLiteral) typeLiteralNode() {
}

func (m MethodDeclaration) declarationNode() {
}

func (a ArraySetMethodDeclaration) declarationNode() {
}

func (i IntegerLiteral) expressionNode() {
}

func (v Variable) expressionNode() {
}

func (m MethodCall) expressionNode() {
}

func (v ValueLiteral) expressionNode() {
}

func (s Select) expressionNode() {
}

func (a ArrayIndex) expressionNode() {
}

func (m MethodDeclaration) GetMethodReceiver() MethodParameter {
	return m.MethodReceiver
}

func (m MethodDeclaration) GetMethodName() string {
	return m.MethodSpecification.MethodName
}

func (a ArraySetMethodDeclaration) GetMethodReceiver() MethodParameter {
	return a.MethodReceiver
}

func (a ArraySetMethodDeclaration) GetMethodName() string {
	return a.MethodName
}
