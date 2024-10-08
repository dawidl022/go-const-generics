package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/shared/auxiliary"
)

func (t TypeCheckingVisitor) typeDeclarationOf(typeName ast.TypeName) ast.TypeDeclaration {
	for _, decl := range t.declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl && typeDecl.TypeName == typeName {
			return typeDecl
		}
	}
	panic(fmt.Sprintf("could not find declaration for typename %q", typeName))
}

func (t TypeCheckingVisitor) VisitTypeName(typeName ast.TypeName) error {
	if slices.Contains(typeDeclarationNames(t.declarations), typeName) || typeName == intTypeName {
		return nil
	}
	return fmt.Errorf("type name not declared: %q", typeName)
}

func (t TypeCheckingVisitor) VisitMethodSpecification(m ast.MethodSpecification) error {
	if err := checkDistinctParameterNames(m); err != nil {
		return fmt.Errorf("argument name %w", err)
	}
	for _, param := range m.MethodSignature.MethodParameters {
		if err := t.TypeCheck(param.TypeName); err != nil {
			return fmt.Errorf("parameter %q: %w", param.ParameterName, err)
		}
	}
	if err := t.TypeCheck(m.MethodSignature.ReturnTypeName); err != nil {
		return fmt.Errorf("return %w", err)
	}
	return nil
}

func checkDistinctParameterNames(m ast.MethodSpecification) error {
	paramNames := []name{}
	for _, param := range m.MethodSignature.MethodParameters {
		paramNames = append(paramNames, name(param.ParameterName))
	}
	return auxiliary.Distinct(paramNames)
}

type name string

func (n name) String() string {
	return string(n)
}

func (t TypeCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if a.Length < 0 {
		return fmt.Errorf("length cannot be less than 0")
	}
	if err := t.TypeCheck(a.ElementTypeName); err != nil {
		return fmt.Errorf("element %w", err)
	}
	return nil
}

func (t TypeCheckingVisitor) VisitStructTypeLiteral(s ast.StructTypeLiteral) error {
	if err := checkDistinctFieldNames(s); err != nil {
		return err
	}
	for _, field := range s.Fields {
		if err := t.TypeCheck(field.TypeName); err != nil {
			return fmt.Errorf("field %q: %w", field.Name, err)
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

func (t TypeCheckingVisitor) VisitInterfaceLiteral(i ast.InterfaceTypeLiteral) error {
	if err := checkUniqueMethodNames(i); err != nil {
		return err
	}
	for _, spec := range i.MethodSpecifications {
		if err := t.TypeCheck(spec); err != nil {
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

func (t TypeCheckingVisitor) VisitTypeDeclaration(d ast.TypeDeclaration) error {
	err := t.TypeCheck(d.TypeLiteral)
	if err != nil {
		return fmt.Errorf("type %q: %w", d.TypeName, err)
	}
	err = newRefCheckingVisitor(d.TypeName, t.declarations).checkSelfRef(d.TypeLiteral)
	if err != nil {
		return fmt.Errorf("type %q: circular reference: %w", d.TypeName, err)
	}
	return nil
}
