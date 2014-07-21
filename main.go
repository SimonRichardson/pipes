package main

import "fmt"

type AddCommand struct{}

func (c AddCommand) Execute(note Note) CommandResult {
	return ContinueResult(note.(Sum).Concat(NewSum(1)))
}

type BadCommand struct{}

func (c BadCommand) Execute(note Note) CommandResult {
	return BreakResult(note)
}

func mapEff(command Command) func(Note) Note {
	return func(note Note) Note {
		res := command.Execute(note)
		return res.Note
	}
}

func chainEff(command Command) func(Note) StateT {
	return func(note Note) StateT {
		res := command.Execute(note)
		return StateInject(ValidationFromBool(res.Continue, NewTuple(note, res.Note)))
	}
}

func main() {
	note := NewSum(1)

	runner := NewStateT().Of(NewSum(1)).
		Map(mapEff(AddCommand{})).
		Map(mapEff(AddCommand{})).
		Chain(chainEff(BadCommand{})).
		Map(mapEff(AddCommand{}))

	fmt.Println("Eval :", runner.EvalState(note))
	fmt.Println("Exec :", runner.ExecState(note))

	fmt.Println("-------")

	free := Suspend{}
	free1 := free.Lift(NewId(1))
	free2 := free.Lift(NewId(2))
	free3 := free.Lift(NewId(3))

	res := free1.Chain(func(x Any) Free {
		return free2.Chain(func(y Any) Free {
			return free3.Map(func(z Any) Any {
				return x.(int) + y.(int) + z.(int)
			})
		})
	})
	fmt.Println(res.Run())
}
