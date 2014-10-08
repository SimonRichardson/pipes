package pipes

import "fmt"

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

func (t Tuple) Fst() Either {
	return t._1
}

func (t Tuple) Snd() []Either {
	return t._2
}

func (t Tuple) String() string {
	return fmt.Sprintf("(%v, %v)", t._1.(Show).String(), t._2)
}
