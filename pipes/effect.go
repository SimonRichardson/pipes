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
			return x.Run(a).Chain(func(t Any) Either {
				merge := func(t Tuple) func(Any) Any {
					return func(x Any) Any {
						tuple := x.(Tuple)
						return NewTuple(append(t._1, tuple._1...), tuple._2)
					}
				}
				tup := t.(Tuple)
				return f(tup._2).Bimap(merge(tup), merge(tup))
			})
		},
	}
}

func (x Effect) EvalState(s Note) Either {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._1
	})
}

func (x Effect) ExecState(s Note) Either {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._2
	})
}
