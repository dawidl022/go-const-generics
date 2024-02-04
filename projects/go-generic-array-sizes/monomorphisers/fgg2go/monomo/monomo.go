package monomo

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func Monomorphise(p ast.Program) ast.Program {
	return newVisitor().monomorphise(p).(ast.Program)
}

type visitor struct {
	// map from type name to instantiations with generic type parameters
	queue              map[ast.TypeName][][]ast.IntegerLiteral
	seenInstantiations map[ast.TypeName][][]ast.IntegerLiteral

	typeEnv           map[ast.TypeParameter]ast.IntegerLiteral
	numericalTypeArgs []ast.IntegerLiteral
}

func newVisitor() visitor {
	return visitor{
		queue:              make(map[ast.TypeName][][]ast.IntegerLiteral),
		seenInstantiations: make(map[ast.TypeName][][]ast.IntegerLiteral),
	}
}

func (v visitor) newArgVisitor(numericalTypeArgs []ast.IntegerLiteral) visitor {
	return visitor{
		queue:              v.queue,
		seenInstantiations: v.seenInstantiations,
		numericalTypeArgs:  numericalTypeArgs,
	}
}

func (v visitor) newEnvVisitor(typeEnv map[ast.TypeParameter]ast.IntegerLiteral) visitor {
	return visitor{
		queue:              v.queue,
		seenInstantiations: v.seenInstantiations,
		typeEnv:            typeEnv,
	}
}

func (v visitor) monomorphise(m ast.MapVisitable) ast.MapVisitable {
	return m.AcceptMapVisitor(v)
}

func (v visitor) enqueue(typeName ast.TypeName, numericalTypeArgs []ast.IntegerLiteral) {
	if typeName == ast.IntTypeName {
		return
	}
	for _, instantiation := range v.seenInstantiations[typeName] {
		if cmp.Equal(instantiation, numericalTypeArgs, cmpopts.EquateEmpty()) {
			return
		}
	}
	v.queue[typeName] = append(v.queue[typeName], numericalTypeArgs)
	v.seenInstantiations[typeName] = append(v.seenInstantiations[typeName], numericalTypeArgs)
}

func (v visitor) dequeue(typeName ast.TypeName) []ast.IntegerLiteral {
	instantiations := v.queue[typeName]
	v.queue[typeName] = instantiations[1:]

	if len(instantiations) == 1 {
		delete(v.queue, typeName)
	}
	return instantiations[0]
}

func (v visitor) isEmpty(typeName ast.TypeName) bool {
	_, isDeclared := v.queue[typeName]
	return !isDeclared
}

func (v visitor) isQueueEmpty() bool {
	return len(v.queue) == 0
}

func (v visitor) monoTypeName(typeName ast.TypeName, numericalTypeArgs []ast.IntegerLiteral) ast.TypeName {
	monoName := string(typeName)
	for _, arg := range numericalTypeArgs {
		monoName += "__" + arg.String()
	}
	return ast.TypeName(monoName)
}

func (v visitor) newTypeEnvVisitor(t ast.TypeDeclaration) visitor {
	numericalTypeParams := make(map[ast.TypeParameter]ast.IntegerLiteral)

	numIndex := 0

	for _, typeParam := range t.TypeParameters {
		if _, isConstParam := typeParam.Bound.(ast.ConstType); isConstParam {
			numericalTypeParams[typeParam.TypeParameter] = v.numericalTypeArgs[numIndex]
			numIndex++
		}
	}
	envVisitor := v.newEnvVisitor(numericalTypeParams)
	envVisitor.numericalTypeArgs = v.numericalTypeArgs
	return envVisitor
}

type typeInstantiation struct {
	typeName          ast.TypeName
	numericalTypeArgs []ast.IntegerLiteral
}

type indexedTypeDeclaration struct {
	decl  ast.TypeDeclaration
	index int
}

type callableDeclaration interface {
	ast.Declaration
	GetMethodReceiver() ast.MethodReceiver
}

