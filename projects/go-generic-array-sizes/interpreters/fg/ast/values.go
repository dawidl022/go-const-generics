package ast

import (
	"fmt"
	"strconv"
)

type Value interface {
	fmt.Stringer
	val()
}

func (i IntegerLiteral) String() string {
	return strconv.Itoa(i.Value)
}

func (i IntegerLiteral) val() {
}

func (v ValueLiteral) String() string {
	str := v.TypeName + "{"

	for i, val := range v.Values {
		if i > 0 {
			str += ", "
		}
		str += val.(Value).String()
	}

	str += "}"
	return str
}

func (v ValueLiteral) val() {
}
