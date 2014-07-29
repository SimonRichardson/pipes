package junction

type Tuple struct {
	_1 Note
	_2 Note
}

func NewTuple(a, b Note) Tuple {
	return Tuple{
		_1: a,
		_2: b,
	}
}
