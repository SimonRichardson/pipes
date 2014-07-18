package main

import "fmt"

type AddCommand struct{}

func (c AddCommand) Execute(note Note) CommandResult {
	return ContinueResult(note.(Sum).Concat(NewSum(1)))
}

func eff(command Command) func(Note) Note {
	return func(note Note) Note {
		res := command.Execute(note)
		return res.Note
	}
}

func main() {
	note := NewSum(1)

	runner := NewStateT().Of(NewSum(1)).
		Map(eff(AddCommand{})).
		Map(eff(AddCommand{})).
		Map(eff(AddCommand{}))

	fmt.Println("Eval :", runner.EvalState(note))
	fmt.Println("Exec :", runner.ExecState(note))
}
