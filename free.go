package main

type Free interface {
	Of(x Any) Free
	Chain(f func(Any) Free) Free
	Map(x func(Any) Any) Free

	Resume() Validation
	Run() Any
}

type Suspend struct {
	x Id
}

func NewSuspend(x Id) Suspend {
	return Suspend{
		x: x,
	}
}

func (f Suspend) Of(x Any) Free {
	return NewReturn(x)
}

func (f Suspend) Chain(x func(Any) Free) Free {
	return NewSuspend(f.x.Map(func(y Any) Any {
		return y.(Free).Chain(x)
	}))
}

func (f Suspend) Map(x func(Any) Any) Free {
	return f.Chain(func(y Any) Free {
		return f.Of(x(y))
	})
}

func (f Suspend) Resume() Validation {
	return NewFailure(f.x)
}

func (f Suspend) Run() Any {
	return run(f, func(x Any) Free {
		return NewReturn(x)
	})
}

func (f Suspend) Lift(x Id) Free {
	return NewSuspend(x.Map(func(x Any) Any {
		return NewReturn(x)
	}))
}

type Return struct {
	x Any
}

func NewReturn(x Any) Return {
	return Return{
		x: x,
	}
}

func (f Return) Of(x Any) Free {
	return NewReturn(x)
}

func (f Return) Chain(x func(Any) Free) Free {
	return x(f.x)
}

func (f Return) Map(x func(Any) Any) Free {
	return f.Chain(func(y Any) Free {
		return f.Of(x(y))
	})
}

func (f Return) Resume() Validation {
	return NewSuccess(f.x)
}

func (f Return) Run() Any {
	return run(f, func(x Any) Free {
		return NewReturn(x)
	})
}

func run(x Free, f func(Any) Free) Any {
	res := x.Resume()
	cont := true
	for {
		res = res.Bimap(
			func(x Any) Any {
				cont = false
				return x
			},
			func(x Any) Any {
				return f(x).Resume()
			},
		)
		if !cont {
			break
		}
	}
	return res.Fold(Identity, Identity)
}
