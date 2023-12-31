package typecheck

import (
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

type typeParamIdentifier struct {
	typeEnvTypeCheckingVisitor
}

func (t typeParamIdentifier) identifyTypeParams(v ast.EnvVisitable) ast.EnvVisitable {
	return v.AcceptEnvMapperVisitor(t)
}

func (typeParamIdentifier) VisitMapConstType(c ast.ConstType) ast.EnvVisitable {
	return c
}

func (typeParamIdentifier) VisitMapTypeParameter(t ast.TypeParameter) ast.EnvVisitable {
	return t
}

func (t typeParamIdentifier) VisitMapNamedType(n ast.NamedType) ast.EnvVisitable {
	// TODO what happens in case type param shadows type decl? Is this allowed in Go?
	typeParam := ast.TypeParameter(n.TypeName)
	if _, isTypeParam := t.typeEnv[typeParam]; isTypeParam {
		return typeParam
	}
	typeArgs := slices.Clone(n.TypeArguments)
	for i, typeArg := range n.TypeArguments {
		if namedTypeArg, isNamedTypeArg := typeArg.(ast.NamedType); isNamedTypeArg {
			typeParam := ast.TypeParameter(namedTypeArg.TypeName)
			if _, isTypeParam := t.typeEnv[typeParam]; isTypeParam {
				typeArgs[i] = typeParam
			}
		}
	}
	return ast.NamedType{
		TypeName:      n.TypeName,
		TypeArguments: typeArgs,
	}
}

func (t typeParamIdentifier) VisitMapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.EnvVisitable {
	return ast.ArrayTypeLiteral{
		Length:      t.identifyTypeParams(a.Length).(ast.Type),
		ElementType: t.identifyTypeParams(a.ElementType).(ast.Type),
	}
}

func (t typeParamIdentifier) VisitMapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.EnvVisitable {
	methodSpecs := make([]ast.MethodSpecification, 0, len(i.MethodSpecifications))

	for _, spec := range i.MethodSpecifications {
		methodSpecs = append(methodSpecs, t.identifyTypeParams(spec).(ast.MethodSpecification))
	}

	return ast.InterfaceTypeLiteral{
		MethodSpecifications: methodSpecs,
	}
}

func (t typeParamIdentifier) VisitMapMethodSpecification(m ast.MethodSpecification) ast.EnvVisitable {
	return ast.MethodSpecification{
		MethodName: m.MethodName,
		MethodSignature: ast.MethodSignature{
			MethodParameters: t.identifyTypeParamsInMethodParams(m),
			ReturnType:       t.identifyTypeParams(m.MethodSignature.ReturnType).(ast.Type),
		},
	}
}

func (t typeParamIdentifier) identifyTypeParamsInMethodParams(spec ast.MethodSpecification) []ast.MethodParameter {
	params := make([]ast.MethodParameter, 0, len(spec.MethodSignature.MethodParameters))
	for _, param := range spec.MethodSignature.MethodParameters {
		params = append(params, ast.MethodParameter{
			ParameterName: param.ParameterName,
			Type:          t.identifyTypeParams(param.Type).(ast.Type),
		})
	}
	return params
}

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
	//TODO implement me
	panic("implement me")
}

// TODO will these ever be called? If not, a separate interface can be extracted

func (s typeParamSubstituter) VisitMapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.EnvVisitable {
	//TODO implement me
	panic("implement me")
}

func (s typeParamSubstituter) VisitMapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.EnvVisitable {
	//TODO implement me
	panic("implement me")
}
