package preprocessor

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

func IdentifyTypeParams(p ast.Program) ast.Program {
	return typeParamVisitor{}.identifyTypeParams(p).(ast.Program)
}

type typeParamVisitor struct {
	typeEnv map[ast.TypeParameter]struct{}
}

func (t typeParamVisitor) identifyTypeParams(v ast.MappingVisitable) ast.MappingVisitable {
	return v.AcceptMappingVisitor(t)
}

func (t typeParamVisitor) VisitMapProgram(p ast.Program) ast.MappingVisitable {
	identifiedDecls := make([]ast.Declaration, 0, len(p.Declarations))
	for _, decl := range p.Declarations {
		identifiedDecls = append(identifiedDecls, t.identifyTypeParams(decl).(ast.Declaration))
	}
	return ast.Program{
		Declarations: identifiedDecls,
		Expression:   p.Expression,
	}
}

func (t typeParamVisitor) VisitMapTypeDeclaration(td ast.TypeDeclaration) ast.MappingVisitable {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range td.TypeParameters {
		envVisitor.typeEnv[param.TypeParameter] = struct{}{}
	}

	identifiedTypeParams := make([]ast.TypeParameterConstraint, 0, len(td.TypeParameters))
	for _, param := range td.TypeParameters {
		identified := envVisitor.identifyTypeParams(param).(ast.TypeParameterConstraint)
		identifiedTypeParams = append(identifiedTypeParams, identified)
	}
	return ast.TypeDeclaration{
		TypeName:       td.TypeName,
		TypeParameters: identifiedTypeParams,
		TypeLiteral:    envVisitor.identifyTypeParams(td.TypeLiteral).(ast.TypeLiteral),
	}
}

func (t typeParamVisitor) VisitMapMethodDeclaration(m ast.MethodDeclaration) ast.MappingVisitable {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range m.MethodReceiver.TypeParameters {
		envVisitor.typeEnv[param] = struct{}{}
	}

	return ast.MethodDeclaration{
		MethodReceiver:      m.MethodReceiver,
		MethodSpecification: envVisitor.identifyTypeParams(m.MethodSpecification).(ast.MethodSpecification),
		ReturnExpression:    envVisitor.identifyTypeParams(m.ReturnExpression).(ast.Expression),
	}
}

func (t typeParamVisitor) VisitMapArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) ast.MappingVisitable {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range a.MethodReceiver.TypeParameters {
		envVisitor.typeEnv[param] = struct{}{}
	}

	return ast.ArraySetMethodDeclaration{
		MethodReceiver:        a.MethodReceiver,
		MethodName:            a.MethodName,
		IndexParameter:        envVisitor.identifyTypeParams(a.IndexParameter).(ast.MethodParameter),
		ValueParameter:        envVisitor.identifyTypeParams(a.ValueParameter).(ast.MethodParameter),
		ReturnType:            envVisitor.identifyTypeParams(a.ReturnType).(ast.Type),
		IndexReceiverVariable: a.IndexReceiverVariable,
		IndexVariable:         a.IndexVariable,
		NewValueVariable:      a.NewValueVariable,
		ReturnVariable:        a.ReturnVariable,
	}
}

func (t typeParamVisitor) VisitMapTypeParameterConstraint(param ast.TypeParameterConstraint) ast.MappingVisitable {
	return ast.TypeParameterConstraint{
		TypeParameter: param.TypeParameter,
		Bound:         t.identifyTypeParams(param.Bound).(ast.Bound),
	}
}

func (t typeParamVisitor) VisitMapStructTypeLiteral(s ast.StructTypeLiteral) ast.MappingVisitable {
	identifiedFields := make([]ast.Field, 0, len(s.Fields))
	for _, field := range s.Fields {
		identifiedFields = append(identifiedFields, t.identifyTypeParams(field).(ast.Field))
	}
	return ast.StructTypeLiteral{
		Fields: identifiedFields,
	}
}

func (t typeParamVisitor) VisitMapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.MappingVisitable {
	identifiedSpecs := make([]ast.MethodSpecification, 0, len(i.MethodSpecifications))
	for _, spec := range i.MethodSpecifications {
		identifiedSpecs = append(identifiedSpecs, t.identifyTypeParams(spec).(ast.MethodSpecification))
	}
	return ast.InterfaceTypeLiteral{
		MethodSpecifications: identifiedSpecs,
	}
}

