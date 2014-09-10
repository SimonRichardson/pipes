package pipes

type Writer struct {
	Run func() Tuple
}

func (w Writer) Of(x Note) Writer {
	return Writer{
		Run: func() Tuple {
			return NewTuple(x, []Note{})
		},
	}
}

func (w Writer) Chain(f func(Note) Writer) Writer {
	return Writer{
		Run: func() Tuple {
			res := w.Run()
			t := f(res._1).Run()
			return NewTuple(res._1, append(res._2, t._2))
		},
	}
}

func (w Writer) Map(f func(Note) Note) Writer {
	return w.Chain(func(x Note) Writer {
		return Writer{
			Run: func() Tuple {
				return NewTuple(f(x), []Note{})
			},
		}
	})
}
