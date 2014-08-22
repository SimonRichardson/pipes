package pipes

type StateT struct {
	Run func(Note) Either
}

func NewStateT() StateT {
	return StateT{
		Run: func(x Note) Either {
			return nil
		},
	}
}

func (x StateT) Of(a Note) StateT {
	return StateT{
		Run: func(b Note) Either {
			return Right{}.Of(NewTuple([]Note{a}, b))
		},
	}
}

func (x StateT) Chain(f func(Note) StateT) StateT {
	return StateT{
		Run: func(a Note) Either {
			res := x.Run(a)
			return res.Chain(func(t Any) Either {
				tup := t.(Tuple)
				return f(tup._1).Run(tup._2)
			})
		},
	}
}

func (x StateT) Map(f func(Note) Note) StateT {
	return x.Chain(func(a Note) StateT {
		return x.Of(f(a))
	})
}

func (x StateT) Eff(f func(Note) Either) StateT {
	return StateT{
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

func (x StateT) EvalState(s Note) Either {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._1
	})
}

func (x StateT) ExecState(s Note) Either {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._2
	})
}

func StateGet() StateT {
	return StateT{
		Run: func(s Note) Either {
			return NewRight(NewTuple([]Note{s}, s))
		},
	}
}

func StateModify(f func(Note) Note) StateT {
	return StateT{
		Run: func(s Note) Either {
			return NewRight(NewTuple([]Note{s}, f(s)))
		},
	}
}

func StateInject(a Either) StateT {
	return StateT{
		Run: func(b Note) Either {
			return a
		},
	}
}

func StatePut(s Note) StateT {
	return StateModify(func(a Note) Note {
		return s
	})
}
