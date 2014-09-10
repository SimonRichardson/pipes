package pipes

type Writer struct {
	Run func() Tuple
}

func (w Writer) Of(x Either) Writer {
	return Writer{
		Run: func() Tuple {
			return NewTuple(x, []Either{})
		},
	}
}

func (w Writer) Chain(f func(Either) Writer) Writer {
	return Writer{
		Run: func() Tuple {
			res := w.Run()
			t := f(res._1).Run()
			return NewTuple(t._1, append(res._2, t._2...))
		},
	}
}

func (w Writer) Map(f func(Either) Either) Writer {
	return w.Chain(func(x Either) Writer {
		return Writer{
			Run: func() Tuple {
				return NewTuple(f(x), []Either{})
			},
		}
	})
}
