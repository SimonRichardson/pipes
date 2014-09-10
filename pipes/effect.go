package pipes

type Effect struct {
	Run func(Note) Either
}

func NewEffect() Effect {
	return Effect{
		Run: func(x Note) Either {
			return nil
		},
	}
}

func (x Effect) Empty() Effect {
	return Effect{
		Run: func(b Note) Either {
			return Right{}.Of(Writer{}.Of(b))
		},
	}
}

func (x Effect) Effect(f func(Note) Either) Effect {
	return Effect{
		Run: func(a Note) Either {
			return x.Run(a).Chain(func(t Writer) Either {
				var e func(Writer) Either
				w := t.Map(func(n Note) Note {
					return f(n).Fold(
						func(a Any) Any {
							e = func(w Writer) Either {
								return NewLeft(w)
							}
							return a
						},
						func(a Any) Any {
							e = func(w Writer) Either {
								return NewRight(w)
							}
							return a
						},
					).(Note)
				})
				return e(t)
			})
		},
	}
}
