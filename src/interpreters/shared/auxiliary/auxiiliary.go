package auxiliary

import "fmt"

type ComparableStringer interface {
	comparable
	fmt.Stringer
}

func Distinct[T ComparableStringer](names []T) error {
	seenNames := make(map[T]struct{})

	for _, name := range names {
		if _, seen := seenNames[name]; seen {
			return fmt.Errorf("redeclared %q", name)
		}
		seenNames[name] = struct{}{}
	}
	return nil
}
