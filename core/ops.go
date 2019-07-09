package core

import (
	"goahead/data"
)

func AddComplex(x data.ComplexNumber, y data.ComplexNumber) data.ComplexNumber {
	return data.ComplexNumber{
		Real:      x.Real + y.Real,
		Imaginary: x.Imaginary + y.Imaginary,
	}
}

func SubComplex(x data.ComplexNumber, y data.ComplexNumber) data.ComplexNumber {
	return AddComplex(x, y.ToInvertedSign())
}
