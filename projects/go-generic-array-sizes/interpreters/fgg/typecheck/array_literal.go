package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeCheckingVisitor) elementType(typeName ast.TypeName) ast.Type {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementType
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral, namedType ast.NamedType) error {
	expectedLen, hasDefinedLen := t.len(namedType).(ast.IntegerLiteral)
	if !hasDefinedLen {
		return fmt.Errorf("cannot create array literal of type %q with non-concrete length %q",
			v.Type, t.len(namedType))
	}
	if len(v.Values) != expectedLen.IntValue {
		return fmt.Errorf("expected %d values in array literal of type %q but got %d",
			expectedLen.IntValue, v.Type, len(v.Values))
	}

	elemType, err := t.substitutedElementType(namedType)
	if err != nil {
		return err
	}

	for _, val := range v.Values {
		valType, err := t.typeOf(val)
		if err != nil {
			return err
		}
		if err := t.checkIsSubtypeOf(valType, elemType); err != nil {
			return fmt.Errorf("cannot use %q as element of array of type %q: %w", val, namedType, err)
		}
	}
	return nil
}

func (t typeVisitor) substitutedElementType(namedType ast.NamedType) (ast.Type, error) {
	typeParams := t.typeParams(namedType.TypeName)
	envChecker := t.newTypeEnvTypeCheckingVisitor(typeParams)

	elemType := t.elementType(namedType.TypeName)
	elemTypeParam, isElemTypeParam := envChecker.identifyTypeParams(elemType).(ast.TypeParameter)
	if isElemTypeParam {
		typeSubstitutions, err := makeTypeSubstitutions(namedType.TypeArguments, typeParams)
		return typeSubstitutions[elemTypeParam], err
	}
	return elemType, nil
}

func (t typeVisitor) typeParams(typeName ast.TypeName) []ast.TypeParameterConstraint {
	return t.typeDeclarationOf(typeName).TypeParameters
}

func (t typeVisitor) typeCheckTypeArgument(typeArg ast.Type, param ast.TypeParameterConstraint) error {
	if err := t.checkConstEquivalence(typeArg, param.Bound); err != nil {
		return err
	}
	return t.checkIsSubtypeOf(typeArg, param.Bound)
}

func (t typeCheckingVisitor) len(namedType ast.NamedType) ast.Type {
	typeDecl := t.typeDeclarationOf(namedType.TypeName)
	envChecker := t.newTypeEnvTypeCheckingVisitor(typeDecl.TypeParameters)

	lenType := envChecker.identifyTypeParams(typeDecl.TypeLiteral.(ast.ArrayTypeLiteral).Length)
	if intLenType, isIntLenType := lenType.(ast.IntegerLiteral); isIntLenType {
		return intLenType
	}
	paramIndex := slices.IndexFunc(typeDecl.TypeParameters, func(p ast.TypeParameterConstraint) bool {
		return p.TypeParameter == lenType
	})
	return namedType.TypeArguments[paramIndex]
}