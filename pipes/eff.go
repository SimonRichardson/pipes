package pipes

func Eff(command Command) func(Note) Either {
	return func(note Note) Either {
		res := command.Execute(note)
		return EitherFromBool(res.Continue, Writer{}.Of(res.Note))
	}
}

func EffF(f func(Note) CommandResult) func(Note) Either {
	return func(note Note) Either {
		res := f(note)
		return EitherFromBool(res.Continue, Writer{}.Of(res.Note))
	}
}

func EffM(f func(Note) Note) func(Note) Either {
	return func(note Note) Either {
		return NewRight(Writer{}.Of(f(note)))
	}
}
