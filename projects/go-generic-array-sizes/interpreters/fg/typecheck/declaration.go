package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) VisitTypeDeclaration(d ast.TypeDeclaration) error {
	err := t.TypeCheck(d.TypeLiteral)
	if err != nil {
		return fmt.Errorf("type %q: %w", d.TypeName, err)
	}
	return nil
}

func (t typeCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if a.Length < 0 {
		panic("untested branch")
	}
	if err := t.TypeCheck(a.ElementTypeName); err != nil {
		return fmt.Errorf("element %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) VisitStructTypeLiteral(s ast.StructTypeLiteral) error {
	if err := distinct(s.Fields); err != nil {
		panic("untested branch")
	}
	for _, f := range s.Fields {
		if err := t.TypeCheck(f.TypeName); err != nil {
			panic("untested branch")
		}
	}
	return nil
}

func (t typeCheckingVisitor) VisitTypeName(typeName ast.TypeName) error {
	if slices.Contains(typeDeclarationNames(t.declarations), typeName) || typeName == "int" {
		return nil
	}
	return fmt.Errorf("type name not ok: %q", typeName)
}
