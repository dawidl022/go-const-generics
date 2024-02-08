package typecheck

import (
	"fmt"

	"golang.org/x/exp/maps"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type selfRefCheckingVisitor struct {
	declarations map[ast.TypeName]ast.TypeDeclaration
	refTypes     map[ast.TypeName]struct{}
}

func newRefCheckingVisitor(initialRefType ast.TypeName, declarations []ast.Declaration) selfRefCheckingVisitor {
	typeDecls := make(map[ast.TypeName]ast.TypeDeclaration)

	for _, decl := range declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl {
			typeDecls[typeDecl.TypeName] = typeDecl
		}
	}

	return selfRefCheckingVisitor{
		declarations: typeDecls,
		refTypes:     map[ast.TypeName]struct{}{initialRefType: {}},
	}
}

func (s selfRefCheckingVisitor) withRefType(refType ast.TypeName) selfRefCheckingVisitor {
	newRefTypes := maps.Clone(s.refTypes)
	newRefTypes[refType] = struct{}{}

	return selfRefCheckingVisitor{
		declarations: s.declarations,
		refTypes:     newRefTypes,
	}
}

func (s selfRefCheckingVisitor) checkSelfRef(r ast.RefVisitable) error {
	return r.AcceptRef(s)
}

func (s selfRefCheckingVisitor) checkSelfRefOfType(t ast.TypeName) error {
	if t == intTypeName {
		return nil
	}
	return s.withRefType(t).checkSelfRef(s.declarations[t].TypeLiteral)
}

func (s selfRefCheckingVisitor) VisitStructTypeLiteral(st ast.StructTypeLiteral) error {
	for _, field := range st.Fields {
		if _, isSelfRef := s.refTypes[field.TypeName]; isSelfRef {
			return fmt.Errorf("field %q of type %q", field.Name, field.TypeName)
		}
	}
	for _, field := range st.Fields {
		err := s.checkSelfRefOfType(field.TypeName)
		if err != nil {
			return fmt.Errorf("field %q of type %q, which has: %w", field.Name, field.TypeName, err)
		}
	}
	return nil
}

func (s selfRefCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if _, isSelfRef := s.refTypes[a.ElementTypeName]; isSelfRef {
		return fmt.Errorf("array element type %q", a.ElementTypeName)
	}
	err := s.checkSelfRefOfType(a.ElementTypeName)
	if err != nil {
		return fmt.Errorf("array element type %q, which has: %w",
			a.ElementTypeName, err)
	}
	return nil
}

func (s selfRefCheckingVisitor) VisitInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) error {
	// interface type literals are allowed to refer to themselves
	return nil
}
