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

func (t typeParamIdentifier) VisitMapConstType(c ast.ConstType) ast.EnvVisitable {
	return c
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
		methodSpecs = append(methodSpecs, ast.MethodSpecification{
			MethodName: spec.MethodName,
			MethodSignature: ast.MethodSignature{
				MethodParameters: t.identifyTypeParamsInMethodParams(spec),
				ReturnType:       t.identifyTypeParams(spec.MethodSignature.ReturnType).(ast.Type),
			},
		})
	}

	return ast.InterfaceTypeLiteral{
		MethodSpecifications: methodSpecs,
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
