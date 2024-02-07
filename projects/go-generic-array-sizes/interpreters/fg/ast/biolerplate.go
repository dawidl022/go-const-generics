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

func (p Program) Accept(visitor Visitor) error {
	return visitor.VisitProgram(p)
}

func (t TypeDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitTypeDeclaration(t)
}

func (s StructTypeLiteral) Accept(visitor Visitor) error {
	return visitor.VisitStructTypeLiteral(s)
}

func (t TypeName) Accept(visitor Visitor) error {
	return visitor.VisitTypeName(t)
}

func (i InterfaceTypeLiteral) Accept(visitor Visitor) error {
	return visitor.VisitInterfaceLiteral(i)
}

func (a ArrayTypeLiteral) Accept(visitor Visitor) error {
	return visitor.VisitArrayTypeLiteral(a)
}

func (m MethodDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitMethodDeclaration(m)
}

func (m MethodSpecification) Accept(visitor Visitor) error {
	return visitor.VisitMethodSpecification(m)
}

func (a ArraySetMethodDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitArraySetMethodDeclaration(a)
}

func (v Variable) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitVariable(v)
}

func (i IntegerLiteral) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitIntLiteral(i)
}

func (m MethodCall) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitMethodCall(m)
}

func (v ValueLiteral) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitValueLiteral(v)
}

func (s Select) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitField(s)
}

func (a ArrayIndex) Accept(visitor TypeVisitor) (Type, error) {
	return visitor.VisitArrayIndex(a)
}

func (t TypeName) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitTypeName(t)
}

func (i IntegerLiteral) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitIntegerLiteral(i)
}

func (t TypeName) typeNode() {
}

func (i IntegerLiteral) typeNode() {
}

func (s StructTypeLiteral) AcceptRef(visitor RefVisitor) error {
	return visitor.VisitStructTypeLiteral(s)
}

func (i InterfaceTypeLiteral) AcceptRef(visitor RefVisitor) error {
	return visitor.VisitInterfaceTypeLiteral(i)
}

func (a ArrayTypeLiteral) AcceptRef(visitor RefVisitor) error {
	return visitor.VisitArrayTypeLiteral(a)
}
