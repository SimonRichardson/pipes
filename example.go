package main

import (
	"fmt"
	. "github.com/SimonRichardson/pipes/junction"
)

type AddCommand struct{}

func (c AddCommand) Execute(note Note) CommandResult {
	return ContinueResult(note.(Sum).Concat(NewSum(1)))
}

type BadCommand struct{}

func (c BadCommand) Execute(note Note) CommandResult {
	return BreakResult(note)
}

func eff(command Command) func(Note) Either {
	return func(note Note) Either {
		res := command.Execute(note)
		return EitherFromBool(res.Continue, NewTuple([]Note{note}, res.Note))
	}
}

func main() {
	note := NewSum(1)

	program := NewStateT().Of(NewSum(0)).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{})).
		Eff(eff(AddCommand{}))

	fmt.Println(program.Run(note))
}
