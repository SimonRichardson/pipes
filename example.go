package main

import (
	"fmt"
	"strconv"

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

func (x Sum) Empty() pipes.Note {
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

func (x Sum) String() string {
	return strconv.Itoa(x.x)
}

type AddCommand struct{}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.ContinueResult(note.(Sum).Concat(NewSum(1)))
}

type BadCommand struct{}

func (c BadCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.BreakResult(note)
}

func main() {
	// Run the commands manually
	x := pipes.EitherT{}.Of(NewSum(1)).
		Eff(pipes.Do(AddCommand{})).
		Eff(pipes.Do(BadCommand{})).
		Eff(pipes.Do(AddCommand{}))
	fmt.Println("Manual : ", x.Run.Run())

	// Run the commands in a runner
	commands := []pipes.Command{
		AddCommand{},
		BadCommand{},
		AddCommand{},
	}
	y := pipes.NewRunner(commands)
	fmt.Println("Runner : ", y.Exec(NewSum(1)))
}
