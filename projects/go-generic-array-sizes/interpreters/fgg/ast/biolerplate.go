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

func (c ConstType) typeNode() {
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
	return visitor.VisitConstType(c)
}

func (c ConstType) AcceptMethodVisitor(visitor MethodVisitor) []MethodSpecification {
	return visitor.VisitConstType(c)
}

func (c ConstType) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapConstType(c)
}

func (t TypeParameter) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapTypeParameter(t)
}

func (n NamedType) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapNamedType(n)
}

func (i IntegerLiteral) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	//TODO implement me
	return i
}

func (s StructTypeLiteral) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapStructTypeLiteral(s)
}

func (i InterfaceTypeLiteral) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapInterfaceTypeLiteral(i)
}

func (a ArrayTypeLiteral) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapArrayTypeLiteral(a)
}

func (m MethodSpecification) AcceptEnvMapperVisitor(visitor EnvMapperVisitor) EnvVisitable {
	return visitor.VisitMapMethodSpecification(m)
}

func (p Program) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapProgram(p)
}

func (t TypeDeclaration) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapTypeDeclaration(t)
}

func (m MethodDeclaration) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapMethodDeclaration(m)
}

func (a ArraySetMethodDeclaration) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapArraySetMethodDeclaration(a)
}

func (t TypeParameterConstraint) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapTypeParameterConstraint(t)
}

func (s StructTypeLiteral) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapStructTypeLiteral(s)
}

func (i InterfaceTypeLiteral) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapInterfaceTypeLiteral(i)
}

func (a ArrayTypeLiteral) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapArrayTypeLiteral(a)
}

func (m MethodSpecification) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapMethodSpecification(m)
}

func (i IntegerLiteral) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapIntegerLiteral(i)
}

func (v Variable) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapVariable(v)
}

func (m MethodCall) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapMethodCall(m)
}

func (v ValueLiteral) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapValueLiteral(v)
}

func (s Select) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapSelect(s)
}

func (a ArrayIndex) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapArrayIndex(a)
}

func (p MethodParameter) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapMethodParameter(p)
}

func (c ConstType) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapConstType(c)
}

func (n NamedType) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapNamedType(n)
}

func (t TypeParameter) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapTypeParameter(t)
}

func (f Field) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapField(f)
}

func (m MethodSignature) AcceptMappingVisitor(visitor MappingVisitor) (MappingVisitable, error) {
	return visitor.VisitMapMethodSignature(m)
}

func (p Program) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapProgram(p)
}

func (t TypeDeclaration) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapTypeDeclaration(t)
}

func (m MethodDeclaration) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodDeclaration(m)
}

func (a ArraySetMethodDeclaration) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapArraySetMethodDeclaration(a)
}

func (t TypeParameterConstraint) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapTypeParameterConstraint(t)
}

func (s StructTypeLiteral) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapStructTypeLiteral(s)
}

func (i InterfaceTypeLiteral) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapInterfaceTypeLiteral(i)
}

func (a ArrayTypeLiteral) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapArrayTypeLiteral(a)
}

func (m MethodSpecification) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodSpecification(m)
}

func (i IntegerLiteral) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapIntegerLiteral(i)
}

func (v Variable) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapVariable(v)
}

func (m MethodCall) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodCall(m)
}

func (v ValueLiteral) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapValueLiteral(v)
}

func (s Select) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapSelect(s)
}

func (a ArrayIndex) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapArrayIndex(a)
}

func (p MethodParameter) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodParameter(p)
}

func (c ConstType) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapConstType(c)
}

func (n NamedType) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapNamedType(n)
}

func (t TypeParameter) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapTypeParameter(t)
}

func (f Field) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapField(f)
}

func (m MethodSignature) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodSignature(m)
}

func (m MethodReceiver) AcceptMapVisitor(visitor MapVisitor) MapVisitable {
	return visitor.MapMethodReceiver(m)
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
