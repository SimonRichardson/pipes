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

func (x Effect) Of(a Note) Effect {
	return Effect{
		Run: func(b Note) Either {
			return Right{}.Of(NewTuple([]Note{a}, b))
		},
	}
}

func (x Effect) Effect(f func(Note) Either) Effect {
	return Effect{
		Run: func(a Note) Either {
			return x.Run(a).Chain(func(t Tuple) Either {
				merge := func(t Tuple) func(Tuple) Tuple {
					return func(x Tuple) Tuple {
						return NewTuple(append(t._1, x._1...), x._2)
					}
				}
				return f(t._2).Bimap(merge(t), merge(t))
			})
		},
	}
}
