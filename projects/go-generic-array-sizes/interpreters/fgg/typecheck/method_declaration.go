package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
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
	receiverType, err := t.getReceiverType(m.MethodReceiver.TypeName)
	if err != nil {
		return err
	}
	err = t.checkParameterTypes(m.MethodSpecification.MethodSignature.MethodParameters)
	if err != nil {
		return err
	}
	err = t.checkReturnType(m)
	if err != nil {
		return err
	}
	expressionType, err := t.typeOf(nil, makeMethodVariableEnv(m, receiverType), m.ReturnExpression) // TODO fill in typing environment
	if err != nil {
		return err
	}
	// TODO consolidate creation of excess EnvTypeCheckingVisitor
	err = t.newTypeEnvTypeCheckingVisitor(nil).checkIsSubtypeOf(expressionType, m.MethodSpecification.MethodSignature.ReturnType) // TODO fill in typing env
	if err != nil {
		return fmt.Errorf("return expression of %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) checkDistinctParameterNames(m ast.MethodDeclaration) error {
	paramNames := []name{name(m.MethodReceiver.ParameterName)}
	for _, param := range m.MethodSpecification.MethodSignature.MethodParameters {
		paramNames = append(paramNames, name(param.ParameterName))
	}
	err := auxiliary.Distinct(paramNames)
	if err != nil {
		return fmt.Errorf("parameter %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) getReceiverType(typeName ast.TypeName) (ast.Type, error) {
	for _, decl := range t.declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if isTypeDecl && typeDecl.TypeName == typeName {
			return ast.NamedType{
				TypeName:      typeName,
				TypeArguments: nil, // TODO fill in type arguments
			}, nil
		}
	}
	return nil, fmt.Errorf("receiver type name not declared: %q", typeName)
}

func makeMethodVariableEnv(m ast.MethodDeclaration, receiverType ast.Type) map[string]ast.Type {
	env := map[string]ast.Type{m.MethodReceiver.ParameterName: receiverType}
	for _, param := range m.MethodSpecification.MethodSignature.MethodParameters {
		env[param.ParameterName] = param.Type
	}
	return env
}

func (t typeCheckingVisitor) checkParameterTypes(parameters []ast.MethodParameter) error {
	for _, param := range parameters {
		err := t.newTypeEnvTypeCheckingVisitor(nil).typeCheck(param.Type) // TODO fill in env
		if err != nil {
			return fmt.Errorf("parameter %q: %w", param.ParameterName, err)
		}
	}
	return nil
}

func (t typeCheckingVisitor) checkReturnType(m ast.MethodDeclaration) error {
	err := t.newTypeEnvTypeCheckingVisitor(nil).typeCheck(m.MethodSpecification.MethodSignature.ReturnType) // TODO fill in env
	if err != nil {
		return fmt.Errorf("return %w", err)
	}
	return nil
}
