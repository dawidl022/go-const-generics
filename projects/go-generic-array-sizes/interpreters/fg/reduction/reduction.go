package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type ProgramReducer struct {
	observers []Observer
}

func NewProgramReducer(observers []Observer) *ProgramReducer {
	return &ProgramReducer{observers: observers}
}

type Observer interface {
	Notify(expression ast.Expression) error
}

func (p ProgramReducer) ReduceToValue(program ast.Program) (ast.Value, error) {
	seenTerms := make(map[string]struct{})

	for program.Expression.Value() == nil {
		if _, alreadySeen := seenTerms[program.Expression.String()]; alreadySeen {
			return nil, newInfiniteLoopErr(program.Expression)
		}
		seenTerms[program.Expression.String()] = struct{}{}

		var err error
		program, err = program.Reduce()

		if err != nil {
			return nil, newStuckProgramErr(err)
		}
		err = p.notifyObservers(program.Expression)
		if err != nil {
			return nil, err
		}
	}
	return program.Expression.Value(), nil
}

func (p ProgramReducer) notifyObservers(expression ast.Expression) error {
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

func newInfiniteLoopErr(expr ast.Expression) InfiniteLoopErr {
	return InfiniteLoopErr{fmt.Errorf("infinite loop detected at term: %q", expr)}
}
