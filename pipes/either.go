package pipes

type Either interface {
	Of(v Writer) Either
	Chain(f func(v Writer) Either) Either
	Map(f func(v Writer) Writer) Either

	Bimap(f func(v Writer) Writer, g func(v Writer) Writer) Either
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}

type Right struct {
	x Writer
}

func NewRight(x Writer) Right {
	return Right{
		x: x,
	}
}

func (x Right) Of(v Writer) Either {
	return NewRight(v)
}

func (x Right) Chain(f func(v Writer) Either) Either {
	return f(x.x)
}

func (x Right) Map(f func(v Writer) Writer) Either {
	return x.Of(f(x.x))
}

func (x Right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Right) Bimap(f func(v Writer) Writer, g func(v Writer) Writer) Either {
	return NewRight(g(x.x))
}

type Left struct {
	x Writer
}

func NewLeft(x Writer) Left {
	return Left{
		x: x,
	}
}

func (x Left) Of(v Writer) Either {
	return NewRight(v)
}

func (x Left) Chain(f func(v Writer) Either) Either {
	return x
}

func (x Left) Map(f func(v Writer) Writer) Either {
	return NewLeft(x.x)
}

func (x Left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Left) Bimap(f func(v Writer) Writer, g func(v Writer) Writer) Either {
	return NewLeft(f(x.x))
}

func EitherFromBool(b bool, val Writer) Either {
	if b {
		return NewRight(val)
	}
	return NewLeft(val)
}
