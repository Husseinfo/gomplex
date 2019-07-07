package service

import (
	"context"
	"goahead/core"
	"goahead/data"
)

func (m ComplexNumber) ToComplexNumberDataType() data.ComplexNumber {
	return data.ComplexNumber{
		Real:      m.Real,
		Imaginary: m.Imaginary,
	}
}

func FromComplexNumberDataType(number data.ComplexNumber) ComplexNumber {
	return ComplexNumber{
		Real:      number.Real,
		Imaginary: number.Imaginary,
	}
}

type complexNumberServer struct {
}

func (s *complexNumberServer) AddComplex(ctx context.Context, numbers ...*ComplexNumber) (*ComplexNumber, error) {
	sum := data.ComplexNumber{
		Real:      0,
		Imaginary: 0,
	}

	for _, number := range numbers {
		sum = core.AddComplex(sum, number.ToComplexNumberDataType())
	}

	res := FromComplexNumberDataType(sum)
	return &res, nil
}

func (s *complexNumberServer) SubComplex(ctx context.Context, numbers ...*ComplexNumber) (*ComplexNumber, error) {
	sum := data.ComplexNumber{
		Real:      0,
		Imaginary: 0,
	}

	for _, number := range numbers {
		sum = core.SubComplex(sum, number.ToComplexNumberDataType())
	}

	return &ComplexNumber{
		Real:      sum.Real,
		Imaginary: sum.Imaginary,
	}, nil
}

func (s *complexNumberServer) InvertSign(ctx context.Context, number *ComplexNumber) (*ComplexNumber, error) {
	x := FromComplexNumberDataType(number.ToComplexNumberDataType().ToInvertedSign())
	return &x, nil
}

func (s *complexNumberServer) GetType(ctx context.Context, number *ComplexNumber) (*NumberType, error) {
	x := NumberType_State(number.ToComplexNumberDataType().GetType())
	return &NumberType{
		Type: x,
	}, nil
}
