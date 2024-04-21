package parsetree

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
)

type AntlrASTBuilder struct {
	*antlr.BaseParseTreeVisitor
	parseTree antlr.ParseTree
}

func NewAntlrASTBuilder(parseTree antlr.ParseTree) *AntlrASTBuilder {
	return &AntlrASTBuilder{parseTree: parseTree}
}

func (a *AntlrASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(a)
}

func (a *AntlrASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	declarations := []ast.Declaration{}
	for _, decl := range ctx.AllDeclaration() {
		declarations = append(declarations, a.Visit(decl).(ast.Declaration))
	}
	return ast.Program{
		Declarations: declarations,
		Expression:   a.Visit(ctx.MainFunc()).(ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitMainPackage(ctx *parser.MainPackageContext) interface{} {
	return nil
}

func (a *AntlrASTBuilder) VisitMainFunc(ctx *parser.MainFuncContext) interface{} {
	return a.Visit(ctx.Expression())
}

func (a *AntlrASTBuilder) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return a.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (a *AntlrASTBuilder) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	decl := ast.TypeDeclaration{
		TypeName:    a.Visit(ctx.TypeName()).(ast.TypeName),
		TypeLiteral: a.Visit(ctx.TypeLiteral()).(ast.TypeLiteral),
	}
	if ctx.TypeParameterConstraints() != nil {
		decl.TypeParameters = a.Visit(ctx.TypeParameterConstraints()).([]ast.TypeParameterConstraint)
	}
	return decl
}

func (a *AntlrASTBuilder) VisitTypeParameterConstraints(ctx *parser.TypeParameterConstraintsContext) interface{} {
	typeParams := []ast.TypeParameterConstraint{}
	for _, param := range ctx.AllTypeParameterConstraint() {
		typeParams = append(typeParams, a.Visit(param).(ast.TypeParameterConstraint))
	}
	return typeParams
}

func (a *AntlrASTBuilder) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	return ast.MethodDeclaration{
		MethodReceiver:      a.Visit(ctx.MethodReceiver()).(ast.MethodReceiver),
		MethodSpecification: a.Visit(ctx.MethodSpecification()).(ast.MethodSpecification),
		ReturnExpression:    a.Visit(ctx.Expression()).(ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitArraySetMethodDeclaration(ctx *parser.ArraySetMethodDeclarationContext) interface{} {
	return ast.ArraySetMethodDeclaration{
		MethodReceiver:        a.Visit(ctx.MethodReceiver()).(ast.MethodReceiver),
		MethodName:            a.Visit(ctx.MethodName()).(string),
		IndexParameter:        a.Visit(ctx.MethodParameter(0)).(ast.MethodParameter),
		ValueParameter:        a.Visit(ctx.MethodParameter(1)).(ast.MethodParameter),
		ReturnType:            a.Visit(ctx.Type_()).(ast.Type),
		IndexReceiverVariable: a.Visit(ctx.Variable(0)).(ast.Variable).Id,
		IndexVariable:         a.Visit(ctx.Variable(1)).(ast.Variable).Id,
		NewValueVariable:      a.Visit(ctx.Variable(2)).(ast.Variable).Id,
		ReturnVariable:        a.Visit(ctx.Variable(3)).(ast.Variable).Id,
	}
}

func (a *AntlrASTBuilder) VisitTypeLiteral(ctx *parser.TypeLiteralContext) interface{} {
	return a.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (a *AntlrASTBuilder) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	var fields []ast.Field
	for _, field := range ctx.AllField() {
		fields = append(fields, a.Visit(field).(ast.Field))
	}
	return ast.StructTypeLiteral{
		Fields: fields,
	}
}

func (a *AntlrASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	return ast.Field{
		Name: a.Visit(ctx.FieldName()).(string),
		Type: a.Visit(ctx.Type_()).(ast.Type),
	}
}

func (a *AntlrASTBuilder) VisitInterfaceLiteral(ctx *parser.InterfaceLiteralContext) interface{} {
	var specifications []ast.MethodSpecification
	for _, spec := range ctx.AllMethodSpecification() {
		specifications = append(specifications, a.Visit(spec).(ast.MethodSpecification))
	}
	return ast.InterfaceTypeLiteral{
		MethodSpecifications: specifications,
	}
}

func (a *AntlrASTBuilder) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	return ast.ArrayTypeLiteral{
		Length:      a.Visit(ctx.Type_(0)).(ast.Type),
		ElementType: a.Visit(ctx.Type_(1)).(ast.Type),
	}
}

func (a *AntlrASTBuilder) VisitMethodReceiver(ctx *parser.MethodReceiverContext) interface{} {
	receiver := ast.MethodReceiver{
		ParameterName: a.Visit(ctx.Variable()).(ast.Variable).Id,
		TypeName:      a.Visit(ctx.TypeName()).(ast.TypeName),
	}
	if ctx.TypeParameters() != nil {
		receiver.TypeParameters = a.Visit(ctx.TypeParameters()).([]ast.TypeParameter)
	}
	return receiver
}

func (a *AntlrASTBuilder) VisitMethodSpecification(ctx *parser.MethodSpecificationContext) interface{} {
	return ast.MethodSpecification{
		MethodName:      a.Visit(ctx.MethodName()).(string),
		MethodSignature: a.Visit(ctx.MethodSignature()).(ast.MethodSignature),
	}
}

func (a *AntlrASTBuilder) VisitMethodSignature(ctx *parser.MethodSignatureContext) interface{} {
	return ast.MethodSignature{
		MethodParameters: a.Visit(ctx.MethodParams()).([]ast.MethodParameter),
		ReturnType:       a.Visit(ctx.Type_()).(ast.Type),
	}
}

func (a *AntlrASTBuilder) VisitMethodParams(ctx *parser.MethodParamsContext) interface{} {
	var methodParams []ast.MethodParameter
	for _, param := range ctx.AllMethodParameter() {
		methodParams = append(methodParams, a.Visit(param).(ast.MethodParameter))
	}
	return methodParams
}

func (a *AntlrASTBuilder) VisitMethodParameter(ctx *parser.MethodParameterContext) interface{} {
	return ast.MethodParameter{
		ParameterName: a.Visit(ctx.Variable()).(ast.Variable).Id,
		Type:          a.Visit(ctx.Type_()).(ast.Type),
	}
}

func (a *AntlrASTBuilder) VisitNamedType(ctx *parser.NamedTypeContext) interface{} {
	namedType := ast.NamedType{
		TypeName: a.Visit(ctx.TypeName()).(ast.TypeName),
	}
	if ctx.TypeArguments() != nil {
		namedType.TypeArguments = a.Visit(ctx.TypeArguments()).([]ast.Type)
	}
	return namedType
}

func (a *AntlrASTBuilder) VisitIntType(ctx *parser.IntTypeContext) interface{} {
	return a.Visit(ctx.IntegerLiteral())
}

func (a *AntlrASTBuilder) VisitTypeParameters(ctx *parser.TypeParametersContext) interface{} {
	typeParams := []ast.TypeParameter{}
	for _, param := range ctx.AllTypeParameter() {
		typeParams = append(typeParams, a.Visit(param).(ast.TypeParameter))
	}
	return typeParams
}

func (a *AntlrASTBuilder) VisitTypeParameterConstraint(ctx *parser.TypeParameterConstraintContext) interface{} {
	return ast.TypeParameterConstraint{
		TypeParameter: a.Visit(ctx.TypeParameter()).(ast.TypeParameter),
		Bound:         a.Visit(ctx.Bound()).(ast.Bound),
	}
}

func (a *AntlrASTBuilder) VisitBound(ctx *parser.BoundContext) interface{} {
	if ctx.GetText() == "const" {
		return ast.ConstType{}
	}
	return a.Visit(ctx.Type_())
}

func (a *AntlrASTBuilder) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	typeArgs := []ast.Type{}
	for _, arg := range ctx.AllType_() {
		typeArgs = append(typeArgs, a.Visit(arg).(ast.Type))
	}
	return typeArgs
}

func (a *AntlrASTBuilder) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	return a.Visit(ctx.DecimalLiteral())
}

func (a *AntlrASTBuilder) VisitDecimalLiteral(ctx *parser.DecimalLiteralContext) interface{} {
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		panic(err)
	}
	return ast.IntegerLiteral{IntValue: int(val)}
}

