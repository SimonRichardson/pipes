package junction

type Sum struct {
	x int
}

func NewSum(x int) Sum {
	return Sum{
		x: x,
	}
}

func (x Sum) Of(v int) Sum {
	return NewSum(v)
}

func (x Sum) Empty() Sum {
	return NewSum(0)
}

func (x Sum) Chain(f func(int) Sum) Sum {
	return f(x.x)
}

func (x Sum) Map(f func(int) int) Sum {
	return x.Chain(func(x int) Sum {
		return NewSum(f(x))
	})
}

func (x Sum) Concat(y Sum) Sum {
	return x.Chain(func(a int) Sum {
		return y.Map(func(b int) int {
			return a + b
		})
	})
}
