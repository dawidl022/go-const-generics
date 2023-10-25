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
	receiverType, typeParams, err := t.getReceiverType(m.MethodReceiver)
	if err != nil {
		return err
	}
	envChecker := t.newTypeEnvTypeCheckingVisitor(typeParams)
	err = envChecker.typeCheck(m.MethodSpecification)
	if err != nil {
		return err
	}
	expressionType, err := envChecker.typeOf(makeMethodVariableEnv(m, receiverType), m.ReturnExpression)
	if err != nil {
		return err
	}
	err = envChecker.checkIsSubtypeOf(expressionType, m.MethodSpecification.MethodSignature.ReturnType)
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

func (t typeCheckingVisitor) getReceiverType(receiver ast.MethodReceiver) (ast.Type, []ast.TypeParameterConstraint, error) {
	for _, decl := range t.declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if isTypeDecl && typeDecl.TypeName == receiver.TypeName {
			typeArgs := []ast.Type{}
			if len(receiver.TypeParameters) != len(typeDecl.TypeParameters) {
				return nil, nil, fmt.Errorf("expected %d type parameters on receiver but got %d",
					len(typeDecl.TypeParameters), len(receiver.TypeParameters))
			}
			for i, param := range receiver.TypeParameters {
				if param != typeDecl.TypeParameters[i].TypeParameter {
					return nil, nil, fmt.Errorf(
						"receiver type parameter name %q does not match type declaration parameter name %q",
						param, typeDecl.TypeParameters[i].TypeParameter)
				}
				typeArgs = append(typeArgs, param)
			}
			return ast.NamedType{
				TypeName:      receiver.TypeName,
				TypeArguments: typeArgs,
			}, typeDecl.TypeParameters, nil
		}
	}
	return nil, nil, fmt.Errorf("receiver type name not declared: %q", receiver.TypeName)
}

func makeMethodVariableEnv(m ast.MethodDeclaration, receiverType ast.Type) map[string]ast.Type {
	env := map[string]ast.Type{m.MethodReceiver.ParameterName: receiverType}
	for _, param := range m.MethodSpecification.MethodSignature.MethodParameters {
		env[param.ParameterName] = param.Type
	}
	return env
}

func (t typeEnvTypeCheckingVisitor) checkParameterTypes(parameters []ast.MethodParameter) error {
	for _, param := range parameters {
		err := t.typeCheck(param.Type)
		if err != nil {
			return fmt.Errorf("parameter %q: %w", param.ParameterName, err)
		}
		if t.isConst(param.Type) {
			return fmt.Errorf("parameter %q: method parameter cannot be of const type %q",
				param.ParameterName, param.Type)
		}
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) checkReturnType(m ast.MethodDeclaration) error {
	err := t.typeCheck(m.MethodSpecification.MethodSignature.ReturnType)
	if err != nil {
		return fmt.Errorf("return %w", err)
	}
	return nil
}
