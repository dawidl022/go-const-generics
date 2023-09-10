package ast

import "fmt"

func (m MethodCall) Reduce(declarations []Declaration) (Expression, error) {
	receiver, isValueLitReceiver := m.Receiver.(ValueLiteral)
	if !isValueLitReceiver {
		return nil, fmt.Errorf("cannot call method %q on primitive value %s", m.MethodName, m.Receiver)
	}
	parameterNames, methodBody, err := body(declarations, receiver.TypeName, m.MethodName)
	if err != nil {
		return nil, err
	}
	allArguments := []Expression{m.Receiver}
	allArguments = append(allArguments, m.Arguments...)
	res, err := bind(methodBody, parameterNames, allArguments)
	if err != nil {
		return nil, fmt.Errorf(`cannot call method "%s.%s": %w`, receiver.TypeName, m.MethodName, err)
	}
	return res, nil
}

func bind(expr Expression, variableNames []string, variables []Expression) (Expression, error) {
	variable, isVariable := expr.(Variable)
	if isVariable {
		for i, varName := range variableNames {
			if variable.Id == varName {
				return variables[i], nil
			}
		}
		return nil, fmt.Errorf("unbound variable %q", variable.Id)
	}
	return expr, nil
}

func body(declarations []Declaration, typeName, methodName string) ([]string, Expression, error) {
	for _, decl := range declarations {
		methodDecl, isMethodDecl := decl.(MethodDeclaration)
		if isMethodDecl && matchesMethod(methodDecl, typeName, methodName) {
			parameters := methodDecl.MethodSpecification.MethodSignature.MethodParameters
			parameterNames := make([]string, 0, len(parameters)+1)
			parameterNames = append(parameterNames, methodDecl.MethodReceiver.ParameterName)

			for _, param := range parameters {
				parameterNames = append(parameterNames, param.ParameterName)
			}
			return parameterNames, methodDecl.ReturnExpression, nil
		}
	}
	return nil, nil, fmt.Errorf("undeclared method %q on type %q", methodName, typeName)
}

func matchesMethod(methodDeclaration MethodDeclaration, typeName, methodName string) bool {
	return methodDeclaration.MethodReceiver.TypeName == typeName &&
		methodDeclaration.MethodSpecification.MethodName == methodName
}

func (m MethodCall) Value() Value {
	//TODO implement me
	panic("implement me")
}
