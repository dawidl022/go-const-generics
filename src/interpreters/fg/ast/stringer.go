package ast

import (
	"fmt"
	"strconv"
)

func (i IntegerLiteral) String() string {
	return strconv.Itoa(i.IntValue)
}

func (v Variable) String() string {
	return v.Id
}

func (m MethodCall) String() string {
	str := fmt.Sprintf("%s.%s(", m.Receiver, m.MethodName)

	for i, arg := range m.Arguments {
		if i > 0 {
			str += ", "
		}
		str += arg.String()
	}

	str += ")"
	return str
}

func (v ValueLiteral) String() string {
	str := v.TypeName.String() + "{"

	for i, val := range v.Values {
		if i > 0 {
			str += ", "
		}
		str += val.String()
	}

	str += "}"
	return str
}

func (s Select) String() string {
	return fmt.Sprintf("%s.%s", s.Receiver, s.FieldName)
}

func (a ArrayIndex) String() string {
	return fmt.Sprintf("%s[%s]", a.Receiver, a.Index)
}

func (t TypeName) String() string {
	return string(t)
}

func (m MethodSpecification) String() string {
	return fmt.Sprintf("%s%s", m.MethodName, m.MethodSignature)
}

func (m MethodSignature) String() string {
	s := "("

	for i, param := range m.MethodParameters {
		if i > 0 {
			s += ", "
		}
		s += param.String()
	}
	s += fmt.Sprintf(") %s", m.ReturnTypeName)
	return s
}

func (p MethodParameter) String() string {
	return fmt.Sprintf("%s %s", p.ParameterName, p.TypeName)
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.Left, a.Right)
}
