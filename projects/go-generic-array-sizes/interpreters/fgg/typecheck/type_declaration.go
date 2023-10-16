package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
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
