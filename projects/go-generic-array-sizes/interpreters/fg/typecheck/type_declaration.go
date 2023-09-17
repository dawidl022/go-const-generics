package typecheck

import (
	"fmt"
	"slices"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) typeDeclarationOf(typeName ast.TypeName) ast.TypeDeclaration {
	for _, decl := range t.declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl && typeDecl.TypeName == typeName {
			return typeDecl
		}
	}
	panic("could not find declaration for typename")
}

func (t typeCheckingVisitor) VisitTypeName(typeName ast.TypeName) error {
	if slices.Contains(typeDeclarationNames(t.declarations), typeName) || typeName == "int" {
		return nil
	}
	return fmt.Errorf("type name not declared: %q", typeName)
}

func (t typeCheckingVisitor) VisitMethodSpecification(m ast.MethodSpecification) error {
	if err := checkDistinctParameterNames(m); err != nil {
		return fmt.Errorf("argument name %w", err)
	}
	for _, param := range m.MethodSignature.MethodParameters {
		if err := t.TypeCheck(param.TypeName); err != nil {
			return fmt.Errorf("argument %q %w", param.ParameterName, err)
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
	return distinct(paramNames)
}

type name string

func (n name) String() string {
	return string(n)
}

func (t typeCheckingVisitor) VisitArrayTypeLiteral(a ast.ArrayTypeLiteral) error {
	if a.Length < 0 {
		return fmt.Errorf("length cannot be less than 0")
	}
	if err := t.TypeCheck(a.ElementTypeName); err != nil {
		return fmt.Errorf("element %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) VisitStructTypeLiteral(s ast.StructTypeLiteral) error {
	if err := checkDistinctFiledNames(s); err != nil {
		return err
	}
	for _, field := range s.Fields {
		if err := t.TypeCheck(field.TypeName); err != nil {
			return fmt.Errorf("field %q %w", field.Name, err)
		}
	}
	return nil
}

func checkDistinctFiledNames(s ast.StructTypeLiteral) error {
	fieldNames := []name{}
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, name(field.Name))
	}
	if err := distinct(fieldNames); err != nil {
		return fmt.Errorf("field name %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) VisitInterfaceLiteral(i ast.InterfaceTypeLiteral) error {
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
	if err := distinct(methodNames); err != nil {
		return fmt.Errorf("method name %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) VisitTypeDeclaration(d ast.TypeDeclaration) error {
	err := t.TypeCheck(d.TypeLiteral)
	if err != nil {
		return fmt.Errorf("type %q: %w", d.TypeName, err)
	}
	return nil
}
