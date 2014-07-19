package main

type Command interface {
	Execute(Note) CommandResult
}

type CommandResult struct {
	Continue bool
	Note     Note
}

func ContinueResult(note Note) CommandResult {
	return CommandResult{
		Continue: true,
		Note:     note,
	}
}

func BreakResult(note Note) CommandResult {
	return CommandResult{
		Continue: false,
		Note:     note,
	}
}
