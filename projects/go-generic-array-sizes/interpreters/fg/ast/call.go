package ast

import "fmt"

func (m MethodCall) Reduce(declarations []Declaration) (Expression, error) {
	if m.Receiver.Value() == nil {
		reducedReceiver, err := m.Receiver.Reduce(declarations)
		return MethodCall{
			Receiver:   reducedReceiver,
			MethodName: m.MethodName,
			Arguments:  m.Arguments,
		}, err
	}
	receiver, isValueLitReceiver := m.Receiver.(ValueLiteral)
	if !isValueLitReceiver {
		return nil, fmt.Errorf("cannot call method %q on primitive value %s", m.MethodName, m.Receiver)
	}
	parameterNames, methodBody, err := body(declarations, receiver.TypeName, m.MethodName)
	if err != nil {
		return nil, err
	}
	arguments := map[string]Expression{parameterNames[0]: m.Receiver}

	if len(m.Arguments) != len(parameterNames[1:]) {
		return nil, fmt.Errorf(
			`expected %d argument(s) in call to "%s.%s", but got %d`,
			len(parameterNames[1:]), receiver.TypeName, m.MethodName, len(m.Arguments),
		)
	}

	reducedArgs := make([]Expression, len(m.Arguments))
	copy(reducedArgs, m.Arguments)
	for i, arg := range m.Arguments {
		if arg.Value() == nil {
			reducedArg, err := arg.Reduce(declarations)
			reducedArgs[i] = reducedArg

			return MethodCall{
				Receiver:   m.Receiver,
				MethodName: m.MethodName,
				Arguments:  reducedArgs,
			}, err
		}
	}

	for i, param := range parameterNames[1:] {
		arguments[param] = m.Arguments[i]
	}
	res, err := methodBody.bind(arguments)
	if err != nil {
		return nil, fmt.Errorf(`cannot call method "%s.%s": %w`, receiver.TypeName, m.MethodName, err)
	}
	return res, nil
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
	return nil
}
