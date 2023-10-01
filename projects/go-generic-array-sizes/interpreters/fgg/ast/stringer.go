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
	str := v.Type.String() + "{"

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

func (t TypeParameter) String() string {
	return string(t)
}

func (n NamedType) String() string {
	s := n.TypeName.String()

	if len(n.TypeArguments) > 0 {
		s += "["

		for i, arg := range n.TypeArguments {
			if i > 0 {
				s += ", "
			}
			s += arg.String()
		}
		s += "]"
	}
	return s
}

func (t TypeName) String() string {
	return string(t)
}
