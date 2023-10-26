package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
)

func (t typeCheckingVisitor) typeDeclarationOf(typeName ast.TypeName) ast.TypeDeclaration {
	if typeName == intTypeName {
		return ast.TypeDeclaration{TypeName: intTypeName}
	}
	for _, decl := range t.declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl && typeDecl.TypeName == typeName {
			return typeDecl
		}
	}
	panic(fmt.Sprintf("could not find declaration for typename %q", typeName))
}

func (t typeCheckingVisitor) VisitTypeDeclaration(tdecl ast.TypeDeclaration) error {
	if err := t.typeCheckTypeDeclaration(tdecl); err != nil {
		return fmt.Errorf("type %q: %w", tdecl.TypeName, err)
	}
	return nil
}

func (t typeCheckingVisitor) typeCheckTypeDeclaration(tdecl ast.TypeDeclaration) error {
	if err := t.typeCheckTypeParams(tdecl.TypeParameters); err != nil {
		return err
	}
	return t.newTypeEnvTypeCheckingVisitor(tdecl.TypeParameters).typeCheck(tdecl.TypeLiteral)
}

func (t typeCheckingVisitor) typeCheckTypeParams(params []ast.TypeParameterConstraint) error {
	if err := checkDistinctTypeParameterNames(params); err != nil {
		return fmt.Errorf("type parameter %w", err)
	}
	envChecker := t.newTypeEnvTypeCheckingVisitor(params)
	for _, param := range params {
		if err := envChecker.typeCheck(param.Bound); err != nil {
			return fmt.Errorf("illegal bound of type parameter %q: %w", param.TypeParameter, err)
		}
		if !envChecker.isValidBoundType(param.Bound) {
			return fmt.Errorf(`cannot use type %q as bound: bound must be interface type or the keyword "const"`, param.Bound)
		}
	}
	return nil
}

func checkDistinctTypeParameterNames(params []ast.TypeParameterConstraint) error {
	paramNames := []ast.TypeParameter{}
	for _, param := range params {
		paramNames = append(paramNames, param.TypeParameter)
	}
	return auxiliary.Distinct(paramNames)
}

type typeEnvTypeCheckingVisitor struct {
	typeCheckingVisitor
	typeEnv map[ast.TypeParameter]ast.Bound
}

func (t typeCheckingVisitor) newTypeEnvTypeCheckingVisitor(typeParams []ast.TypeParameterConstraint) typeEnvTypeCheckingVisitor {
	env := make(map[ast.TypeParameter]ast.Bound)
	for _, param := range typeParams {
		env[param.TypeParameter] = param.Bound
	}
	return typeEnvTypeCheckingVisitor{
		typeCheckingVisitor: t,
		typeEnv:             env,
	}
}

func (t typeEnvTypeCheckingVisitor) typeOf(variableEnv map[string]ast.Type, expr ast.Expression) (ast.Type, error) {
	return t.typeCheckingVisitor.typeOf(t.typeEnv, variableEnv, expr)
}

func (t typeEnvTypeCheckingVisitor) typeCheck(v ast.EnvVisitable) error {
	return t.identifyTypeParams(v).AcceptEnvVisitor(t)
}

// since there is no way to syntactically distinguish between a type parameter
// and a named type with 0 type parameters, before type checking, it is
// necessary to identify all type parameters
func (t typeEnvTypeCheckingVisitor) identifyTypeParams(v ast.EnvVisitable) ast.EnvVisitable {
	return typeParamIdentifier{t}.identifyTypeParams(v)
}

func (t typeEnvTypeCheckingVisitor) VisitConstType(c ast.ConstType) error {
	return nil
}

func (t typeEnvTypeCheckingVisitor) VisitTypeParameter(typeParam ast.TypeParameter) error {
	if _, inEnv := t.typeEnv[typeParam]; !inEnv {
		panic("misidentified type parameter")
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) VisitNamedType(n ast.NamedType) error {
	for _, arg := range n.TypeArguments {
		err := t.typeCheck(arg)
		if err != nil {
			return fmt.Errorf("type %q badly instantiated: %w", n.TypeName, err)
		}
	}
	if !(slices.Contains(typeDeclarationNames(t.declarations), n.TypeName) || n.TypeName == intTypeName) {
		return fmt.Errorf("type name not declared: %q", n.TypeName)
	}
	if err := t.makeTypeSubstitutionsCheckingBounds(n); err != nil {
		return fmt.Errorf("type %q badly instantiated: %w", n.TypeName, err)
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) makeTypeSubstitutionsCheckingBounds(n ast.NamedType) error {
	decl := t.typeDeclarationOf(n.TypeName)
	substitutor, err := newTypeParamSubstituter(n.TypeArguments, decl.TypeParameters)
	if err != nil {
		return err
	}

	for _, typeParam := range decl.TypeParameters {
		typeArg := substitutor.substituteTypeParams(typeParam.TypeParameter).(ast.Type)

		if err := t.checkConstEquivalence(typeArg, typeParam.Bound); err != nil {
			return err
		}
		if err := t.checkIsSubtypeOf(typeArg, typeParam.Bound); err != nil {
			return err
		}
	}
	return nil
}

func makeTypeSubstitutions(typeArguments []ast.Type, typeParams []ast.TypeParameterConstraint) (map[ast.TypeParameter]ast.Type, error) {
	if len(typeArguments) != len(typeParams) {
		return nil, fmt.Errorf("expected %d type arguments but got %d", len(typeParams), len(typeArguments))
	}
	typeSubstitutions := make(map[ast.TypeParameter]ast.Type)
	for i, typeParam := range typeParams {
		typeSubstitutions[typeParam.TypeParameter] = typeArguments[i]
	}
	return typeSubstitutions, nil
}

func (t typeEnvTypeCheckingVisitor) checkConstEquivalence(typeArg ast.Type, typeParamBound ast.Bound) error {
	if t.isConst(typeParamBound) && !t.isConst(typeArg) {
		return fmt.Errorf("type %q cannot be used as const type argument", typeArg)
	}
	if !t.isConst(typeParamBound) && t.isConst(typeArg) {
		return fmt.Errorf("type %q cannot be used as non-const type argument", typeArg)
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) isValidBoundType(bound ast.Bound) bool {
	switch t.identifyTypeParams(bound).(type) {
	case ast.ConstType:
		return true
	case ast.NamedType:
		tdecl := t.typeDeclarationOf(bound.(ast.NamedType).TypeName)
		_, isInterfaceType := tdecl.TypeLiteral.(ast.InterfaceTypeLiteral)
		return isInterfaceType
	default:
		return false
	}
}
