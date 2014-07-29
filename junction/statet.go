package junction

type StateT struct {
	mon Either
	Run func(Note) Either
}

func NewStateT() StateT {
	return StateT{
		mon: Right{},
		Run: func(x Note) Either {
			return nil
		},
	}
}

func (x StateT) Of(a Note) StateT {
	return StateT{
		mon: x.mon,
		Run: func(b Note) Either {
			return x.mon.Of(NewTuple(a, b))
		},
	}
}

func (x StateT) Chain(f func(Note) StateT) StateT {
	return StateT{
		mon: x.mon,
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
		mon: Right{},
		Run: func(s Note) Either {
			return NewRight(NewTuple(s, s))
		},
	}
}

func StateModify(f func(Note) Note) StateT {
	return StateT{
		mon: Right{},
		Run: func(s Note) Either {
			return NewRight(NewTuple(s, f(s)))
		},
	}
}

func StateInject(a Either) StateT {
	return StateT{
		mon: Right{},
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
