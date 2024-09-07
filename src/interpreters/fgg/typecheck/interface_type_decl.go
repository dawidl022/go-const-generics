package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
)

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
			return fmt.Errorf("parameter %q: %w", param.ParameterName, err)
		}
		if t.isConst(param.Type) {
			return fmt.Errorf("parameter %q: const type %q cannot be used as parameter type",
				param.ParameterName, param.Type)
		}
	}
	if err := t.typeCheck(m.MethodSignature.ReturnType); err != nil {
		return fmt.Errorf("return %w", err)
	}
	if t.isConst(m.MethodSignature.ReturnType) {
		return fmt.Errorf("const type %q cannot be used as return type",
			m.MethodSignature.ReturnType)
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
