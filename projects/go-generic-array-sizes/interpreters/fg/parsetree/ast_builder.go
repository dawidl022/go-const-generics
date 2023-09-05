package parsetree

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
)

type ASTBuilder interface {
	BuildAST() ast.Program
}

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

func (a *AntlrASTBuilder) VisitMainPackage(_ *parser.MainPackageContext) interface{} {
	return nil
}

func (a *AntlrASTBuilder) VisitMainFunc(ctx *parser.MainFuncContext) interface{} {
	return a.Visit(ctx.Expression())
}

func (a *AntlrASTBuilder) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return a.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

func (a *AntlrASTBuilder) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	return ast.TypeDeclaration{
		TypeName:    a.Visit(ctx.TypeName()).(string),
		TypeLiteral: a.Visit(ctx.TypeLiteral()).(ast.TypeLiteral),
	}
}

func (a *AntlrASTBuilder) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	return ast.MethodDeclaration{
		MethodReceiver:      a.Visit(ctx.MethodReceiver()).(ast.MethodParameter),
		MethodSpecification: a.Visit(ctx.MethodSpecification()).(ast.MethodSpecification),
		ReturnExpression:    a.Visit(ctx.Expression()).(ast.Expression),
	}
}

func (a *AntlrASTBuilder) VisitArraySetMethodDeclaration(ctx *parser.ArraySetMethodDeclarationContext) interface{} {
	return ast.ArraySetMethodDeclaration{
		MethodReceiver:        a.Visit(ctx.MethodReceiver()).(ast.MethodParameter),
		MethodName:            a.Visit(ctx.MethodName()).(string),
		IndexParameter:        a.Visit(ctx.MethodParameter(0)).(ast.MethodParameter),
		ValueParameter:        a.Visit(ctx.MethodParameter(1)).(ast.MethodParameter),
		ReturnType:            a.Visit(ctx.TypeName()).(string),
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
		Name:     a.Visit(ctx.FieldName()).(string),
		TypeName: a.Visit(ctx.TypeName()).(string),
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

func (a *AntlrASTBuilder) VisitMethodSpecification(ctx *parser.MethodSpecificationContext) interface{} {
	return ast.MethodSpecification{
		MethodName:      a.Visit(ctx.MethodName()).(string),
		MethodSignature: a.Visit(ctx.MethodSignature()).(ast.MethodSignature),
	}
}

func (a *AntlrASTBuilder) VisitMethodSignature(ctx *parser.MethodSignatureContext) interface{} {
	signature := ast.MethodSignature{}
	signature.MethodParameters = a.Visit(ctx.MethodParams()).([]ast.MethodParameter)
	signature.ReturnTypeName = a.Visit(ctx.TypeName()).(string)
	return signature
}

func (a *AntlrASTBuilder) VisitMethodParams(ctx *parser.MethodParamsContext) interface{} {
	var methodParams []ast.MethodParameter
	for _, param := range ctx.AllMethodParameter() {
		methodParams = append(methodParams, a.Visit(param).(ast.MethodParameter))
	}
	return methodParams
}

func (a *AntlrASTBuilder) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	arrayLit := ast.ArrayTypeLiteral{}
	arrayLit.Length = a.Visit(ctx.IntegerLiteral()).(ast.IntegerLiteral).Value
	arrayLit.ElementTypeName = a.Visit(ctx.TypeName()).(string)
	return arrayLit
}

func (a *AntlrASTBuilder) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	return a.Visit(ctx.DecimalLiteral())
}

func (a *AntlrASTBuilder) VisitDecimalLiteral(ctx *parser.DecimalLiteralContext) interface{} {
	val, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return ast.IntegerLiteral{Value: val}
}

func (a *AntlrASTBuilder) VisitMethodReceiver(ctx *parser.MethodReceiverContext) interface{} {
	return a.Visit(ctx.MethodParameter())
}

func (a *AntlrASTBuilder) VisitMethodParameter(ctx *parser.MethodParameterContext) interface{} {
	receiver := ast.MethodParameter{}
	receiver.ParameterName = a.Visit(ctx.Variable()).(ast.Variable).Id
	receiver.TypeName = a.Visit(ctx.TypeName()).(string)
	return receiver
}

func (a *AntlrASTBuilder) VisitValueLiteral(ctx *parser.ValueLiteralContext) interface{} {
	valLiteral := ast.ValueLiteral{}
	valLiteral.TypeName = a.Visit(ctx.TypeName()).(string)
	valLiteral.Values = a.Visit(ctx.ExpressionList()).([]ast.Expression)
	return valLiteral
}

func (a *AntlrASTBuilder) VisitVar(ctx *parser.VarContext) interface{} {
	return a.Visit(ctx.Variable())
}

func (a *AntlrASTBuilder) VisitFieldSelect(ctx *parser.FieldSelectContext) interface{} {
	sel := ast.Select{}
	sel.FieldName = ctx.FieldName().GetText()
	sel.Expression = a.Visit(ctx.Expression()).(ast.Expression)
	return sel
}

func (a *AntlrASTBuilder) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	return a.Visit(ctx.IntegerLiteral())
}

func (a *AntlrASTBuilder) VisitArrIndex(ctx *parser.ArrIndexContext) interface{} {
	arrIndex := ast.ArrayIndex{}
	arrIndex.Receiver = a.Visit(ctx.Expression(0)).(ast.Expression)
	arrIndex.Index = a.Visit(ctx.Expression(1)).(ast.Expression)
	return arrIndex
}

func (a *AntlrASTBuilder) VisitMethodCall(ctx *parser.MethodCallContext) interface{} {
	methodCall := ast.MethodCall{
		MethodName: a.Visit(ctx.MethodName()).(string),
	}
	methodCall.Expression = a.Visit(ctx.Expression()).(ast.Expression)
	methodCall.Arguments = a.Visit(ctx.ExpressionList()).([]ast.Expression)

	return methodCall
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
	return ctx.ID().GetText()
}

func (a *AntlrASTBuilder) VisitMethodName(ctx *parser.MethodNameContext) interface{} {
	return ctx.ID().GetText()
}

func (a *AntlrASTBuilder) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	return ctx.ID().GetText()
}

func (a *AntlrASTBuilder) BuildAST() ast.Program {
	return a.Visit(a.parseTree).(ast.Program)
}
