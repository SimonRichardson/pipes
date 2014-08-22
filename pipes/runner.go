package pipes

type Runner struct {
	Commands []Command
}

func NewRunner(commands []Command) Runner {
	return Runner{
		Commands: commands,
	}
}

func (r Runner) Execute(note Note) Either {
	program := NewEffect().Empty()
	for _, v := range r.Commands {
		program = program.Effect(Eff(v))
	}
	return program.Run(note)
}
