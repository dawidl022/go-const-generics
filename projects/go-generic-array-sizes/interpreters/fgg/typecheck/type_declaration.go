package typecheck

import (
	"fmt"
	"slices"

	"golang.org/x/exp/maps"

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
	envChecker := t.newTypeEnvTypeCheckingVisitor(tdecl.TypeParameters)
	// TODO may be worth moving identification into separate struct so it only does one thing
	declWithIdentifiedTypeParams, err := envChecker.identifyTypeLiteralParams(tdecl.TypeLiteral)
	if err != nil {
		return nil
	}
	return envChecker.typeCheck(declWithIdentifiedTypeParams)
}

func (t typeCheckingVisitor) typeCheckTypeParams(params []ast.TypeParameterConstraint) error {
	// TODO
	return nil
}

type typeEnvTypeCheckingVisitor struct {
	typeCheckingVisitor
	typeEnv map[ast.TypeParameter]ast.Bound
}

func (t typeEnvTypeCheckingVisitor) VisitTypeParameter(typeParam ast.TypeParameter) error {
	if _, inEnv := t.typeEnv[typeParam]; !inEnv {
		panic("untested path")
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) VisitEnvNamedType(n ast.NamedType) (ast.Type, error) {
	// TODO what happens in case type param shadows type decl? Is this allowed in Go?
	typeParam := ast.TypeParameter(n.TypeName)
	if _, isTypeParam := t.typeEnv[typeParam]; isTypeParam {
		return typeParam, nil
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
	}, nil
}

func (t typeEnvTypeCheckingVisitor) VisitEnvArrayTypeLiteral(a ast.ArrayTypeLiteral) (ast.TypeLiteral, error) {
	lengthType, err := t.identifyTypeParams(a.Length)
	if err != nil {
		return nil, err
	}
	elementType, err := t.identifyTypeParams(a.ElementType)
	if err != nil {
		return nil, err
	}
	return ast.ArrayTypeLiteral{
		Length:      lengthType,
		ElementType: elementType,
	}, nil
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

func (t typeEnvTypeCheckingVisitor) typeCheck(v ast.EnvVisitable) error {
	return v.AcceptEnvVisitor(t)
}

func (t typeEnvTypeCheckingVisitor) identifyTypeParams(v ast.EnvTypeVisitable) (ast.Type, error) {
	return v.AcceptEnvTypeVisitor(t)
}

func (t typeEnvTypeCheckingVisitor) identifyTypeLiteralParams(v ast.EnvTypeLiteralVisitable) (ast.TypeLiteral, error) {
	return v.AcceptEnvTypeVisitor(t)
}

func (t typeEnvTypeCheckingVisitor) VisitNamedType(n ast.NamedType) error {
	for _, arg := range n.TypeArguments {
		err := t.typeCheck(arg)
		if err != nil {
			return fmt.Errorf("type %q badly instantiated: %w", n.TypeName, err)
		}
	}
	if slices.Contains(maps.Keys(t.typeEnv), ast.TypeParameter(n.TypeName)) {
		return nil
	}
	if !(slices.Contains(typeDeclarationNames(t.declarations), n.TypeName) || n.TypeName == intTypeName) {
		return fmt.Errorf("type name not declared: %q", n.TypeName)
	}
	_, err := t.makeTypeSubstitutionsCheckingBounds(n)
	return err
}

func (t typeEnvTypeCheckingVisitor) makeTypeSubstitutionsCheckingBounds(n ast.NamedType) (map[ast.TypeParameter]ast.Type, error) {
	decl := t.typeDeclarationOf(n.TypeName)
	typeSubstitutions := makeTypeSubstitutions(n, decl.TypeParameters)

	for _, typeParam := range decl.TypeParameters {
		err := t.checkIsSubtypeOf(typeSubstitutions[typeParam.TypeParameter], typeParam.Bound)
		if err != nil {
			return nil, err
		}
	}
	return typeSubstitutions, nil
}

func makeTypeSubstitutions(n ast.NamedType, typeParams []ast.TypeParameterConstraint) map[ast.TypeParameter]ast.Type {
	if len(n.TypeArguments) != len(typeParams) {
		panic("untested branch")
	}
	typeSubstitutions := make(map[ast.TypeParameter]ast.Type)
	for i, typeParam := range typeParams {
		typeSubstitutions[typeParam.TypeParameter] = n.TypeArguments[i]
	}
	return typeSubstitutions
}

func (t typeEnvTypeCheckingVisitor) VisitInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) error {
	if err := checkUniqueMethodNames(i); err != nil {
		return err
	}
	for _, spec := range i.MethodSpecifications {
		if err := t.typeCheck(spec); err != nil {
			return fmt.Errorf("method specification %q: %w", spec.MethodName, err)
		}
	}
	return nil
}

func checkUniqueMethodNames(i ast.InterfaceTypeLiteral) error {
	methodNames := []name{}
	for _, spec := range i.MethodSpecifications {
		methodNames = append(methodNames, name(spec.MethodName))
	}
	if err := auxiliary.Distinct(methodNames); err != nil {
		return fmt.Errorf("method name %w", err)
	}
	return nil
}

type name string

func (n name) String() string {
	return string(n)
}

func (t typeEnvTypeCheckingVisitor) VisitMethodSpecification(m ast.MethodSpecification) error {
	if err := checkDistinctParameterNames(m); err != nil {
		return fmt.Errorf("argument name %w", err)
	}
	for _, param := range m.MethodSignature.MethodParameters {
		if err := t.typeCheck(param.Type); err != nil {
			return fmt.Errorf("argument %q %w", param.ParameterName, err)
		}
	}
	if err := t.typeCheck(m.MethodSignature.ReturnType); err != nil {
		return fmt.Errorf("return %w", err)
	}
	// TODO check for non constant types
	return nil
}

func checkDistinctParameterNames(m ast.MethodSpecification) error {
	paramNames := []name{}
	for _, param := range m.MethodSignature.MethodParameters {
		paramNames = append(paramNames, name(param.ParameterName))
	}
	return auxiliary.Distinct(paramNames)
}

func (t typeEnvTypeCheckingVisitor) VisitStructTypeLiteral(s ast.StructTypeLiteral) error {
	if err := checkDistinctFieldNames(s); err != nil {
		return err
	}
	for _, field := range s.Fields {
		if err := t.typeCheck(field.Type); err != nil {
			return fmt.Errorf("field %q %w", field.Name, err)
		}
	}
	return nil
}

func checkDistinctFieldNames(s ast.StructTypeLiteral) error {
	fieldNames := []name{}
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, name(field.Name))
	}
	if err := auxiliary.Distinct(fieldNames); err != nil {
		return fmt.Errorf("field name %w", err)
	}
	return nil
}
