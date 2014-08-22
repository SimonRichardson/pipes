package main

import (
	"fmt"

	"github.com/SimonRichardson/pipes/pipes"
)

type Sum struct {
	x int
}

func NewSum(x int) Sum {
	return Sum{
		x: x,
	}
}

func (x Sum) Of(v int) Sum {
	return NewSum(v)
}

func (x Sum) Empty() Sum {
	return NewSum(0)
}

func (x Sum) Chain(f func(int) Sum) Sum {
	return f(x.x)
}

func (x Sum) Map(f func(int) int) Sum {
	return x.Chain(func(x int) Sum {
		return NewSum(f(x))
	})
}

func (x Sum) Concat(y Sum) Sum {
	return x.Chain(func(a int) Sum {
		return y.Map(func(b int) int {
			return a + b
		})
	})
}

type AddCommand struct{}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.ContinueResult(note.(Sum).Concat(NewSum(1)))
}

type BadCommand struct{}

func (c BadCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.BreakResult(note)
}

func eff(command pipes.Command) func(pipes.Note) pipes.Either {
	return func(note pipes.Note) pipes.Either {
		res := command.Execute(note)
		return pipes.EitherFromBool(res.Continue, pipes.NewTuple([]pipes.Note{note}, res.Note))
	}
}

func main() {
	note := NewSum(1)

	program := pipes.NewStateT().Of(Sum{}.Empty()).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{}))

	fmt.Println(program.Run(note))
}
