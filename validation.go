package main

type Validation interface {
	Of(v Any) Validation
	Chain(f func(v Any) Validation) Validation
	Map(f func(v Any) Any) Validation

	Bimap(f func(v Any) Any, g func(v Any) Any) Validation
	Fold(f func(v Any) Any, g func(v Any) Any) Any
}

type Success struct {
	x Any
}

func NewSuccess(x Any) Success {
	return Success{
		x: x,
	}
}

func (x Success) Of(v Any) Validation {
	return NewSuccess(v)
}

func (x Success) Chain(f func(v Any) Validation) Validation {
	return f(x.x)
}

func (x Success) Map(f func(v Any) Any) Validation {
	return x.Of(f(x.x))
}

func (x Success) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return g(x.x)
}

func (x Success) Bimap(f func(v Any) Any, g func(v Any) Any) Validation {
	return NewSuccess(g(x.x))
}

type Failure struct {
	x Any
}

func NewFailure(x Any) Failure {
	return Failure{
		x: x,
	}
}

func (x Failure) Of(v Any) Validation {
	return NewSuccess(v)
}

func (x Failure) Chain(f func(v Any) Validation) Validation {
	return x
}

func (x Failure) Map(f func(v Any) Any) Validation {
	return NewFailure(x.x)
}

func (x Failure) Fold(f func(v Any) Any, g func(v Any) Any) Any {
	return f(x.x)
}

func (x Failure) Bimap(f func(v Any) Any, g func(v Any) Any) Validation {
	return NewFailure(f(x.x))
}

func ValidationFromBool(b bool, val Note) Validation {
	if b {
		return NewSuccess(val)
	}
	return NewFailure(val)
}
