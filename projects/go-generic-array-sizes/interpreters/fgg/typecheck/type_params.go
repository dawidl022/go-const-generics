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
