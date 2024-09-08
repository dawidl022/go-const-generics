package typecheck

import (
	"fmt"
	"maps"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

type selfRefTracker struct {
	declarations map[ast.TypeName]ast.TypeDeclaration
	refTypes     map[ast.TypeName]struct{}
}

type selfRefCheckingVisitor struct {
	selfRefTracker
}

func newRefCheckingVisitor(initialRefType ast.TypeName, declarations []ast.Declaration) selfRefCheckingVisitor {
	return selfRefCheckingVisitor{newSelfRefTracker(initialRefType, declarations)}
}

func newSelfRefTracker(initialRefType ast.TypeName, declarations []ast.Declaration) selfRefTracker {
	typeDecls := make(map[ast.TypeName]ast.TypeDeclaration)

	for _, decl := range declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl {
			typeDecls[typeDecl.TypeName] = typeDecl
		}
	}

	return selfRefTracker{
		declarations: typeDecls,
		refTypes:     map[ast.TypeName]struct{}{initialRefType: {}},
	}
}

func (s selfRefCheckingVisitor) withRefType(refType ast.TypeName) selfRefCheckingVisitor {
	newRefTypes := maps.Clone(s.refTypes)
	newRefTypes[refType] = struct{}{}

	return selfRefCheckingVisitor{
		selfRefTracker{
			declarations: s.declarations,
			refTypes:     newRefTypes,
		},
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

type typeParamSelfRefCheckingVisitor struct {
	selfRefTracker
}

func newTypeParamRefCheckingVisitor(
	initialRefType ast.TypeName, declarations []ast.Declaration,
) typeParamSelfRefCheckingVisitor {
	return typeParamSelfRefCheckingVisitor{newSelfRefTracker(initialRefType, declarations)}
}

func (t typeParamSelfRefCheckingVisitor) withRefType(refType ast.TypeName) typeParamSelfRefCheckingVisitor {
	newRefTypes := maps.Clone(t.refTypes)
	newRefTypes[refType] = struct{}{}

	return typeParamSelfRefCheckingVisitor{
		selfRefTracker{
			declarations: t.declarations,
			refTypes:     newRefTypes,
		},
	}
}

func (t typeParamSelfRefCheckingVisitor) checkSelfRef(visitable ast.TypeRefVisitable) error {
	return visitable.AcceptRefVisitor(t)
}

func (t typeParamSelfRefCheckingVisitor) checkSelfRefOfNamedType(n ast.TypeName) error {
	decl := t.declarations[n]
	for _, typeParamConstraint := range decl.TypeParameters {
		err := t.checkSelfRef(typeParamConstraint.Bound)
		if err != nil {
			return fmt.Errorf("bound of %q references %w", typeParamConstraint.TypeParameter, err)
		}
	}
	return nil
}

func (t typeParamSelfRefCheckingVisitor) VisitMapTypeNamedType(n ast.NamedType) error {
	if _, isSelfRef := t.refTypes[n.TypeName]; isSelfRef {
		return fmt.Errorf("%q", n.TypeName)
	}
	for _, typeArg := range n.TypeArguments {
		if err := t.checkSelfRef(typeArg); err != nil {
			return err
		}
	}
	err := t.withRefType(n.TypeName).checkSelfRefOfNamedType(n.TypeName)
	if err != nil {
		return fmt.Errorf("%q, where: %w", n.TypeName, err)
	}
	return nil
}

func (t typeParamSelfRefCheckingVisitor) VisitMapTypeConstType(c ast.ConstType) error {
	return nil
}

func (t typeParamSelfRefCheckingVisitor) VisitMapTypeTypeParameter(tp ast.TypeParameter) error {
	return nil
}

func (t typeParamSelfRefCheckingVisitor) VisitMapTypeIntegerLiteral(i ast.IntegerLiteral) error {
	return nil
}
