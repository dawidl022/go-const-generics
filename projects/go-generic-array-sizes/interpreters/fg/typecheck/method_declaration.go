package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) VisitMethodDeclaration(m ast.MethodDeclaration) error {
	err := t.typeCheckMethodDeclaration(m)
	if err != nil {
		return fmt.Errorf(`method "%s.%s": %w`, m.MethodReceiver.TypeName, m.GetMethodName(), err)
	}
	return nil
}

func (t typeCheckingVisitor) typeCheckMethodDeclaration(m ast.MethodDeclaration) error {
	err := t.checkDistinctParameterNames(m)
	if err != nil {
		return err
	}
	err = t.TypeCheck(m.MethodReceiver.TypeName)
	if err != nil {
		return fmt.Errorf("receiver %w", err)
	}
	err = t.checkParameterTypeNames(m.MethodSpecification.MethodSignature.MethodParameters)
	if err != nil {
		return err
	}
	err = t.checkReturnType(m)
	if err != nil {
		return err
	}
	expressionType, err := t.typeOf(makeMethodVariableEnv(m), m.ReturnExpression)
	if err != nil {
		return err
	}
	err = t.checkIsSubtypeOf(expressionType, m.MethodSpecification.MethodSignature.ReturnTypeName)
	if err != nil {
		return fmt.Errorf("return expression of %w", err)
	}
	return nil
}

func makeMethodVariableEnv(m ast.MethodDeclaration) map[string]ast.TypeName {
	env := map[string]ast.TypeName{m.MethodReceiver.ParameterName: m.MethodReceiver.TypeName}
	for _, param := range m.MethodSpecification.MethodSignature.MethodParameters {
		env[param.ParameterName] = param.TypeName
	}
	return env
}

func (t typeCheckingVisitor) checkDistinctParameterNames(m ast.MethodDeclaration) error {
	paramNames := []name{name(m.MethodReceiver.ParameterName)}
	for _, param := range m.MethodSpecification.MethodSignature.MethodParameters {
		paramNames = append(paramNames, name(param.ParameterName))
	}
	err := distinct(paramNames)
	if err != nil {
		return fmt.Errorf("parameter %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) checkParameterTypeNames(parameters []ast.MethodParameter) error {
	for _, param := range parameters {
		err := t.TypeCheck(param.TypeName)
		if err != nil {
			return fmt.Errorf("parameter %q: %w", param.ParameterName, err)
		}
	}
	return nil
}

func (t typeCheckingVisitor) checkReturnType(m ast.MethodDeclaration) error {
	err := t.TypeCheck(m.MethodSpecification.MethodSignature.ReturnTypeName)
	if err != nil {
		return fmt.Errorf("return %w", err)
	}
	return nil
}
