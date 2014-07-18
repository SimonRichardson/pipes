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
			return res.Chain(func(t Tuple) Validation {
				return f(t._1).Run(t._2)
			})
		},
	}
}

func (x StateT) Map(f func(Note) Note) StateT {
	return x.Chain(func(a Note) StateT {
		return x.Of(f(a))
	})
}
