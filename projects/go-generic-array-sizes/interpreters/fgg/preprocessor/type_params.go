package preprocessor

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func IdentifyTypeParams(p ast.Program) (ast.Program, error) {
	identified, err := typeParamVisitor{}.identifyTypeParams(p)
	if err != nil {
		return ast.Program{}, err
	}
	return identified.(ast.Program), nil
}

type typeParamVisitor struct {
	typeEnv map[ast.TypeParameter]struct{}
}

func (t typeParamVisitor) identifyTypeParams(v ast.MappingVisitable) (ast.MappingVisitable, error) {
	return v.AcceptMappingVisitor(t)
}

func (t typeParamVisitor) VisitMapProgram(p ast.Program) (ast.MappingVisitable, error) {
	identifiedDecls := make([]ast.Declaration, 0, len(p.Declarations))
	for _, decl := range p.Declarations {
		identified, err := t.identifyTypeParams(decl)
		if err != nil {
			return nil, err
		}
		identifiedDecls = append(identifiedDecls, identified.(ast.Declaration))
	}
	return ast.Program{
		Declarations: identifiedDecls,
		Expression:   p.Expression,
	}, nil
}

func (t typeParamVisitor) VisitMapTypeDeclaration(td ast.TypeDeclaration) (ast.MappingVisitable, error) {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range td.TypeParameters {
		envVisitor.typeEnv[param.TypeParameter] = struct{}{}
	}

	identifiedTypeParams := make([]ast.TypeParameterConstraint, 0, len(td.TypeParameters))
	for _, param := range td.TypeParameters {
		identified, err := envVisitor.identifyTypeParams(param)
		if err != nil {
			return nil, err
		}
		identifiedTypeParams = append(identifiedTypeParams, identified.(ast.TypeParameterConstraint))
	}
	identifiedTypeLiteral, err := envVisitor.identifyTypeParams(td.TypeLiteral)
	if err != nil {
		return nil, err
	}
	return ast.TypeDeclaration{
		TypeName:       td.TypeName,
		TypeParameters: identifiedTypeParams,
		TypeLiteral:    identifiedTypeLiteral.(ast.TypeLiteral),
	}, nil
}

func (t typeParamVisitor) VisitMapMethodDeclaration(m ast.MethodDeclaration) (ast.MappingVisitable, error) {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range m.MethodReceiver.TypeParameters {
		envVisitor.typeEnv[param] = struct{}{}
	}

	identifiedSpec, err := envVisitor.identifyTypeParams(m.MethodSpecification)
	if err != nil {
		return nil, err
	}
	identifiedExpr, err := envVisitor.identifyTypeParams(m.ReturnExpression)
	if err != nil {
		return nil, err
	}

	return ast.MethodDeclaration{
		MethodReceiver:      m.MethodReceiver,
		MethodSpecification: identifiedSpec.(ast.MethodSpecification),
		ReturnExpression:    identifiedExpr.(ast.Expression),
	}, nil
}

func (t typeParamVisitor) VisitMapArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) (ast.MappingVisitable, error) {
	envVisitor := typeParamVisitor{typeEnv: make(map[ast.TypeParameter]struct{})}
	for _, param := range a.MethodReceiver.TypeParameters {
		envVisitor.typeEnv[param] = struct{}{}
	}

	identifiedIndex, err := envVisitor.identifyTypeParams(a.IndexParameter)
	if err != nil {
		return nil, err
	}
	identifiedValue, err := envVisitor.identifyTypeParams(a.ValueParameter)
	if err != nil {
		return nil, err
	}
	identifiedReturn, err := envVisitor.identifyTypeParams(a.ReturnType)
	if err != nil {
		return nil, err
	}

	return ast.ArraySetMethodDeclaration{
		MethodReceiver:        a.MethodReceiver,
		MethodName:            a.MethodName,
		IndexParameter:        identifiedIndex.(ast.MethodParameter),
		ValueParameter:        identifiedValue.(ast.MethodParameter),
		ReturnType:            identifiedReturn.(ast.Type),
		IndexReceiverVariable: a.IndexReceiverVariable,
		IndexVariable:         a.IndexVariable,
		NewValueVariable:      a.NewValueVariable,
		ReturnVariable:        a.ReturnVariable,
	}, nil
}

