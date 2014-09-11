package pipes

import "sync"

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

type Parallel struct {
	Commands []Command
}

func (c Parallel) Execute(note Note) CommandResult {
	commands := c.Commands
	num := len(commands)

	results := make([]CommandResult, num, num)

	var group sync.WaitGroup
	for k, v := range commands {
		group.Add(1)

		go func(i int, command Command) {
			defer group.Done()
			results[i] = command.Execute(note)
		}(k, v)
	}

	group.Wait()

	for i := 0; i < num; i++ {
		result := results[i]
		if !result.Continue {
			return result
		}
	}

	return results[num-1]
}
