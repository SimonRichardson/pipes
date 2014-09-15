package pipes

func identity() func(Any) Any {
	return func(x Any) Any {
		return x
	}
}

func constant(x Any) func(Any) Any {
	return func(y Any) Any {
		return x
	}
}
