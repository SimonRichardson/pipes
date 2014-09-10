package pipes

type Tuple struct {
	_1 Either
	_2 []Either
}

func NewTuple(a Either, b []Either) Tuple {
	return Tuple{
		_1: a,
		_2: b,
	}
}
