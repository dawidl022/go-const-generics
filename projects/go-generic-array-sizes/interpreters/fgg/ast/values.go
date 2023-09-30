package ast

func (i IntegerLiteral) IsValue() bool {
	return true
}

func (v Variable) IsValue() bool {
	return false
}

func (m MethodCall) IsValue() bool {
	//TODO implement me
	panic("implement me")
}

func (v ValueLiteral) IsValue() bool {
	for _, val := range v.Values {
		if !val.IsValue() {
			return false
		}
	}
	return true
}

func (s Select) IsValue() bool {
	return false
}

func (a ArrayIndex) IsValue() bool {
	return false
}
