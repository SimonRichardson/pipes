package pipes

type EitherT struct {
	Run Writer
}

func (e EitherT) Of(x Note) EitherT {
	return EitherT{
		Run: Writer{}.Of(NewRight(x)),
	}
}

func (e EitherT) Swap() EitherT {
	return e.Fold(
		func(l Any) Any {
			return NewRight(l.(Note))
		},
		func(r Any) Any {
			return NewLeft(r.(Note))
		},
	)
}

func (e EitherT) Bimap(f func(Note) Note, g func(Note) Note) EitherT {
	return e.Fold(
		func(l Any) Any {
			return NewLeft(f(l.(Note)))
		},
		func(r Any) Any {
			return NewRight(g(r.(Note)))
		},
	)
}

func (e EitherT) Fold(f func(Any) Any, g func(Any) Any) EitherT {
	return EitherT{
		Run: e.Run.Chain(func(o Either) Writer {
			return Writer{}.Of(o.Fold(f, g).(Either))
		}),
	}
}

func (e EitherT) Chain(f func(Note) EitherT) EitherT {
	return EitherT{
		Run: e.Run.Chain(func(n Either) Writer {
			return n.Fold(
				func(a Any) Any {
					return Writer{}.Of(NewLeft(a.(Note)))
				},
				func(a Any) Any {
					return f(a.(Note)).Run
				},
			).(Writer)
		}),
	}
}

func (e EitherT) Map(f func(Note) Note) EitherT {
	return e.Chain(func(a Note) EitherT {
		return e.Of(f(a))
	})
}

func (e EitherT) Eff(f func(Note) Either) EitherT {
	return EitherT{
		Run: e.Run.Chain(func(n Either) Writer {
			return n.Fold(
				func(a Any) Any {
					return Writer{}.Of(NewLeft(a.(Note)))
				},
				func(a Any) Any {
					return Writer{}.Of(f(a.(Note)))
				},
			).(Writer)
		}),
	}
}
