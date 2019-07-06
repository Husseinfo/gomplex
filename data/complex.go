package data

type Complex struct {
	Real      float32
	Imaginary float32
}

func (number Complex) ToInvertedSign() Complex {
	return Complex{
		Real:      number.Real * -1,
		Imaginary: number.Imaginary * -1,
	}
}
