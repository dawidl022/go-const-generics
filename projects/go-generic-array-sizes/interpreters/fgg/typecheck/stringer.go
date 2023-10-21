package typecheck

import "fmt"

func (m MethodSet) String() string {
	s := ""

	for i, method := range m {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%q", method)
	}

	return s
}
