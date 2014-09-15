package main

import (
	"fmt"

	"github.com/SimonRichardson/pipes/pipes"
)

type Semigroup interface {
	Concat(a pipes.Any) pipes.Any
}

type Int struct {
	x int
}

func NewInt(x int) Int {
	return Int{
		x: x,
	}
}

func (i Int) Concat(a pipes.Any) pipes.Any {
	return i.x + a.(int)
}

type Get interface {
	Get() Int
}

type Sum struct{}

func NewSum() Sum {
	return Sum{}
}

func (s Sum) GetValue() pipes.Reader {
	return pipes.Reader{
		Run: func(e pipes.Any) pipes.Any {
			return e.(Get).Get()
		},
	}
}

type AddCommand struct{}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
	fmt.Println(run(note.(Sum).GetValue()))
	return pipes.ContinueResult(note)
}

type BadCommand struct{}

func (c BadCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.BreakResult(note)
}

func main() {
	conf := NewConf(NewInt(1))

	// Run the commands manually
	x := pipes.EitherT{}.Of(NewSum()).
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
	fmt.Println("Runner : ", y.Exec(NewSum()))
}