func (v visitor) MapProgram(p ast.Program) ast.MapVisitable {
	monoExpr := v.monomorphise(p.Expression).(ast.Expression)

	var typeDecls []indexedTypeDeclaration
	methodDeclIndices := make(map[ast.TypeName][]int)

	for i, decl := range p.Declarations {
		switch decl := decl.(type) {
		case ast.TypeDeclaration:
			typeDecls = append(typeDecls, indexedTypeDeclaration{decl, i})
		case callableDeclaration:
			typeName := decl.GetMethodReceiver().TypeName
			methodDeclIndices[typeName] = append(methodDeclIndices[typeName], i)
		}
	}

	// each declaration in original FGG program gets bucket for 0 or more
	// instantiations
	monoDecls := make([][]ast.Declaration, len(p.Declarations))

	for !v.isQueueEmpty() {
		for _, indexedDecl := range typeDecls {
			i := indexedDecl.index
			decl := indexedDecl.decl

			for !v.isEmpty(decl.TypeName) {
				inst := v.dequeue(decl.TypeName)
				argVisitor := v.newArgVisitor(inst)
				envVisitor := argVisitor.newTypeEnvVisitor(decl)

				monoDecls[i] = append(monoDecls[i], envVisitor.monomorphise(decl).(ast.Declaration))

				for _, methodIndex := range methodDeclIndices[decl.TypeName] {
					monoDecls[methodIndex] = append(monoDecls[methodIndex],
						envVisitor.monomorphise(p.Declarations[methodIndex]).(ast.Declaration))
				}
			}
		}
	}

	flatMonoDecl := []ast.Declaration{}
	for _, bucket := range monoDecls {
		for _, decl := range bucket {
			flatMonoDecl = append(flatMonoDecl, decl)
		}
	}

	return ast.Program{
		Declarations: flatMonoDecl,
		Expression:   monoExpr,
	}
}

func (v visitor) MapTypeDeclaration(t ast.TypeDeclaration) ast.MapVisitable {
	var monoNonConstTypeParams []ast.TypeParameterConstraint

	for _, typeParam := range t.TypeParameters {
		if _, isConstParam := typeParam.Bound.(ast.ConstType); !isConstParam {
			monoNonConstTypeParams = append(monoNonConstTypeParams, v.monomorphise(typeParam).(ast.TypeParameterConstraint))
		}
	}

	return ast.TypeDeclaration{
		TypeName:       v.monoTypeName(t.TypeName, v.numericalTypeArgs),
		TypeParameters: monoNonConstTypeParams,
		TypeLiteral:    v.monomorphise(t.TypeLiteral).(ast.TypeLiteral),
	}
}

func (v visitor) MapMethodDeclaration(m ast.MethodDeclaration) ast.MapVisitable {
	return ast.MethodDeclaration{
		MethodReceiver:      v.monomorphise(m.MethodReceiver).(ast.MethodReceiver),
		MethodSpecification: v.monomorphise(m.MethodSpecification).(ast.MethodSpecification),
		ReturnExpression:    v.monomorphise(m.ReturnExpression).(ast.Expression),
	}
}

func (v visitor) MapArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) ast.MapVisitable {
	return ast.ArraySetMethodDeclaration{
		MethodReceiver:        v.monomorphise(a.MethodReceiver).(ast.MethodReceiver),
		MethodName:            a.MethodName,
		IndexParameter:        a.IndexParameter, // should always be int, so no need to monomorphise
		ValueParameter:        v.monomorphise(a.ValueParameter).(ast.MethodParameter),
		ReturnType:            v.monomorphise(a.ReturnType).(ast.Type),
		IndexReceiverVariable: a.IndexReceiverVariable,
		IndexVariable:         a.IndexVariable,
		NewValueVariable:      a.NewValueVariable,
		ReturnVariable:        a.ReturnVariable,
	}
}

func (v visitor) MapMethodReceiver(m ast.MethodReceiver) ast.MapVisitable {
	var monoTypeParams []ast.TypeParameter
	for _, typeParam := range m.TypeParameters {
		if _, isConst := v.typeEnv[typeParam]; !isConst {
			monoTypeParams = append(monoTypeParams, typeParam)
		}
	}

	return ast.MethodReceiver{
		ParameterName:  m.ParameterName,
		TypeName:       v.monoTypeName(m.TypeName, v.numericalTypeArgs),
		TypeParameters: monoTypeParams,
	}
}

func (v visitor) MapTypeParameterConstraint(t ast.TypeParameterConstraint) ast.MapVisitable {
	return ast.TypeParameterConstraint{
		TypeParameter: t.TypeParameter,
		Bound:         v.monomorphise(t.Bound).(ast.Bound),
	}
}

func (v visitor) MapStructTypeLiteral(s ast.StructTypeLiteral) ast.MapVisitable {
	monoFields := make([]ast.Field, 0, len(s.Fields))
	for _, field := range s.Fields {
		monoFields = append(monoFields, v.monomorphise(field).(ast.Field))
	}
	return ast.StructTypeLiteral{
		Fields: monoFields,
	}
}

func (v visitor) MapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.MapVisitable {
	var monoSpecs []ast.MethodSpecification
	for _, spec := range i.MethodSpecifications {
		monoSpecs = append(monoSpecs, v.monomorphise(spec).(ast.MethodSpecification))
	}
	return ast.InterfaceTypeLiteral{
		MethodSpecifications: monoSpecs,
	}
}

