package reduction

import (
	"errors"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type ProgramReducer struct {
	observers []Observer
}

func NewProgramReducer(observers []Observer) *ProgramReducer {
	return &ProgramReducer{observers: observers}
}

type Observer interface {
	Notify(expression ast.Expression)
}

func (p ProgramReducer) ReduceToValue(program ast.Program) (ast.Value, error) {
	seenTerms := make(map[string]struct{})

	for program.Expression.Value() == nil {
		if _, alreadySeen := seenTerms[program.Expression.String()]; alreadySeen {
			return nil, errors.New("infinite loop detected")
		}
		seenTerms[program.Expression.String()] = struct{}{}

		var err error
		program, err = program.Reduce()

		if err != nil {
			return nil, err
		}
		p.notifyObservers(program.Expression)
	}
	return program.Expression.Value(), nil
}

func (p ProgramReducer) notifyObservers(expression ast.Expression) {
	for _, o := range p.observers {
		o.Notify(expression)
	}
}
