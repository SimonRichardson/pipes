package main

type Validation interface {
	Of(v Tuple) Validation
	Chain(f func(v Tuple) Validation) Validation
	Map(f func(v Tuple) Tuple) Validation
}

type Success struct {
	x Tuple
}

func NewSuccess(x Tuple) Success {
	return Success{
		x: x,
	}
}

func (x Success) Of(v Tuple) Validation {
	return NewSuccess(v)
}

func (x Success) Chain(f func(v Tuple) Validation) Validation {
	return f(x.x)
}

func (x Success) Map(f func(v Tuple) Tuple) Validation {
	return NewSuccess(f(x.x))
}

type Failure struct {
	x Tuple
}

func NewFailure(x Tuple) Failure {
	return Failure{
		x: x,
	}
}

func (x Failure) Of(v Tuple) Validation {
	return NewSuccess(v)
}

func (x Failure) Chain(f func(v Tuple) Validation) Validation {
	return x
}

func (x Failure) Map(f func(v Tuple) Tuple) Validation {
	return NewFailure(x.x)
}
