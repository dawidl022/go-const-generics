package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (r ReducingVisitor) VisitMethodCall(m ast.MethodCall) (ast.Expression, error) {
	if !m.Receiver.IsValue() {
		return r.methodCallWithReducedReceiver(m)
	}
	for i, arg := range m.Arguments {
		if !arg.IsValue() {
			return r.methodCallWithReducedArg(m, i)
		}
	}
	receiver, isValLitReceiver := m.Receiver.(ast.ValueLiteral)
	if !isValLitReceiver {
		return nil, fmt.Errorf("cannot call method %q on primitive value %s", m.MethodName, m.Receiver)
	}

	namedReceiverType, isNamedReceiverType := receiver.Type.(ast.NamedType)
	if !isNamedReceiverType {
		return nil, fmt.Errorf("type %q is not a valid value literal type", receiver.Type)
	}

	if r.isArraySetMethod(namedReceiverType, m.MethodName) {
		return r.reduceArraySetMethodCall(m, receiver, namedReceiverType)
	}

	return r.reduceMethodCall(m, namedReceiverType)
}

func (r ReducingVisitor) methodCallWithReducedReceiver(m ast.MethodCall) (ast.Expression, error) {
	reducedReceiver, err := r.Reduce(m.Receiver)
	return ast.MethodCall{
		Receiver:   reducedReceiver,
		MethodName: m.MethodName,
		Arguments:  m.Arguments,
	}, err
}

func (r ReducingVisitor) methodCallWithReducedArg(m ast.MethodCall, i int) (ast.Expression, error) {
	reducedArgs := make([]ast.Expression, len(m.Arguments))
	copy(reducedArgs, m.Arguments)

	reducedArg, err := r.Reduce(m.Arguments[i])
	reducedArgs[i] = reducedArg

	return ast.MethodCall{
		Receiver:   m.Receiver,
		MethodName: m.MethodName,
		Arguments:  reducedArgs,
	}, err
}

func (r ReducingVisitor) reduceMethodCall(m ast.MethodCall, namedReceiverType ast.NamedType) (ast.Expression, error) {
	parameterNames, methodBody, err := r.body(namedReceiverType, m.MethodName)
	if err != nil {
		return nil, err
	}

	if len(m.Arguments) != len(parameterNames[1:]) {
		return nil, fmt.Errorf(
			`expected %d argument(s) in call to "%s.%s", but got %d`,
			len(parameterNames[1:]), namedReceiverType, m.MethodName, len(m.Arguments),
		)
	}

	res, err := bindArguments(m, methodBody, parameterNames)
	if err != nil {
		return nil, fmt.Errorf(`cannot call method "%s.%s": %w`, namedReceiverType, m.MethodName, err)
	}
	return res, nil
}

func (r ReducingVisitor) body(receiverType ast.NamedType, methodName string) ([]string, ast.Expression, error) {
	for _, decl := range r.declarations {
		methodDecl, isMethodDecl := decl.(ast.MethodDeclaration)
		if isMethodDecl && matchesMethod(methodDecl, receiverType, methodName) {
			parameters := methodDecl.MethodSpecification.MethodSignature.MethodParameters
			parameterNames := make([]string, 0, len(parameters)+1)
			parameterNames = append(parameterNames, methodDecl.MethodReceiver.ParameterName)

			for _, param := range parameters {
				parameterNames = append(parameterNames, param.ParameterName)
			}
			return parameterNames, methodDecl.ReturnExpression, nil
		}
	}
	return nil, nil, fmt.Errorf("undeclared method %q on type %q", methodName, receiverType)
}

type CallableDeclaration interface {
	GetMethodReceiver() ast.MethodReceiver
	GetMethodName() string
}

func matchesMethod(methodDeclaration CallableDeclaration, receiverType ast.NamedType, methodName string) bool {
	return methodDeclaration.GetMethodReceiver().TypeName == receiverType.TypeName &&
		methodDeclaration.GetMethodName() == methodName
}

func bindArguments(m ast.MethodCall, methodBody ast.Expression, parameterNames []string) (ast.Expression, error) {
	arguments := map[string]ast.Expression{parameterNames[0]: m.Receiver}
	for i, param := range parameterNames[1:] {
		arguments[param] = m.Arguments[i]
	}
	return newBindingVisitor(arguments).bind(methodBody)
}
