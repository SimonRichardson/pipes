package pipes

type Runner struct {
	run func(Note) EitherT
}

func NewRunner(commands []Command) Runner {
	return Runner{
		run: func(note Note) EitherT {
			x := EitherT{}.Of(note)
			for _, v := range commands {
				x = x.Eff(Do(v))
			}
			return x
		},
	}
}

func (r Runner) Exec(note Note) Tuple {
	return r.run(note).Run.Run()
}
