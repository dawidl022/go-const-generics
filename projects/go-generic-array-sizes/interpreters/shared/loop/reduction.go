package loop

import (
	"fmt"
)

type Reducer[P Program[E], E Expression] interface {
	Reduce(program P) (P, error)
}

type ProgramReducer[P Program[E], E Expression] struct {
	reducer   Reducer[P, E]
	observers []Observer[E]
}

func NewProgramReducer[P Program[E], E Expression](
	reducer Reducer[P, E], observers []Observer[E],
) *ProgramReducer[P, E] {
	return &ProgramReducer[P, E]{reducer: reducer, observers: observers}
}

type Observer[E Expression] interface {
	Notify(expression E) error
}

func (p ProgramReducer[P, E]) ReduceToValue(program P) (E, error) {
	var nilExpression E
	seenTerms := make(map[string]struct{})

	for !program.Expression().IsValue() {
		if _, alreadySeen := seenTerms[program.Expression().String()]; alreadySeen {
			return nilExpression, newInfiniteLoopErr(program.Expression())
		}
		seenTerms[program.Expression().String()] = struct{}{}

		var err error
		program, err = p.reducer.Reduce(program)

		if err != nil {
			return nilExpression, newStuckProgramErr(err)
		}
		err = p.notifyObservers(program.Expression())
		if err != nil {
			return nilExpression, err
		}
	}
	return program.Expression(), nil
}

func (p ProgramReducer[P, E]) notifyObservers(expression E) error {
	for _, o := range p.observers {
		err := o.Notify(expression)
		if err != nil {
			return err
		}
	}
	return nil
}

type StuckProgramErr struct {
	error
}

func newStuckProgramErr(err error) StuckProgramErr {
	return StuckProgramErr{fmt.Errorf("program stuck: %w", err)}
}

type InfiniteLoopErr struct {
	error
}

func newInfiniteLoopErr(expr Expression) InfiniteLoopErr {
	return InfiniteLoopErr{fmt.Errorf("infinite loop detected at term: %q", expr)}
}
