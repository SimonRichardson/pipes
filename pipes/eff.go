package pipes

func Do(command Command) func(Note) Either {
	return func(note Note) Either {
		res := command.Execute(note)
		return EitherFromBool(res.Continue, res.Note)
	}
}