func (v visitor) MapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.MapVisitable {
	return ast.ArrayTypeLiteral{
		Length:      v.monomorphise(a.Length).(ast.IntegerLiteral),
		ElementType: v.monomorphise(a.ElementType).(ast.Type),
	}
}

func (v visitor) MapMethodSpecification(m ast.MethodSpecification) ast.MapVisitable {
	return ast.MethodSpecification{
		MethodName:      m.MethodName,
		MethodSignature: v.monomorphise(m.MethodSignature).(ast.MethodSignature),
	}
}

func (v visitor) MapIntegerLiteral(i ast.IntegerLiteral) ast.MapVisitable {
	return i
}

func (v visitor) MapVariable(variable ast.Variable) ast.MapVisitable {
	return variable
}

func (v visitor) MapMethodCall(m ast.MethodCall) ast.MapVisitable {
	var monoArgs []ast.Expression
	for _, arg := range m.Arguments {
		monoArgs = append(monoArgs, v.monomorphise(arg).(ast.Expression))
	}

	return ast.MethodCall{
		Receiver:   v.monomorphise(m.Receiver).(ast.Expression),
		MethodName: m.MethodName,
		Arguments:  monoArgs,
	}
}

func (v visitor) MapValueLiteral(value ast.ValueLiteral) ast.MapVisitable {
	monoType := v.monomorphise(value.Type).(ast.Type)
	monoValues := make([]ast.Expression, 0, len(value.Values))
	for _, val := range value.Values {
		monoValues = append(monoValues, v.monomorphise(val).(ast.Expression))
	}
	return ast.ValueLiteral{
		Type:   monoType,
		Values: monoValues,
	}
}

func (v visitor) MapSelect(s ast.Select) ast.MapVisitable {
	return ast.Select{
		Receiver:  v.monomorphise(s.Receiver).(ast.Expression),
		FieldName: s.FieldName,
	}
}

func (v visitor) MapArrayIndex(a ast.ArrayIndex) ast.MapVisitable {
	return ast.ArrayIndex{
		Receiver: v.monomorphise(a.Receiver).(ast.Expression),
		Index:    v.monomorphise(a.Index).(ast.Expression),
	}
}

func (v visitor) MapMethodParameter(p ast.MethodParameter) ast.MapVisitable {
	return ast.MethodParameter{
		ParameterName: p.ParameterName,
		Type:          v.monomorphise(p.Type).(ast.Type),
	}
}

func (v visitor) MapConstType(c ast.ConstType) ast.MapVisitable {
	panic("const type parameter should have been eliminated")
}

func (v visitor) MapNamedType(n ast.NamedType) ast.MapVisitable {
	// instantiate any numerical type params
	var monoTypeArgs = make([]ast.Type, 0, len(n.TypeArguments))
	for _, arg := range n.TypeArguments {
		monoTypeArgs = append(monoTypeArgs, v.monomorphise(arg).(ast.Type))
	}

	// split numerical and non-numerical type args
	var numericalTypeArgs []ast.IntegerLiteral
	var monoNonNumericTypeArgs []ast.Type

	for _, arg := range monoTypeArgs {
		if numArg, isNumArg := arg.(ast.IntegerLiteral); isNumArg {
			numericalTypeArgs = append(numericalTypeArgs, numArg)
		} else {
			monoNonNumericTypeArgs = append(monoNonNumericTypeArgs, arg)
		}
	}
	// enqueue the instantiation
	v.enqueue(n.TypeName, numericalTypeArgs)

	// monomorphise the instantation
	return ast.NamedType{
		TypeName:      v.monoTypeName(n.TypeName, numericalTypeArgs),
		TypeArguments: monoNonNumericTypeArgs,
	}
}

func (v visitor) MapTypeParameter(t ast.TypeParameter) ast.MapVisitable {
	if typeArg, shouldMonomorphise := v.typeEnv[t]; shouldMonomorphise {
		return typeArg
	}
	return t
}

func (v visitor) MapField(f ast.Field) ast.MapVisitable {
	return ast.Field{
		Name: f.Name,
		Type: v.monomorphise(f.Type).(ast.Type),
	}
}

func (v visitor) MapMethodSignature(m ast.MethodSignature) ast.MapVisitable {
	var monoParams []ast.MethodParameter
	for _, param := range m.MethodParameters {
		monoParams = append(monoParams, v.monomorphise(param).(ast.MethodParameter))
	}

	return ast.MethodSignature{
		MethodParameters: monoParams,
		ReturnType:       v.monomorphise(m.ReturnType).(ast.Type),
	}
}
