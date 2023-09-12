package ast

import (
	"fmt"
)

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

	if isArraySetMethod(declarations, receiver.TypeName, m.MethodName) {
		if len(m.Arguments) != 2 {
			return nil, fmt.Errorf(
				`expected 2 arguments in call to "%s.%s", but got %d`,
				receiver.TypeName, m.MethodName, len(m.Arguments),
			)
		}

		intIndex, isIntIndex := m.Arguments[0].(IntegerLiteral)
		if !isIntIndex {
			return nil, fmt.Errorf(
				`non-integer index %q in array set method call: %s.%s`,
				m.Arguments[0], receiver.TypeName, m.MethodName,
			)
		}
		i := intIndex.IntValue

		withinBounds, err := inIndexBounds(declarations, receiver.TypeName, i)
		if err != nil {
			return nil, err
		}
		if !withinBounds {
			return nil, fmt.Errorf("array set index %d out of bounds for array of type %q", i, receiver.TypeName)
		}

		reducedArrayValues := make([]Expression, len(receiver.Values))
		copy(reducedArrayValues, receiver.Values)

		reducedArrayValues[i] = m.Arguments[1]

		return ValueLiteral{
			TypeName: receiver.TypeName,
			Values:   reducedArrayValues,
		}, nil
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

	for i, param := range parameterNames[1:] {
		arguments[param] = m.Arguments[i]
	}
	res, err := methodBody.bind(arguments)
	if err != nil {
		return nil, fmt.Errorf(`cannot call method "%s.%s": %w`, receiver.TypeName, m.MethodName, err)
	}
	return res, nil
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
