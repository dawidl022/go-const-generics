package typecheck

import (
	"fmt"
	"maps"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

// TODO test with generic types

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

func (s selfRefCheckingVisitor) checkSelfRefOfType(t ast.Type) error {
	switch typ := t.(type) {
	case ast.TypeParameter:
		return nil
	case ast.NamedType:
		typeName := typ.TypeName
		if typeName == intTypeName {
			return nil
		}
		if _, isSelfRef := s.refTypes[typeName]; isSelfRef {
			return fmt.Errorf("type %q", typeName)
		}
		// TODO need to instantiate the type literal with type args
		// how is eta substitution done in other places?
		substituter, err := newTypeParamSubstituter(typ.TypeArguments, s.declarations[typ.TypeName].TypeParameters)
		if err != nil {
			return err
		}
		substitutedLiteral := substituter.substituteTypeParams(s.declarations[typeName].TypeLiteral).(ast.TypeLiteral)
		err = s.withRefType(typeName).checkSelfRef(substitutedLiteral)
		if err != nil {
			return fmt.Errorf("type %q, which has: %w", typeName, err)
		}
		return nil
	}
	panic("unhandled type")
}

func (s selfRefCheckingVisitor) VisitStructTypeLiteral(st ast.StructTypeLiteral) error {
	for _, field := range st.Fields {
		if err := s.checkSelfRefOfType(field.Type); err != nil {
			return fmt.Errorf("field %q of %w", field.Name, err)
		}
	}
	return nil
}

func (s selfRefCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if err := s.checkSelfRefOfType(a.ElementType); err != nil {
		return fmt.Errorf("array element %w", err)
	}
	return nil
}

func (s selfRefCheckingVisitor) VisitInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) error {
	// interface type literals are allowed to refer to themselves
	return nil
}
