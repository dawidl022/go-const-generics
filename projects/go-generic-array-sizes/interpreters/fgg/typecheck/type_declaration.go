package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
)

func (t typeCheckingVisitor) VisitTypeDeclaration(tdecl ast.TypeDeclaration) error {
	if err := t.typeCheckDeclaration(tdecl); err != nil {
		return fmt.Errorf("type %q: %w", tdecl.TypeName, err)
	}
	return nil
}

func (t typeCheckingVisitor) typeCheckDeclaration(tdecl ast.TypeDeclaration) error {
	if err := t.typeCheckTypeParams(tdecl.TypeParameters); err != nil {
		return err
	}
	return t.newTypeEnvTypeCheckingVisitor(tdecl.TypeParameters).typeCheck(tdecl.TypeLiteral)
}

func (t typeCheckingVisitor) typeCheckTypeParams(params []ast.TypeParameterConstraint) error {
	return nil
}

type typeEnvTypeCheckingVisitor struct {
	declarations []ast.Declaration
	typeEnv      map[ast.TypeParameter]ast.Bound
}

func (t typeCheckingVisitor) newTypeEnvTypeCheckingVisitor(typeParams []ast.TypeParameterConstraint) typeEnvTypeCheckingVisitor {
	env := make(map[ast.TypeParameter]ast.Bound)
	for _, param := range typeParams {
		env[param.TypeParameter] = param.Bound
	}
	return typeEnvTypeCheckingVisitor{
		declarations: t.declarations,
		typeEnv:      env,
	}
}

func (t typeEnvTypeCheckingVisitor) typeCheck(v ast.EnvVisitable) error {
	return v.AcceptEnvVisitor(t)
}

func (t typeEnvTypeCheckingVisitor) AcceptArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	// TODO might need to t.typeCheck(a.Length) (not currently in formal rules)
	if err := t.typeCheck(a.ElementType); err != nil {
		return fmt.Errorf("element %w", err)
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) VisitNamedType(n ast.NamedType) error {
	// TODO type check each type argument
	// TODO check type arguments satisfy parameter bounds
	if !(slices.Contains(typeDeclarationNames(t.declarations), n.TypeName) || n.TypeName == intTypeName) {
		return fmt.Errorf("type name not declared: %q", n.TypeName)
	}
	return nil
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
