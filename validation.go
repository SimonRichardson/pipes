package main

type Validation interface {
	Of(v Any) Validation
	Chain(f func(v Any) Validation) Validation
	Map(f func(v Any) Any) Validation
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
