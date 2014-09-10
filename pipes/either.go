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
		func(l Any) Any {
			return NewRight(l)
		},
		func(r Any) Any {
			return NewLeft(r)
		},
	)
}

func (e EitherT) Bimap(f func(Note) Note, g func(Note) Note) EitherT {
	return e.Fold(
		func(l Any) Any {
			return NewLeft(f(l))
		},
		func(r Any) Any {
			return NewRight(g(r))
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
					return Writer{}.Of(NewLeft(a))
				},
				func(a Any) Any {
					return f(a).Run
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
