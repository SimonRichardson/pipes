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
	return NewInt(i.x + a.(Int).x)
}

func (i Int) Extract() int {
	return i.x
}

type Val struct {
	x Int
}

func NewVal(x Int) Val {
	return Val{
		x: x,
	}
}

func (s Val) Sum() pipes.Reader {
	return pipes.Reader{
		Run: func(e pipes.Any) pipes.Any {
			return e.(Int).Concat(s.x)
		},
	}
}

type Conf struct {
	x Val
}

func NewConf(x Val) Conf {
	return Conf{
		x: x,
	}
}

func (c Conf) Map(f func(Val) Val) Conf {
	return NewConf(f(c.x))
}

type AddCommand struct {
	x Int
}

func NewAddCommand(x Int) AddCommand {
	return AddCommand{
		x: x,
	}
}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
	conf := note.(Conf).Map(func(x Val) Val {
		return NewVal(x.Sum().Run(c.x).(Int))
	})
	return pipes.ContinueResult(conf)
}

type BadCommand struct{}

func (c BadCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.BreakResult(note)
}

func main() {
	conf := NewConf(NewVal(NewInt(0)))

	// Run the commands manually
	x := pipes.EitherT{}.Of(conf).
		Eff(pipes.Do(NewAddCommand(NewInt(1)))).
		Eff(pipes.Do(NewAddCommand(NewInt(1)))).
		Eff(pipes.Do(NewAddCommand(NewInt(1))))
	fmt.Println("Manual : ", x.Run.Run())
}
