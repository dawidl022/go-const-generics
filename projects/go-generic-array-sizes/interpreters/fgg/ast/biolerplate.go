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

func (m MethodDeclaration) GetMethodReceiver() MethodReceiver {
	return m.MethodReceiver
}

func (m MethodDeclaration) GetMethodName() string {
	return m.MethodSpecification.MethodName
}

func (a ArraySetMethodDeclaration) GetMethodReceiver() MethodReceiver {
	return a.MethodReceiver
}

func (a ArraySetMethodDeclaration) GetMethodName() string {
	return a.MethodName
}

func (i IntegerLiteral) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitIntegerLiteral(i)
}

func (v Variable) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitVariable(v)
}

func (m MethodCall) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitMethodCall(m)
}

func (v ValueLiteral) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitValueLiteral(v)
}

func (s Select) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitSelect(s)
}

func (a ArrayIndex) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitArrayIndex(a)
}

func (p Program) Accept(visitor Visitor) error {
	return visitor.VisitProgram(p)
}

func (i IntegerLiteral) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitIntegerLiteral(i)
}

func (v Variable) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitVariable(v)
}

func (m MethodCall) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitMethodCall(m)
}

func (v ValueLiteral) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitValueLiteral(v)
}

func (s Select) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitSelect(s)
}

func (a ArrayIndex) AcceptTypeVisitor(visitor TypeVisitor) (Type, error) {
	return visitor.VisitArrayIndex(a)
}

func (t TypeDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitTypeDeclaration(t)
}

func (m MethodDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitMethodDeclaration(m)
}

func (a ArraySetMethodDeclaration) Accept(visitor Visitor) error {
	return visitor.VisitArraySetMethodDeclaration(a)
}

func (s StructTypeLiteral) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitStructTypeLiteral(s)
}

func (i InterfaceTypeLiteral) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitInterfaceTypeLiteral(i)
}

func (a ArrayTypeLiteral) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitArrayTypeLiteral(a)
}

func (i IntegerLiteral) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitIntegerLiteral(i)
}

func (n NamedType) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitNamedType(n)
}

func (t TypeParameter) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitTypeParameter(t)
}

func (m MethodSpecification) AcceptEnvVisitor(visitor EnvVisitor) error {
	return visitor.VisitMethodSpecification(m)
}

func (i IntegerLiteral) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitIntegerLiteral(i)
}

func (n NamedType) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitNamedType(n)
}

func (t TypeParameter) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitTypeParameter(t)
}

func (c ConstType) AcceptEnvVisitor(visitor EnvVisitor) error {
	//TODO implement me
	panic("implement me")
}

func (c ConstType) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	//TODO implement me
	panic("implement me")
}

func (c ConstType) Equal(other Type) bool {
	_, isOtherConstType := other.(ConstType)
	return isOtherConstType
}

func (c ConstType) typeNode() {
	//TODO implement me
	panic("implement me")
}

func (c ConstType) AcceptEnvTypeVisitor(visitor EnvTypeVisitor) (Type, error) {
	//TODO implement me
	panic("implement me")
}

func (t TypeParameter) AcceptEnvTypeVisitor(visitor EnvTypeVisitor) (Type, error) {
	//TODO implement me
	panic("implement me")
}

func (n NamedType) AcceptEnvTypeVisitor(visitor EnvTypeVisitor) (Type, error) {
	return visitor.VisitEnvNamedType(n)
}

func (i IntegerLiteral) AcceptEnvTypeVisitor(visitor EnvTypeVisitor) (Type, error) {
	//TODO implement me
	return i, nil
}

func (s StructTypeLiteral) AcceptEnvTypeVisitor(visitor EnvTypeLiteralVisitor) (TypeLiteral, error) {
	//TODO implement me
	return s, nil
}

func (i InterfaceTypeLiteral) AcceptEnvTypeVisitor(visitor EnvTypeLiteralVisitor) (TypeLiteral, error) {
	//TODO implement me
	return i, nil
}

func (a ArrayTypeLiteral) AcceptEnvTypeVisitor(visitor EnvTypeLiteralVisitor) (TypeLiteral, error) {
	return visitor.VisitEnvArrayTypeLiteral(a)
}
