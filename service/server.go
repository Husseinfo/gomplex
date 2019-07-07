package service

import (
	"context"
	"goahead/core"
	"goahead/data"
	"io"
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

type ComplexNumberServer struct {
}

func (s *ComplexNumberServer) Add(stream Complex_AddServer) error {
	sum := data.ComplexNumber{
		Real:      0,
		Imaginary: 0,
	}

	for {
		number, err := stream.Recv()
		if err == io.EOF {
			res := FromComplexNumberDataType(sum)
			_error := stream.SendAndClose(&res)
			if _error != nil {
				panic(_error)
			}
		}

		sum = core.AddComplex(sum, number.ToComplexNumberDataType())
	}
}

func (s *ComplexNumberServer) Sub(stream Complex_SubServer) error {
	sum := data.ComplexNumber{
		Real:      0,
		Imaginary: 0,
	}

	for {
		number, err := stream.Recv()
		if err == io.EOF {
			res := FromComplexNumberDataType(sum)
			_error := stream.SendAndClose(&res)
			if _error != nil {
				panic(_error)
			}
		}

		sum = core.SubComplex(sum, number.ToComplexNumberDataType())
	}
}

func (s *ComplexNumberServer) InvertSign(ctx context.Context, number *ComplexNumber) (*ComplexNumber, error) {
	x := FromComplexNumberDataType(number.ToComplexNumberDataType().ToInvertedSign())
	return &x, nil
}

func (s *ComplexNumberServer) Type(ctx context.Context, number *ComplexNumber) (*NumberType, error) {
	x := NumberType_State(number.ToComplexNumberDataType().GetType())
	return &NumberType{
		Type: x,
	}, nil
}