func (t typeParamVisitor) VisitMapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.MappingVisitable {
	return ast.ArrayTypeLiteral{
		Length:      t.identifyTypeParams(a.Length).(ast.Type),
		ElementType: t.identifyTypeParams(a.ElementType).(ast.Type),
	}
}

func (t typeParamVisitor) VisitMapMethodSpecification(m ast.MethodSpecification) ast.MappingVisitable {
	return ast.MethodSpecification{
		MethodName:      m.MethodName,
		MethodSignature: t.identifyTypeParams(m.MethodSignature).(ast.MethodSignature),
	}
}

func (t typeParamVisitor) VisitMapIntegerLiteral(i ast.IntegerLiteral) ast.MappingVisitable {
	return i
}

func (t typeParamVisitor) VisitMapVariable(v ast.Variable) ast.MappingVisitable {
	return v
}

func (t typeParamVisitor) VisitMapMethodCall(m ast.MethodCall) ast.MappingVisitable {
	identifiedArgs := make([]ast.Expression, 0, len(m.Arguments))
	for _, arg := range m.Arguments {
		identifiedArgs = append(identifiedArgs, t.identifyTypeParams(arg).(ast.Expression))
	}
	return ast.MethodCall{
		Receiver:   t.identifyTypeParams(m.Receiver).(ast.Expression),
		MethodName: m.MethodName,
		Arguments:  identifiedArgs,
	}
}

func (t typeParamVisitor) VisitMapValueLiteral(v ast.ValueLiteral) ast.MappingVisitable {
	identifiedValues := make([]ast.Expression, 0, len(v.Values))
	for _, val := range v.Values {
		identifiedValues = append(identifiedValues, t.identifyTypeParams(val).(ast.Expression))
	}
	return ast.ValueLiteral{
		Type:   t.identifyTypeParams(v.Type).(ast.Type),
		Values: identifiedValues,
	}
}

func (t typeParamVisitor) VisitMapSelect(s ast.Select) ast.MappingVisitable {
	return ast.Select{
		Receiver:  t.identifyTypeParams(s.Receiver).(ast.Expression),
		FieldName: s.FieldName,
	}
}

func (t typeParamVisitor) VisitMapArrayIndex(a ast.ArrayIndex) ast.MappingVisitable {
	return ast.ArrayIndex{
		Receiver: t.identifyTypeParams(a.Receiver).(ast.Expression),
		Index:    t.identifyTypeParams(a.Index).(ast.Expression),
	}
}

func (t typeParamVisitor) VisitMapMethodParameter(p ast.MethodParameter) ast.MappingVisitable {
	return ast.MethodParameter{
		ParameterName: p.ParameterName,
		Type:          t.identifyTypeParams(p.Type).(ast.Type),
	}
}

func (t typeParamVisitor) VisitMapConstType(c ast.ConstType) ast.MappingVisitable {
	return c
}

func (t typeParamVisitor) VisitMapNamedType(n ast.NamedType) ast.MappingVisitable {
	// TODO if there are type arguments, then let the type checker fail this somehow
	// how do we do this correctly if a type parameter hides a valid named type?
	// might have to error here instead (is it a type or syntax error in actual Go?)
	if _, isTypeParam := t.typeEnv[ast.TypeParameter(n.TypeName)]; isTypeParam {
		return ast.TypeParameter(n.TypeName)
	}
	identifiedArgs := make([]ast.Type, 0, len(n.TypeArguments))
	for _, arg := range n.TypeArguments {
		identifiedArgs = append(identifiedArgs, t.identifyTypeParams(arg).(ast.Type))
	}
	return ast.NamedType{
		TypeName:      n.TypeName,
		TypeArguments: identifiedArgs,
	}
}

func (t typeParamVisitor) VisitMapTypeParameter(param ast.TypeParameter) ast.MappingVisitable {
	return param
}

func (t typeParamVisitor) VisitMapField(f ast.Field) ast.MappingVisitable {
	return ast.Field{
		Name: f.Name,
		Type: t.identifyTypeParams(f.Type).(ast.Type),
	}
}

func (t typeParamVisitor) VisitMapMethodSignature(m ast.MethodSignature) ast.MappingVisitable {
	identifiedParams := make([]ast.MethodParameter, 0, len(m.MethodParameters))
	for _, param := range m.MethodParameters {
		identifiedParams = append(identifiedParams, t.identifyTypeParams(param).(ast.MethodParameter))
	}
	return ast.MethodSignature{
		MethodParameters: identifiedParams,
		ReturnType:       t.identifyTypeParams(m.ReturnType).(ast.Type),
	}
}
