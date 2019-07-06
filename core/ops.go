package core

import "goahead/data"

func AddComplex(x data.Complex, y data.Complex) data.Complex {
	return data.Complex{
		Real:      x.Real + y.Real,
		Imaginary: x.Imaginary + y.Imaginary,
	}
}

func SubComplex(x data.Complex, y data.Complex) data.Complex {
	return AddComplex(x, y.ToInvertedSign())
}
