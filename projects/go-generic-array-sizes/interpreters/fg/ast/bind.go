package ast

import "fmt"

func (i IntegerLiteral) bind(variables map[string]Expression) (Expression, error) {
	return i, nil
}

func (v Variable) bind(variables map[string]Expression) (Expression, error) {
	if val, isBound := variables[v.Id]; isBound {
		return val, nil
	}
	return nil, fmt.Errorf("unbound variable %q", v.Id)
}

func (m MethodCall) bind(variables map[string]Expression) (Expression, error) {
	boundReceiver, err := m.Receiver.bind(variables)
	if err != nil {
		return nil, err
	}
	boundArgs := []Expression{}
	for _, arg := range m.Arguments {
		boundArg, err := arg.bind(variables)
		if err != nil {
			return nil, err
		}
		boundArgs = append(boundArgs, boundArg)
	}
	return MethodCall{
		Receiver:   boundReceiver,
		MethodName: m.MethodName,
		Arguments:  boundArgs,
	}, nil
}

func (v ValueLiteral) bind(variables map[string]Expression) (Expression, error) {
	boundValues := []Expression{}
	for _, val := range v.Values {
		boundVal, err := val.bind(variables)
		if err != nil {
			return nil, err
		}
		boundValues = append(boundValues, boundVal)
	}
	return ValueLiteral{TypeName: v.TypeName, Values: boundValues}, nil
}

func (s Select) bind(variables map[string]Expression) (Expression, error) {
	boundExpression, err := s.Receiver.bind(variables)
	return Select{FieldName: s.FieldName, Receiver: boundExpression}, err
}

func (a ArrayIndex) bind(variables map[string]Expression) (Expression, error) {
	boundReceiver, err := a.Receiver.bind(variables)
	if err != nil {
		return nil, err
	}
	boundIndex, err := a.Index.bind(variables)
	return ArrayIndex{Receiver: boundReceiver, Index: boundIndex}, err
}
