package service

import (
	"context"
	"goahead/core"
	"goahead/data"
	"io"
	"log"
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
		in, err := stream.Recv()
		if err == io.EOF {
			res := FromComplexNumberDataType(sum)
			err = stream.SendAndClose(&res)
			if err != nil {
				panic(err)
			}
			return nil
		}
		if err != nil {
			return err
		}
		number := in.ToComplexNumberDataType()
		sum = core.AddComplex(sum, number)
		log.Printf("Received %v for addition", number)
		log.Printf("Current Sum is %v", sum)
	}
}

func (s *ComplexNumberServer) Sub(stream Complex_SubServer) error {
	first, _ := stream.Recv()
	sum := first.ToComplexNumberDataType()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			res := FromComplexNumberDataType(sum)
			err = stream.SendAndClose(&res)
			if err != nil {
				panic(err)
			}
			return nil
		}
		if err != nil {
			return err
		}

		number := in.ToComplexNumberDataType()
		sum = core.SubComplex(sum, number)
		log.Printf("Received %v for substraction", number)
		log.Printf("Current Sum is %v", sum)
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
