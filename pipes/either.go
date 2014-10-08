package pipes

import "fmt"

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

func (x Right) String() string {
	return fmt.Sprintf("Right(%s)", x.x.(Show).String())
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

func (x Left) String() string {
	return fmt.Sprintf("Left(%s)", x.x.(Show).String())
}

func EitherFromBool(b bool, val Note) Either {
	if b {
		return NewRight(val)
	}
	return NewLeft(val)
}
