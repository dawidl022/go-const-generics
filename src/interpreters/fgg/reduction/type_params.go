package reduction

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

type typeParamSubstituter struct {
	substitutions map[ast.TypeParameter]ast.Type
}

func newTypeParamSubstituter(
	typeParams []ast.TypeParameter,
	typeArgs []ast.Type,
) (typeParamSubstituter, error) {
	substitutions, err := makeTypeSubstitutions(typeParams, typeArgs)
	return typeParamSubstituter{
		substitutions: substitutions,
	}, err
}

func makeTypeSubstitutions(typeParams []ast.TypeParameter, typeArgs []ast.Type) (map[ast.TypeParameter]ast.Type, error) {
	if len(typeArgs) != len(typeParams) {
		return nil, fmt.Errorf("expected %d type arguments but got %d", len(typeParams), len(typeArgs))
	}
	typeSubstitutions := make(map[ast.TypeParameter]ast.Type)
	for i, typeParam := range typeParams {
		typeSubstitutions[typeParam] = typeArgs[i]
	}
	return typeSubstitutions, nil
}

func (t typeParamSubstituter) substituteTypeParams(e ast.Expression) (ast.Expression, error) {
	return e.Accept(t)
}

func (t typeParamSubstituter) VisitIntegerLiteral(i ast.IntegerLiteral) (ast.Expression, error) {
	return i, nil
}

func (t typeParamSubstituter) VisitVariable(v ast.Variable) (ast.Expression, error) {
	return v, nil
}

func (t typeParamSubstituter) VisitMethodCall(m ast.MethodCall) (ast.Expression, error) {
	receiver, err := t.substituteTypeParams(m.Receiver)
	if err != nil {
		return nil, err
	}

	var args []ast.Expression
	for _, arg := range m.Arguments {
		substitutedArg, err := t.substituteTypeParams(arg)
		if err != nil {
			return nil, err
		}
		args = append(args, substitutedArg)
	}

	return ast.MethodCall{
		Receiver:   receiver,
		MethodName: m.MethodName,
		Arguments:  args,
	}, nil
}

func (t typeParamSubstituter) VisitValueLiteral(v ast.ValueLiteral) (ast.Expression, error) {
	var typeArgs []ast.Type
	for _, typeArg := range v.Type.TypeArguments {
		typeParam, isTypeParam := typeArg.(ast.TypeParameter)
		if isTypeParam {
			typeArgs = append(typeArgs, t.substitutions[typeParam])
		} else {
			typeArgs = append(typeArgs, typeArg)
		}
	}

	var values []ast.Expression
	for _, value := range v.Values {
		substitutedValue, err := t.substituteTypeParams(value)
		if err != nil {
			return nil, err
		}
		values = append(values, substitutedValue)
	}

	return ast.ValueLiteral{Type: ast.NamedType{
		TypeName: v.Type.TypeName, TypeArguments: typeArgs},
		Values: values,
	}, nil
}

func (t typeParamSubstituter) VisitSelect(s ast.Select) (ast.Expression, error) {
	receiver, err := t.substituteTypeParams(s.Receiver)
	return ast.Select{
		Receiver:  receiver,
		FieldName: s.FieldName,
	}, err
}

func (t typeParamSubstituter) VisitArrayIndex(a ast.ArrayIndex) (ast.Expression, error) {
	receiver, err := t.substituteTypeParams(a.Receiver)
	if err != nil {
		return nil, err
	}
	index, err := t.substituteTypeParams(a.Index)
	return ast.ArrayIndex{
		Receiver: receiver,
		Index:    index,
	}, err
}

func (t typeParamSubstituter) VisitAdd(a ast.Add) (ast.Expression, error) {
	left, err := t.substituteTypeParams(a.Left)
	if err != nil {
		return nil, err
	}
	right, err := t.substituteTypeParams(a.Right)
	return ast.Add{
		Left:  left,
		Right: right,
	}, err
}
