package pipes

type Either interface {
	Of(v Note) Either
	Chain(f func(v Note) Either) Either
	Map(f func(v Note) Note) Either

	Bimap(f func(v Note) Note, g func(v Note) Note) Either
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}

type Right struct {
	x Note
}

func NewRight(x Note) Right {
	return Right{
		x: x,
	}
}

func (x Right) Of(v Note) Either {
	return NewRight(v)
}

func (x Right) Chain(f func(v Note) Either) Either {
	return f(x.x)
}

func (x Right) Map(f func(v Note) Note) Either {
	return x.Of(f(x.x))
}

func (x Right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Right) Bimap(f func(v Note) Note, g func(v Note) Note) Either {
	return NewRight(g(x.x))
}

type Left struct {
	x Note
}

func NewLeft(x Note) Left {
	return Left{
		x: x,
	}
}

func (x Left) Of(v Note) Either {
	return NewRight(v)
}

func (x Left) Chain(f func(v Note) Either) Either {
	return x
}

func (x Left) Map(f func(v Note) Note) Either {
	return NewLeft(x.x)
}

func (x Left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Left) Bimap(f func(v Note) Note, g func(v Note) Note) Either {
	return NewLeft(f(x.x))
}

func EitherFromBool(b bool, val Note) Either {
	if b {
		return NewRight(val)
	}
	return NewLeft(val)
}

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
		func(l Note) Either {
			return NewRight(l)
		},
		func(r Note) Either {
			return NewLeft(r)
		},
	)
}

func (e EitherT) Bimap(f func(Note) Note, g func(Note) Note) EitherT {
	return e.Fold(
		func(l Note) Either {
			return NewLeft(f(l))
		},
		func(r Note) Either {
			return NewRight(g(r))
		},
	)
}

func (e EitherT) Fold(f func(Any) Any, g func(Any) Any) Any {
	return e.Run.Chain(func(o Either) Writer {
		return e.Of(o.Fold(f, g))
	})
}

func (e EitherT) Chain(f func(Note) EitherT) EitherT {
	return EitherT{
		Run: e.Run.Chain(func(n Either) {
			return n.Fold(
				func(a Note) Writer {
					return Writer{}.Of(NewLeft(a))
				},
				func(a Note) Writer {
					return f(a).Run
				},
			)
		}),
	}
}

func (e EitherT) Map(f func(Note) Note) EitherT {
	return e.Chain(func(a Note) EitherT {
		return e.Of(f(a))
	})
}
