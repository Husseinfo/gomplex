package data

type ComplexNumber struct {
	Real      float32
	Imaginary float32
}

type ComplexNumberType int

const (
	Real      ComplexNumberType = 0
	Imaginary ComplexNumberType = 1
	Complex   ComplexNumberType = 2
)

func (number ComplexNumber) ToInvertedSign() ComplexNumber {
	return ComplexNumber{
		Real:      number.Real * -1,
		Imaginary: number.Imaginary * -1,
	}
}

func (number ComplexNumber) GetType() ComplexNumberType {
	if number.Imaginary == 0 {
		return Real
	}
	if number.Real == 0 {
		return Imaginary
	}

	return Complex
}