func (t typeParamVisitor) VisitMapTypeParameterConstraint(param ast.TypeParameterConstraint) (ast.MappingVisitable, error) {
	identifiedBound, err := t.identifyTypeParams(param.Bound)
	if err != nil {
		return nil, err
	}
	return ast.TypeParameterConstraint{
		TypeParameter: param.TypeParameter,
		Bound:         identifiedBound.(ast.Bound),
	}, nil
}

func (t typeParamVisitor) VisitMapStructTypeLiteral(s ast.StructTypeLiteral) (ast.MappingVisitable, error) {
	identifiedFields := make([]ast.Field, 0, len(s.Fields))
	for _, field := range s.Fields {
		identified, err := t.identifyTypeParams(field)
		if err != nil {
			return nil, err
		}
		identifiedFields = append(identifiedFields, identified.(ast.Field))
	}
	return ast.StructTypeLiteral{
		Fields: identifiedFields,
	}, nil
}

func (t typeParamVisitor) VisitMapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) (ast.MappingVisitable, error) {
	identifiedSpecs := make([]ast.MethodSpecification, 0, len(i.MethodSpecifications))
	for _, spec := range i.MethodSpecifications {
		identified, err := t.identifyTypeParams(spec)
		if err != nil {
			return nil, err
		}
		identifiedSpecs = append(identifiedSpecs, identified.(ast.MethodSpecification))
	}
	return ast.InterfaceTypeLiteral{
		MethodSpecifications: identifiedSpecs,
	}, nil
}

func (t typeParamVisitor) VisitMapArrayTypeLiteral(a ast.ArrayTypeLiteral) (ast.MappingVisitable, error) {
	identifiedLength, err := t.identifyTypeParams(a.Length)
	if err != nil {
		return nil, err
	}
	identifiedElement, err := t.identifyTypeParams(a.ElementType)
	if err != nil {
		return nil, err
	}

	return ast.ArrayTypeLiteral{
		Length:      identifiedLength.(ast.Type),
		ElementType: identifiedElement.(ast.Type),
	}, nil
}

func (t typeParamVisitor) VisitMapMethodSpecification(m ast.MethodSpecification) (ast.MappingVisitable, error) {
	identifiedSignature, err := t.identifyTypeParams(m.MethodSignature)
	if err != nil {
		return nil, err
	}

	return ast.MethodSpecification{
		MethodName:      m.MethodName,
		MethodSignature: identifiedSignature.(ast.MethodSignature),
	}, nil
}

func (t typeParamVisitor) VisitMapIntegerLiteral(i ast.IntegerLiteral) (ast.MappingVisitable, error) {
	return i, nil
}

func (t typeParamVisitor) VisitMapVariable(v ast.Variable) (ast.MappingVisitable, error) {
	return v, nil
}

func (t typeParamVisitor) VisitMapMethodCall(m ast.MethodCall) (ast.MappingVisitable, error) {
	identifiedArgs := make([]ast.Expression, 0, len(m.Arguments))
	for _, arg := range m.Arguments {
		identified, err := t.identifyTypeParams(arg)
		if err != nil {
			return nil, err
		}
		identifiedArgs = append(identifiedArgs, identified.(ast.Expression))
	}
	identifiedReceiver, err := t.identifyTypeParams(m.Receiver)
	if err != nil {
		return nil, err
	}

	return ast.MethodCall{
		Receiver:   identifiedReceiver.(ast.Expression),
		MethodName: m.MethodName,
		Arguments:  identifiedArgs,
	}, nil
}

