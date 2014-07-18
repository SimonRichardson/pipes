package main

import "fmt"

type AddCommand struct{}

func (c AddCommand) Execute(note Note) Note {
	return note.(Sum).Concat(NewSum(1))
}

func eff(command Command) func(Note) Note {
	return func(note Note) Note {
		return command.Execute(note)
	}
}

func main() {
	runner := NewStateT().Of(NewSum(1)).Map(eff(AddCommand{}))
	fmt.Println(runner.Run(NewSum(1)))
}
