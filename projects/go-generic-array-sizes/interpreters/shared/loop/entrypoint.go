package loop

import (
	"fmt"
	"io"
)

type Program[E Expression] interface {
	Expression() E
}

type Expression interface {
	IsValue() bool
	fmt.Stringer
}

type Type interface {
	fmt.Stringer
}

type Interpreter[P Program[E], E Expression, T Type] interface {
	ParseProgram(program io.Reader) (P, error)
	TypeCheck(program P) error
	TypeOf(program P) (T, error)
	CheckIsSubtypeOf(program P, subtype, supertype T) error
	ProgramWithExpression(program P, expression E) P
	Reducer[P, E]
}

func Interpret[P Program[E], E Expression, T Type](
	program io.Reader, debugOutput io.Writer, interpreter Interpreter[P, E, T], maxSteps int,
) (string, error) {
	p, err := interpreter.ParseProgram(program)
	if err != nil {
		return "", fmt.Errorf("failed to parse program: %w", err)
	}

	err = interpreter.TypeCheck(p)
	if err != nil {
		return "", fmt.Errorf("type error: %w", err)
	}
	exprType, err := interpreter.TypeOf(p)
	if err != nil {
		panic("call to TypeCheck should have failed if TypeOf main expression returns error")
	}

	debugObserver := &debugObserver[E]{writer: debugOutput}
	typeObserver := &typeCheckingObserver[P, E, T]{interpreter: interpreter, writer: debugOutput, program: p, exprType: exprType}

	reducer := NewProgramReducer[P, E](interpreter, []Observer[E]{debugObserver, typeObserver})

	val, err := reducer.ReduceToValue(p, maxSteps)
	if err != nil {
		return "", err
	}
	return val.String(), nil
}

type debugObserver[E Expression] struct {
	writer     io.Writer
	stepNumber int
}

func (d *debugObserver[E]) Notify(expression E) error {
	d.stepNumber++
	_, err := fmt.Fprintf(d.writer, "reduction step %d: %s\n", d.stepNumber, expression)
	if err != nil {
		return fmt.Errorf("failed to write to debug output: %w", err)
	}
	return nil
}

type typeCheckingObserver[P Program[E], E Expression, T Type] struct {
	interpreter Interpreter[P, E, T]
	writer      io.Writer
	program     P
	exprType    T
}

func (t *typeCheckingObserver[P, E, T]) Notify(expression E) error {
	newProgram := t.ProgramWithExpression(expression)
	err := t.interpreter.TypeCheck(newProgram)
	if err != nil {
		return fmt.Errorf("type error: %s\n", err)
	}
	// TODO consider unit testing preservation using mock visitor
	newExprType, err := t.interpreter.TypeOf(newProgram)
	if err != nil {
		panic("call to TypeCheck should have failed if TypeOf main expression returns error")
	}
	err = t.interpreter.CheckIsSubtypeOf(newProgram, newExprType, t.exprType)
	if err != nil {
		return fmt.Errorf("type preservation violated: %w", err)
	}
	err = handleWriteErrors(func() error {
		_, err = fmt.Fprint(t.writer, "program well typed\n")
		if err != nil {
			return err
		}
		_, err := fmt.Fprintf(t.writer,
			"expression type preserved: expression of type %q is a subtype of previous expression type %q\n\n",
			newExprType, t.exprType,
		)
		return err
	})
	t.exprType = newExprType
	return err
}

func (t *typeCheckingObserver[P, E, T]) ProgramWithExpression(expression E) P {
	return t.interpreter.ProgramWithExpression(t.program, expression)
}

func handleWriteErrors(f func() error) error {
	if err := f(); err != nil {
		return fmt.Errorf("failed to write to debug output: %w", err)
	}
	return nil
}