func (a *AntlrASTBuilder) VisitValueLiteral(ctx *parser.ValueLiteralContext) interface{} {
	return ast.ValueLiteral{
		Type:   a.Visit(ctx.Type_()).(ast.Type),
		Values: a.Visit(ctx.ExpressionList()).([]ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitVar(ctx *parser.VarContext) interface{} {
	return a.Visit(ctx.Variable())
}

func (a *AntlrASTBuilder) VisitFieldSelect(ctx *parser.FieldSelectContext) interface{} {
	return ast.Select{
		Receiver:  a.Visit(ctx.Expression()).(ast.Expression),
		FieldName: ctx.FieldName().GetText(),
	}
}

func (a *AntlrASTBuilder) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	return a.Visit(ctx.IntegerLiteral())
}

func (a *AntlrASTBuilder) VisitArrIndex(ctx *parser.ArrIndexContext) interface{} {
	return ast.ArrayIndex{
		Receiver: a.Visit(ctx.Expression(0)).(ast.Expression),
		Index:    a.Visit(ctx.Expression(1)).(ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitMethodCall(ctx *parser.MethodCallContext) interface{} {
	return ast.MethodCall{
		Receiver:   a.Visit(ctx.Expression()).(ast.Expression),
		MethodName: a.Visit(ctx.MethodName()).(string),
		Arguments:  a.Visit(ctx.ExpressionList()).([]ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	var expressions []ast.Expression
	for _, expr := range ctx.AllExpression() {
		expressions = append(expressions, a.Visit(expr).(ast.Expression))
	}
	return expressions
}

func (a *AntlrASTBuilder) VisitVariable(ctx *parser.VariableContext) interface{} {
	return ast.Variable{Id: ctx.ID().GetText()}
}

func (a *AntlrASTBuilder) VisitTypeName(ctx *parser.TypeNameContext) interface{} {
	return ast.TypeName(ctx.ID().GetText())
}

func (a *AntlrASTBuilder) VisitMethodName(ctx *parser.MethodNameContext) interface{} {
	return ctx.ID().GetText()
}

func (a *AntlrASTBuilder) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	return ctx.ID().GetText()
}

func (a *AntlrASTBuilder) VisitTypeParameter(ctx *parser.TypeParameterContext) interface{} {
	return ast.TypeParameter(ctx.ID().GetText())
}

func (a *AntlrASTBuilder) VisitAdd(ctx *parser.AddContext) interface{} {
	return ast.Add{
		Left:  a.Visit(ctx.Expression(0)).(ast.Expression),
		Right: a.Visit(ctx.Expression(1)).(ast.Expression),
	}
}

func (a *AntlrASTBuilder) BuildAST() ast.Program {
	return a.Visit(a.parseTree).(ast.Program)
}
