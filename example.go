package main

import (
	"fmt"

	"github.com/SimonRichardson/pipes/pipes"
)

type AddCommand struct{}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
	return pipes.ContinueResult(note.(pipes.Sum).Concat(pipes.NewSum(1)))
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
	note := pipes.NewSum(1)

	program := pipes.NewStateT().Of(pipes.NewSum(0)).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{}))

	fmt.Println(program.Run(note))
}
