package main

type Free interface {
	Of(x Any) Free
	Resume() Validation
	Run() Any
}

type Suspend struct {
	x Validation
}

func NewSuspend(x Any) Suspend {
	return Suspend{
		x: NewFailure(x),
	}
}

func (f Suspend) Of(x Any) Free {
	return NewReturn(x)
}

func (f Suspend) Resume() Validation {
	return f.x
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
	x Validation
}

func NewReturn(x Any) Return {
	return Return{
		x: NewSuccess(x),
	}
}

func (f Return) Of(x Any) Free {
	return NewReturn(x)
}

func (f Return) Resume() Validation {
	return f.x
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
