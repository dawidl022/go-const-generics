package typecheck

import "github.com/dawidl022/go-const-generics/interpreters/fgg/ast"

type typeParamSubstituter struct {
	substitutions map[ast.TypeParameter]ast.Type
}

func newTypeParamSubstituter(
	typeArguments []ast.Type,
	typeParams []ast.TypeParameterConstraint,
) (typeParamSubstituter, error) {
	substitutions, err := makeTypeSubstitutions(typeArguments, typeParams)
	return typeParamSubstituter{
		substitutions: substitutions,
	}, err
}

func (s typeParamSubstituter) substituteTypeParams(e ast.EnvVisitable) ast.EnvVisitable {
	switch e.(type) {
	case ast.NamedType:
		e := e.(ast.NamedType)
		var substitutedTypeArgs []ast.Type

		for _, typeArg := range e.TypeArguments {
			substitutedTypeArgs = append(substitutedTypeArgs, s.substituteTypeParams(typeArg).(ast.Type))
		}
		return ast.NamedType{
			TypeName:      e.TypeName,
			TypeArguments: substitutedTypeArgs,
		}
	default:
		return e.AcceptEnvMapperVisitor(s)
	}
}

func (s typeParamSubstituter) VisitMapTypeParameter(t ast.TypeParameter) ast.EnvVisitable {
	return s.substitutions[t]
}

func (s typeParamSubstituter) VisitMapNamedType(n ast.NamedType) ast.EnvVisitable {
	typeArgs := make([]ast.Type, 0, len(n.TypeArguments))
	for _, typeArg := range n.TypeArguments {
		typeArgs = append(typeArgs, s.substituteTypeParams(typeArg).(ast.Type))
	}
	return ast.NamedType{
		TypeName:      n.TypeName,
		TypeArguments: typeArgs,
	}
}

func (s typeParamSubstituter) VisitMapMethodSpecification(m ast.MethodSpecification) ast.EnvVisitable {
	params := make([]ast.MethodParameter, 0, len(m.MethodSignature.MethodParameters))
	for _, param := range m.MethodSignature.MethodParameters {
		params = append(params, ast.MethodParameter{
			ParameterName: param.ParameterName,
			Type:          s.substituteTypeParams(param.Type).(ast.Type),
		})
	}
	return ast.MethodSpecification{
		MethodName: m.MethodName,
		MethodSignature: ast.MethodSignature{
			MethodParameters: params,
			ReturnType:       s.substituteTypeParams(m.MethodSignature.ReturnType).(ast.Type),
		},
	}
}

func (s typeParamSubstituter) VisitMapConstType(c ast.ConstType) ast.EnvVisitable {
	return c
}

func (s typeParamSubstituter) VisitMapIntegerLiteral(i ast.IntegerLiteral) ast.EnvVisitable {
	return i
}

func (s typeParamSubstituter) VisitMapStructTypeLiteral(st ast.StructTypeLiteral) ast.EnvVisitable {
	substitutedFields := []ast.Field{}
	for _, field := range st.Fields {
		substitutedFields = append(substitutedFields, ast.Field{
			Name: field.Name,
			Type: s.substituteTypeParams(field.Type).(ast.Type),
		})
	}
	return ast.StructTypeLiteral{Fields: substitutedFields}
}

func (s typeParamSubstituter) VisitMapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.EnvVisitable {
	return ast.ArrayTypeLiteral{
		Length:      a.Length,
		ElementType: s.substituteTypeParams(a.ElementType).(ast.Type),
	}
}

func (s typeParamSubstituter) VisitMapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.EnvVisitable {
	substitutedSpecs := []ast.MethodSpecification{}
	for _, spec := range i.MethodSpecifications {
		substitutedSpecs = append(substitutedSpecs, s.substituteTypeParams(spec).(ast.MethodSpecification))
	}
	return ast.InterfaceTypeLiteral{MethodSpecifications: substitutedSpecs}
}
