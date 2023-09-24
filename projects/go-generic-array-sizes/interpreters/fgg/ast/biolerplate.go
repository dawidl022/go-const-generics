package ast

func (t TypeDeclaration) declarationNode() {
}

func (t TypeParameter) typeNode() {
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

func (n NamedType) typeNode() {
}

func (i IntegerLiteral) typeNode() {
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
