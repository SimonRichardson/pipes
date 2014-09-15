package pipes

type Reader struct {
	Run func(Any) Any
}

func (r Reader) Of(a Any) Reader {
	return Reader{
		Run: constant(a),
	}
}

func (r Reader) Chain(f func(Any) Reader) Reader {
	return Reader{
		Run: func(e Any) Any {
			return f(r.Run(e)).Run(e)
		},
	}
}

func (r Reader) Map(f func(Any) Any) Reader {
	return r.Chain(func(a Any) Reader {
		return Reader{}.Of(f(a))
	})
}

func (r Reader) Ask() Reader {
	return Reader{
		Run: identity(),
	}
}
