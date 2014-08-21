package junction

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
			res := x.Run(a)
			return res.Chain(func(t Any) Either {
				tup0 := t.(Tuple)
				return f(tup0._2).Map(func(x Any) Any {
					tup1 := x.(Tuple)
					return NewTuple(append(tup0._1, tup1._1...), tup1._2)
				})
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
