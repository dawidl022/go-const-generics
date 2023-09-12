package ast

import (
	"fmt"
)

func (m MethodCall) Reduce(declarations []Declaration) (Expression, error) {
	if m.Receiver.Value() == nil {
		return m.withReducedReceiver(declarations)
	}
	for i, arg := range m.Arguments {
		if arg.Value() == nil {
			return m.withReducedArg(declarations, i)
		}
	}

	receiver, isValueLitReceiver := m.Receiver.(ValueLiteral)
	if !isValueLitReceiver {
		return nil, fmt.Errorf("cannot call method %q on primitive value %s", m.MethodName, m.Receiver)
	}

	if isArraySetMethod(declarations, receiver.TypeName, m.MethodName) {
		return m.reduceArraySet(declarations, receiver)
	}

	return m.reduceCall(declarations, receiver)
}

func (m MethodCall) withReducedReceiver(declarations []Declaration) (MethodCall, error) {
	reducedReceiver, err := m.Receiver.Reduce(declarations)
	return MethodCall{
		Receiver:   reducedReceiver,
		MethodName: m.MethodName,
		Arguments:  m.Arguments,
	}, err
}

func (m MethodCall) withReducedArg(declarations []Declaration, i int) (MethodCall, error) {
	reducedArgs := make([]Expression, len(m.Arguments))
	copy(reducedArgs, m.Arguments)

	reducedArg, err := m.Arguments[i].Reduce(declarations)
	reducedArgs[i] = reducedArg

	return MethodCall{
		Receiver:   m.Receiver,
		MethodName: m.MethodName,
		Arguments:  reducedArgs,
	}, err
}

func isArraySetMethod(declarations []Declaration, typeName, methodName string) bool {
	for _, decl := range declarations {
		methodDecl, isArraySetMethodDecl := decl.(ArraySetMethodDeclaration)
		if isArraySetMethodDecl && matchesMethod(methodDecl, typeName, methodName) {
			return true
		}
	}
	return false
}

func (m MethodCall) reduceArraySet(declarations []Declaration, receiver ValueLiteral) (Expression, error) {
	index, err := m.getArraySetIndex(declarations, receiver)
	if err != nil {
		return nil, err
	}
	return m.reduceArrayValues(receiver, index)
}

func (m MethodCall) getArraySetIndex(declarations []Declaration, receiver ValueLiteral) (IntegerLiteral, error) {
	if len(m.Arguments) != 2 {
		return IntegerLiteral{}, fmt.Errorf(
			`expected 2 arguments in call to "%s.%s", but got %d`,
			receiver.TypeName, m.MethodName, len(m.Arguments),
		)
	}

	index, isIntIndex := m.Arguments[0].(IntegerLiteral)
	if !isIntIndex {
		return IntegerLiteral{}, fmt.Errorf(
			`non-integer index %q in array set method call: %s.%s`,
			m.Arguments[0], receiver.TypeName, m.MethodName,
		)
	}
	withinBounds, err := inIndexBounds(declarations, receiver.TypeName, index.IntValue)
	if err != nil {
		return IntegerLiteral{}, err
	}
	if !withinBounds {
		return IntegerLiteral{}, fmt.Errorf(
			"array set index %d out of bounds for array of type %q",
			index.IntValue, receiver.TypeName,
		)
	}
	return index, nil
}

func (m MethodCall) reduceArrayValues(receiver ValueLiteral, index IntegerLiteral) (Expression, error) {
	reducedArrayValues := make([]Expression, len(receiver.Values))
	copy(reducedArrayValues, receiver.Values)

	reducedArrayValues[index.IntValue] = m.Arguments[1]

	return ValueLiteral{
		TypeName: receiver.TypeName,
		Values:   reducedArrayValues,
	}, nil
}

func (m MethodCall) reduceCall(declarations []Declaration, receiver ValueLiteral) (Expression, error) {
	parameterNames, methodBody, err := body(declarations, receiver.TypeName, m.MethodName)
	if err != nil {
		return nil, err
	}

	if len(m.Arguments) != len(parameterNames[1:]) {
		return nil, fmt.Errorf(
			`expected %d argument(s) in call to "%s.%s", but got %d`,
			len(parameterNames[1:]), receiver.TypeName, m.MethodName, len(m.Arguments),
		)
	}

	res, err := m.bindArguments(methodBody, parameterNames)
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

func (m MethodCall) bindArguments(methodBody Expression, parameterNames []string) (Expression, error) {
	arguments := map[string]Expression{parameterNames[0]: m.Receiver}
	for i, param := range parameterNames[1:] {
		arguments[param] = m.Arguments[i]
	}
	return methodBody.bind(arguments)
}

type CallableDeclaration interface {
	GetMethodReceiver() MethodParameter
	GetMethodName() string
}

func matchesMethod(methodDeclaration CallableDeclaration, typeName, methodName string) bool {
	return methodDeclaration.GetMethodReceiver().TypeName == typeName &&
		methodDeclaration.GetMethodName() == methodName
}

func (m MethodCall) Value() Value {
	return nil
}
