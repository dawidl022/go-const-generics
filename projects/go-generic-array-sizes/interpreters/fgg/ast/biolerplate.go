package ast

func (t TypeDeclaration) declarationNode() {
}

func (t TypeParameter) typeNode() {
}

func (t TypeParameter) boundNode() {
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

func (n NamedType) boundNode() {
}

func (i IntegerLiteral) typeNode() {
}

func (i IntegerLiteral) boundNode() {
}

func (c ConstType) boundNode() {
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