func (t typeParamVisitor) VisitMapValueLiteral(v ast.ValueLiteral) (ast.MappingVisitable, error) {
	identifiedValues := make([]ast.Expression, 0, len(v.Values))
	for _, val := range v.Values {
		identified, err := t.identifyTypeParams(val)
		if err != nil {
			return nil, err
		}
		identifiedValues = append(identifiedValues, identified.(ast.Expression))
	}
	identifiedParams, err := t.identifyTypeParams(v.Type)
	if err != nil {
		return nil, err
	}
	return ast.ValueLiteral{
		Type:   identifiedParams.(ast.Type),
		Values: identifiedValues,
	}, nil
}

func (t typeParamVisitor) VisitMapSelect(s ast.Select) (ast.MappingVisitable, error) {
	identifiedReceiver, err := t.identifyTypeParams(s.Receiver)
	if err != nil {
		return nil, err
	}
	return ast.Select{
		Receiver:  identifiedReceiver.(ast.Expression),
		FieldName: s.FieldName,
	}, nil
}

func (t typeParamVisitor) VisitMapArrayIndex(a ast.ArrayIndex) (ast.MappingVisitable, error) {
	identifiedReceiver, err := t.identifyTypeParams(a.Receiver)
	if err != nil {
		return nil, err
	}
	identifiedIndex, err := t.identifyTypeParams(a.Index)
	if err != nil {
		return nil, err
	}

	return ast.ArrayIndex{
		Receiver: identifiedReceiver.(ast.Expression),
		Index:    identifiedIndex.(ast.Expression),
	}, nil
}

func (t typeParamVisitor) VisitMapMethodParameter(p ast.MethodParameter) (ast.MappingVisitable, error) {
	identifiedParams, err := t.identifyTypeParams(p.Type)
	if err != nil {
		return nil, err
	}
	return ast.MethodParameter{
		ParameterName: p.ParameterName,
		Type:          identifiedParams.(ast.Type),
	}, nil
}

func (t typeParamVisitor) VisitMapConstType(c ast.ConstType) (ast.MappingVisitable, error) {
	return c, nil
}

func (t typeParamVisitor) VisitMapNamedType(n ast.NamedType) (ast.MappingVisitable, error) {
	if _, isTypeParam := t.typeEnv[ast.TypeParameter(n.TypeName)]; isTypeParam {
		if len(n.TypeArguments) > 0 {
			return nil, fmt.Errorf("type parameter %q does not accept any type arguments", n.TypeName)
		}
		return ast.TypeParameter(n.TypeName), nil
	}
	identifiedArgs := make([]ast.Type, 0, len(n.TypeArguments))
	for _, arg := range n.TypeArguments {
		identified, err := t.identifyTypeParams(arg)
		if err != nil {
			return nil, err
		}
		identifiedArgs = append(identifiedArgs, identified.(ast.Type))
	}
	return ast.NamedType{
		TypeName:      n.TypeName,
		TypeArguments: identifiedArgs,
	}, nil
}

func (t typeParamVisitor) VisitMapTypeParameter(param ast.TypeParameter) (ast.MappingVisitable, error) {
	return param, nil
}

func (t typeParamVisitor) VisitMapField(f ast.Field) (ast.MappingVisitable, error) {
	identifiedType, err := t.identifyTypeParams(f.Type)
	if err != nil {
		return nil, err
	}
	return ast.Field{
		Name: f.Name,
		Type: identifiedType.(ast.Type),
	}, nil
}

func (t typeParamVisitor) VisitMapMethodSignature(m ast.MethodSignature) (ast.MappingVisitable, error) {
	identifiedParams := make([]ast.MethodParameter, 0, len(m.MethodParameters))
	for _, param := range m.MethodParameters {
		identified, err := t.identifyTypeParams(param)
		if err != nil {
			return nil, err
		}
		identifiedParams = append(identifiedParams, identified.(ast.MethodParameter))
	}
	identifiedReturn, err := t.identifyTypeParams(m.ReturnType)
	if err != nil {
		return nil, err
	}

	return ast.MethodSignature{
		MethodParameters: identifiedParams,
		ReturnType:       identifiedReturn.(ast.Type),
	}, nil
}
