package pipes

type Either interface {
	Of(v Tuple) Either
	Chain(f func(v Tuple) Either) Either
	Map(f func(v Tuple) Tuple) Either

	Bimap(f func(v Tuple) Tuple, g func(v Tuple) Tuple) Either
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}

type Right struct {
	x Tuple
}

func NewRight(x Tuple) Right {
	return Right{
		x: x,
	}
}

func (x Right) Of(v Tuple) Either {
	return NewRight(v)
}

func (x Right) Chain(f func(v Tuple) Either) Either {
	return f(x.x)
}

func (x Right) Map(f func(v Tuple) Tuple) Either {
	return x.Of(f(x.x))
}

func (x Right) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Right) Bimap(f func(v Tuple) Tuple, g func(v Tuple) Tuple) Either {
	return NewRight(g(x.x))
}

type Left struct {
	x Tuple
}

func NewLeft(x Tuple) Left {
	return Left{
		x: x,
	}
}

func (x Left) Of(v Tuple) Either {
	return NewRight(v)
}

func (x Left) Chain(f func(v Tuple) Either) Either {
	return x
}

func (x Left) Map(f func(v Tuple) Tuple) Either {
	return NewLeft(x.x)
}

func (x Left) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Left) Bimap(f func(v Tuple) Tuple, g func(v Tuple) Tuple) Either {
	return NewLeft(f(x.x))
}

func EitherFromBool(b bool, val Tuple) Either {
	if b {
		return NewRight(val)
	}
	return NewLeft(val)
}
