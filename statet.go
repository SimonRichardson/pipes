package main

type StateT struct {
	mon Validation
	Run func(Note) Validation
}

func NewStateT() StateT {
	return StateT{
		mon: Success{},
		Run: func(x Note) Validation {
			return nil
		},
	}
}

func (x StateT) Of(a Note) StateT {
	return StateT{
		mon: x.mon,
		Run: func(b Note) Validation {
			return x.mon.Of(NewTuple(a, b))
		},
	}
}

func (x StateT) Chain(f func(Note) StateT) StateT {
	return StateT{
		mon: x.mon,
		Run: func(a Note) Validation {
			res := x.Run(a)
			return res.Chain(func(t Any) Validation {
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

func (x StateT) EvalState(s Note) Validation {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._1
	})
}

func (x StateT) ExecState(s Note) Validation {
	return x.Run(s).Map(func(t Any) Any {
		return t.(Tuple)._2
	})
}

func StateGet() StateT {
	return StateT{
		mon: Success{},
		Run: func(s Note) Validation {
			return NewSuccess(NewTuple(s, s))
		},
	}
}

func StateModify(f func(Note) Note) StateT {
	return StateT{
		mon: Success{},
		Run: func(s Note) Validation {
			return NewSuccess(NewTuple(s, f(s)))
		},
	}
}

func StateReshape(f func(Note) Validation) StateT {
	return StateT{
		mon: Success{},
		Run: f,
	}
}

func StatePut(s Note) StateT {
	return StateModify(func(a Note) Note {
		return s
	})
}