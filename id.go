package main

type Id struct {
	x Any
}

func NewId(x Any) Id {
	return Id{
		x: x,
	}
}

func (x Id) Of(v Any) Id {
	return NewId(v)
}

func (x Id) Chain(f func(Any) Id) Id {
	return f(x.x)
}

func (x Id) Map(f func(Any) Any) Id {
	return x.Chain(func(x Any) Id {
		return NewId(f(x))
	})
}
